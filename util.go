package cesrgo

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/jasoncolburne/cesrgo/types"
	"github.com/vmihailenco/msgpack/v5"
)

var (
	VER1FULLSPAN = 17
	VER1TERM     = '_'
	VEREX1       = "([A-Z]{4})([0-9a-f])([0-9a-f])([A-Z]{4})([0-9a-f]{6})_"

	VER2FULLSPAN = 19
	VER2TERM     = '.'
	VEREX2       = "([A-Z]{4})([0-9A-Za-z_-])([0-9A-Za-z_-]{2})([0-9A-Za-z_-])([0-9A-Za-z_-]{2})([A-Z]{4})([0-9A-Za-z_-]{4})\\."

	REVER *regexp.Regexp

	MAXVERFULLSPAN = max(VER1FULLSPAN, VER2FULLSPAN)
)

var (
	MAXVSOFFSET = 12
	SMELLSIZE   = MAXVSOFFSET + MAXVERFULLSPAN // # min buffer size to inhale
)

func Rever() (*regexp.Regexp, error) {
	if REVER == nil {
		var err error
		REVER, err = regexp.Compile(VEREX1 + "|" + VEREX2)
		if err != nil {
			return nil, err
		}
	}

	return REVER, nil
}

func validateCode(code types.Code, validCodes []types.Code) bool {
	return slices.Contains(validCodes, code)
}

var b64Runes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func b64CharToIndex(c byte) (uint8, error) {
	index := strings.Index(b64Runes, string(c))
	if index == -1 {
		return 0, fmt.Errorf("invalid base64 character: %c", c)
	}

	if index > 63 {
		return 0, fmt.Errorf("programmer error:invalid base64 character: %c", c)
	}

	return uint8(index), nil
}

func b64IndexToChar(i uint8) (byte, error) {
	if i > 63 {
		return 0, fmt.Errorf("programmer error:invalid base64 index: %d", i)
	}

	return b64Runes[i], nil
}

func nabSextets(binary []byte, count int) ([]byte, error) {
	n := ((count + 1) * 3) / 4

	if n > len(binary) {
		return nil, fmt.Errorf("binary is too small")
	}

	bps := 3 - (len(binary) % 3)
	padded := make([]byte, len(binary)+bps)
	copy(padded, binary)

	out := make([]byte, len(padded)*4/3)
	i := 0
	j := 0
	for {
		n := uint32(padded[i])<<16 | uint32(padded[i+1])<<8 | uint32(padded[i+2])

		out[j] = byte((n & 0xfc0000) >> 18)
		out[j+1] = byte((n & 0x03f000) >> 12)
		out[j+2] = byte((n & 0x000fc0) >> 6)
		out[j+3] = byte(n & 0x00003f)

		j += 4
		i += 3
		if i >= len(padded) {
			break
		}
	}

	return out[:count], nil
}

func codeB2ToB64(b2 []byte, length int) (string, error) {
	if length == 0 {
		return "", nil
	}

	n := int(math.Ceil(float64(length) * 3 / 4))

	if n > len(b2) {
		return "", fmt.Errorf("not enough bytes")
	}

	tbs := 2 * (length % 4)
	if length <= 4 {
		bytes := [4]byte{}
		copy(bytes[:], b2[:n])

		i := binary.BigEndian.Uint32(bytes[:])
		return u32ToB64(i>>tbs, length)
	} else if length <= 8 {
		bytes := [8]byte{}
		copy(bytes[:], b2[:n])

		i := binary.BigEndian.Uint64(bytes[:])
		return u64ToB64(i>>tbs, length)
	} else {
		return "", fmt.Errorf("unexpected length")
	}
}

// func codeB64ToB2(code string) ([]byte, error) {
// 	i, err := b64ToU64(code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	i <<= 2 * (len(code) % 4)
// 	n := ((len(code) + 1) * 3) / 4
// 	return binary.BigEndian.AppendUint64(make([]byte, 0, 8), i)[8-n:], nil
// }

func b64ToU32(b64 string) (uint32, error) {
	var out uint32 = 0

	for _, c := range b64 {
		i, err := b64CharToIndex(byte(c))
		if err != nil {
			return 0, err
		}
		out = (out << 6) + uint32(i)
	}

	return out, nil
}

func b64ToU64(b64 string) (uint64, error) {
	var out uint64 = 0

	for _, c := range b64 {
		i, err := b64CharToIndex(byte(c))
		if err != nil {
			return 0, err
		}
		out = (out << 6) + uint64(i)
	}

	return out, nil
}

func u32ToB64(n uint32, length int) (string, error) {
	if length == 0 {
		return "", nil
	}

	var x uint32 = n
	out := ""

	var overflow float64 = float64(length) - math.Log2(float64(n))/math.Log2(64)
	for x > 0 {
		if overflow >= 0.0 {
			i, err := b64IndexToChar(byte(x % 64))
			if err != nil {
				return "", err
			}
			out = string(i) + out
		} else {
			overflow += 1
		}
		x /= 64
	}

	for len(out) < length {
		out = "A" + out
	}

	return out, nil
}

func u64ToB64(n uint64, length int) (string, error) {
	if length == 0 {
		return "", nil
	}

	var x uint64 = n
	out := ""

	var overflow float64 = float64(length) - math.Log2(float64(n))/math.Log2(64)
	for x > 0 {
		if overflow >= 0.0 {
			c, err := b64IndexToChar(byte(x % 64))
			if err != nil {
				return "", err
			}
			out = string([]byte{c}) + out
		} else {
			overflow += 1
		}
		x /= 64
	}

	for len(out) < length {
		out = "A" + out
	}

	return out, nil
}

func intToB64(n, length int) (string, error) {
	s := ""

	for n > 0 {
		//nolint:gosec
		c, err := b64IndexToChar(uint8(n % 64))
		if err != nil {
			return "", err
		}
		s = string([]byte{c}) + s
		n /= 64
	}

	limit := length - len(s)
	for i := 0; i < limit; i++ {
		s = "A" + s
	}

	return s, nil
}

func hexToUint32(hex []byte) (uint32, error) {
	val, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return 0, err
	}

	return uint32(val), nil
}

//nolint:gocritic
func smell(raw types.Raw) (types.Proto, types.Version, types.Kind, types.Size, *types.Version, error) {
	re, err := Rever()
	if err != nil {
		return "", types.Version{}, "", 0, nil, err
	}

	match := re.FindSubmatch(raw)
	if len(match) != 13 {
		return "", types.Version{}, "", 0, nil, fmt.Errorf("invalid version")
	}

	if len(match[1]) > 0 {
		proto := types.Proto(match[1])

		pmajor, err := hexToUint32(match[2])
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		pminor, err := hexToUint32(match[3])
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		kind := types.Kind(match[4])

		sizeInt, err := hexToUint32(match[5])
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		size := types.Size(sizeInt)

		return proto, types.Version{
			Major: pmajor,
			Minor: pminor,
		}, kind, size, nil, nil
	} else if len(match[6]) > 0 {
		proto := types.Proto(match[6])

		pmajor, err := b64ToU32(string(match[7]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		pminor, err := b64ToU32(string(match[8]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		gmajor, err := b64ToU32(string(match[9]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		gminor, err := b64ToU32(string(match[10]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		kind := types.Kind(match[11])

		sizeInt, err := b64ToU32(string(match[12]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		size := types.Size(sizeInt)

		return proto, types.Version{
				Major: pmajor,
				Minor: pminor,
			}, kind, size, &types.Version{
				Major: gmajor,
				Minor: gminor,
			}, nil
	}

	return "", types.Version{}, "", 0, nil, fmt.Errorf("invalid version")
}

//nolint:gocritic
func sizeify(ked types.Map, kind *types.Kind, version *types.Version) (
	types.Raw,
	types.Proto,
	types.Kind,
	types.Map,
	types.Version,
	error,
) {
	vAny, ok := ked.Get("vs")
	if !ok {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("version string not found")
	}

	v, ok := vAny.(string)
	if !ok {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("version string not a string")
	}

	if version == nil {
		version = &VERSION
	}

	proto, pvrsn, knd, _, gvrsn, err := deversify(v)
	if err != nil {
		return nil, "", "", types.Map{}, types.Version{}, err
	}

	if pvrsn.Major != version.Major || pvrsn.Minor != version.Minor {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("protocol version mismatch")
	}

	if gvrsn != nil && (gvrsn.Major != version.Major || gvrsn.Minor != version.Minor) {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("genus version mismatch")
	}

	if kind == nil {
		kind = &knd
	}

	if !slices.Contains(KINDS, *kind) {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("kind not supported")
	}

	raw, err := marshal(ked, kind)
	if err != nil {
		return nil, "", "", types.Map{}, types.Version{}, err
	}

	length := len(raw)
	if length > 1<<32-1 {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("size too large")
	}

	//nolint:gosec
	size := types.Size(len(raw))

	re, err := Rever()
	if err != nil {
		return nil, "", "", types.Map{}, types.Version{}, err
	}

	offset := re.FindIndex(raw)
	if offset == nil {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("version string not found")
	}

	fore := offset[0]
	back := offset[1]

	vs, err := versify(&proto, &pvrsn, kind, size, nil)
	if err != nil {
		return nil, "", "", types.Map{}, types.Version{}, err
	}

	rawOut := make([]byte, len(raw[:fore])+len(vs)+len(raw[back:]))
	copy(rawOut, raw[:fore])
	copy(rawOut[fore:], vs)
	copy(rawOut[fore+len(vs):], raw[back:])

	ked.Set("vs", vs)

	return rawOut, proto, *kind, ked, pvrsn, nil
}

func versify(proto *types.Proto, pvrsn *types.Version, kind *types.Kind, size types.Size, gvrsn *types.Version) (string, error) {
	if proto == nil {
		protoKeri := Proto_KERI
		proto = &protoKeri
	}

	if pvrsn == nil {
		pvrsn = &VERSION
	}

	if kind == nil {
		kindJson := Kind_JSON
		kind = &kindJson
	}

	if gvrsn == nil {
		gvrsn = pvrsn
	}

	if !slices.Contains(PROTOS, *proto) {
		return "", fmt.Errorf("proto not supported")
	}

	if !slices.Contains(KINDS, *kind) {
		return "", fmt.Errorf("kind not supported")
	}

	if pvrsn.Major < 2 || gvrsn.Major < 2 {
		return "", fmt.Errorf("major versions must be >= 2")
	}

	pvmaj, err := intToB64(int(pvrsn.Major), 1)
	if err != nil {
		return "", err
	}

	pvmin, err := intToB64(int(pvrsn.Minor), 2)
	if err != nil {
		return "", err
	}

	gvmaj, err := intToB64(int(gvrsn.Major), 1)
	if err != nil {
		return "", err
	}

	gvmin, err := intToB64(int(gvrsn.Minor), 2)
	if err != nil {
		return "", err
	}

	sz, err := intToB64(int(size), 4)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s%s%s%s%s%s.", *proto, pvmaj, pvmin, gvmaj, gvmin, *kind, sz), nil
}

//nolint:gocritic
func deversify(v string) (
	types.Proto,
	types.Version,
	types.Kind,
	types.Size,
	*types.Version,
	error,
) {
	re, err := Rever()
	if err != nil {
		return "", types.Version{}, "", 0, &types.Version{}, err
	}

	match := re.FindSubmatch([]byte(v))
	if match == nil {
		return "", types.Version{}, "", 0, &types.Version{}, fmt.Errorf("version string not found")
	}

	return rematch(match)
}

//nolint:gocritic
func rematch(match [][]byte) (
	types.Proto,
	types.Version,
	types.Kind,
	types.Size,
	*types.Version,
	error,
) {
	if len(match) != 13 {
		return "", types.Version{}, "", 0, &types.Version{}, fmt.Errorf("invalid version")
	}

	if len(match[6]) > 0 {
		proto := types.Proto(match[6])
		pmajor, err := b64ToU32(string(match[7]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		pminor, err := b64ToU32(string(match[8]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		gmajor, err := b64ToU32(string(match[9]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		gminor, err := b64ToU32(string(match[10]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		kind := types.Kind(match[11])

		sizeInt, err := b64ToU32(string(match[12]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		size := types.Size(sizeInt)

		return proto, types.Version{
				Major: pmajor,
				Minor: pminor,
			}, kind, size, &types.Version{
				Major: gmajor,
				Minor: gminor,
			}, nil
	} else if len(match[1]) > 0 {
		proto := types.Proto(match[1])
		pmajor, err := hexToUint32(match[2])
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		pminor, err := hexToUint32(match[3])
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		kind := types.Kind(match[4])

		sizeInt, err := hexToUint32(match[5])
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		size := types.Size(sizeInt)

		return proto, types.Version{
			Major: pmajor,
			Minor: pminor,
		}, kind, size, nil, nil
	}

	return "", types.Version{}, "", 0, &types.Version{}, fmt.Errorf("invalid version")
}

func marshal(ked types.Map, kind *types.Kind) (types.Raw, error) {
	if kind == nil {
		kindJson := Kind_JSON
		kind = &kindJson
	}

	var (
		raw types.Raw
		err error
	)

	switch *kind {
	case Kind_JSON:
		raw, err = json.Marshal(ked)
		if err != nil {
			return nil, err
		}
	case Kind_CBOR:
		raw, err = cbor.Marshal(ked)
		if err != nil {
			return nil, err
		}
	case Kind_MGPK:
		raw, err = msgpack.Marshal(ked)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported kind: %s", *kind)
	}

	return raw, nil
}

func unmarshal(kind types.Kind, raw types.Raw) (types.Map, error) {
	var ked types.Map

	switch kind {
	case Kind_JSON:
		err := json.Unmarshal(raw, &ked)
		if err != nil {
			return types.Map{}, err
		}
	case Kind_CBOR:
		err := cbor.Unmarshal(raw, &ked)
		if err != nil {
			return types.Map{}, err
		}
	case Kind_MGPK:
		err := msgpack.Unmarshal(raw, &ked)
		if err != nil {
			return types.Map{}, err
		}
	default:
		return types.Map{}, fmt.Errorf("unsupported kind: %s", kind)
	}

	return ked, nil
}
