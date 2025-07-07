package common

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/jasoncolburne/cesrgo"
	"github.com/jasoncolburne/cesrgo/core/types"
	"github.com/vmihailenco/msgpack/v5"
)

var (
	VER1FULLSPAN = 17
	VER1TERM     = '_'
	VEREX1       = "([A-Z]{4})([0-9a-f])([0-9a-f])([A-Z]{4})([0-9a-f]{6})_"

	VER2FULLSPAN = 19
	VER2TERM     = '.'
	VEREX2       = "([A-Z]{4})([0-9A-Za-z_-])([0-9A-Za-z_-]{2})([0-9A-Za-z_-])([0-9A-Za-z_-]{2})([A-Z]{4})([0-9A-Za-z_-]{4})\\."

	B64EX = "^[A-Za-z0-9_-]*$"

	REVER *regexp.Regexp
	REB64 *regexp.Regexp

	MAXVERFULLSPAN = max(VER1FULLSPAN, VER2FULLSPAN)
	MAXVSOFFSET    = 12

	SMELLSIZE = MAXVSOFFSET + MAXVERFULLSPAN // # min buffer size to inhale
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

func Reb64() (*regexp.Regexp, error) {
	if REB64 == nil {
		var err error
		REB64, err = regexp.Compile(B64EX)
		if err != nil {
			return nil, err
		}
	}

	return REB64, nil
}

func ValidateCode(code types.Code, validCodes []types.Code) bool {
	return slices.Contains(validCodes, code)
}

var b64Runes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
var b64Indices = map[rune]uint8{}

func generateb64Indices() {
	if len(b64Indices) > 0 {
		return
	}

	for i, c := range b64Runes[:64] {
		//nolint:gosec
		b64Indices[c] = uint8(i)
	}
}

func B64CharToIndex(c rune) (uint8, error) {
	generateb64Indices()

	index, ok := b64Indices[c]
	if !ok {
		return 0, fmt.Errorf("invalid url-safe base64 character: %c", c)
	}

	return index, nil
}

func B64IndexToChar(i uint8) (byte, error) {
	if i > 63 {
		return 0, fmt.Errorf("programmer error:invalid base64 index: %d", i)
	}

	return b64Runes[i], nil
}

func NabSextets(bin []byte, count int) ([]byte, error) {
	n := int(math.Ceil(float64(count) * 3 / 4))

	if n > len(bin) {
		return nil, fmt.Errorf("not enough bytes in %v to nab %d sextets", bin, count)
	}

	inInt := BytesToBigInt(bin[:n])
	p := 2 * (count % 4)

	inInt.Rsh(inInt, uint(p))
	out := make([]byte, count)
	for j := count - 1; j >= 0; j-- {
		out[j] = byte(inInt.Uint64() & 0x3F)
		inInt.Rsh(inInt, 6)
	}

	return out, nil
}

func BytesToBigInt(bin []byte) *big.Int {
	return big.NewInt(0).SetBytes(bin)
}

func BigIntToB64(n *big.Int, b64Length int) (string, error) {
	if b64Length == 0 {
		return "", nil
	}

	x := big.NewInt(0).Set(n)
	out := ""

	for x.Cmp(big.NewInt(0)) > 0 {
		c, err := B64IndexToChar(byte(x.Uint64() % 64))
		if err != nil {
			return "", err
		}
		out = string(c) + out

		x.Div(x, big.NewInt(64))
	}

	for len(out) < b64Length {
		out = "A" + out
	}

	return out[:b64Length], nil
}

func B64ToBigInt(b64 string) (*big.Int, error) {
	n := big.NewInt(0)

	for _, c := range b64 {
		i, err := B64CharToIndex(c)
		if err != nil {
			return nil, err
		}
		n.Mul(n, big.NewInt(64))
		n.Add(n, big.NewInt(int64(i)))
	}

	return n, nil
}

func CodeB2ToB64(b2 []byte, length int) (string, error) {
	if length == 0 {
		return "", nil
	}

	n := int(math.Ceil(float64(length) * 3 / 4))

	if n > len(b2) {
		return "", fmt.Errorf("not enough bytes")
	}

	sextets, err := NabSextets(b2[:n], length)
	if err != nil {
		return "", err
	}

	out := make([]byte, length)
	for i, sextet := range sextets {
		c, err := B64IndexToChar(sextet)
		if err != nil {
			return "", err
		}
		out[i] = c
	}

	return string(out), nil
}

func CodeB64ToB2(code string) ([]byte, error) {
	i, err := B64ToBigInt(code)

	if err != nil {
		return nil, err
	}

	i.Lsh(i, uint(2*(len(code)%4)))
	n := int(math.Ceil(float64(len(code)) * 3 / 4))
	out := make([]byte, n)
	i.FillBytes(out)
	return out, nil
}

func B64ToU32(b64 string) (uint32, error) {
	var out uint32 = 0

	for _, c := range b64 {
		i, err := B64CharToIndex(c)
		if err != nil {
			return 0, err
		}
		out = (out << 6) + uint32(i)
	}

	return out, nil
}

func B64ToU64(b64 string) (uint64, error) {
	var out uint64

	for _, c := range b64 {
		i, err := B64CharToIndex(c)
		if err != nil {
			return 0, err
		}
		out = (out << 6) + uint64(i)
	}

	return out, nil
}

func IntToB64(n, length int) (string, error) {
	s := ""

	for n > 0 {
		//nolint:gosec
		c, err := B64IndexToChar(uint8(n % 64))
		if err != nil {
			return "", err
		}
		s = string([]byte{c}) + s
		n /= 64
	}

	limit := length - len(s)
	for range limit {
		s = "A" + s
	}

	return s, nil
}

func HexToUint32(hex []byte) (uint32, error) {
	val, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return 0, err
	}

	return uint32(val), nil
}

//nolint:gocritic
func Smell(raw types.Raw) (types.Proto, types.Version, types.Kind, types.Size, *types.Version, error) {
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

		pmajor, err := HexToUint32(match[2])
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		pminor, err := HexToUint32(match[3])
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		kind := types.Kind(match[4])

		sizeInt, err := HexToUint32(match[5])
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

		pmajor, err := B64ToU32(string(match[7]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		pminor, err := B64ToU32(string(match[8]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		gmajor, err := B64ToU32(string(match[9]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		gminor, err := B64ToU32(string(match[10]))
		if err != nil {
			return "", types.Version{}, "", 0, nil, err
		}

		kind := types.Kind(match[11])

		sizeInt, err := B64ToU32(string(match[12]))
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
func Sizeify(ked types.Map, kind *types.Kind, version *types.Version) (
	types.Raw,
	types.Proto,
	types.Kind,
	types.Map,
	types.Version,
	error,
) {
	vAny, ok := ked.Get("v")
	if !ok {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("version string not found")
	}

	v, ok := vAny.(string)
	if !ok {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("version string not a string")
	}

	if version == nil {
		version = &cesrgo.VERSION
	}

	proto, pvrsn, knd, _, gvrsn, err := Deversify(v)
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

	if !slices.Contains(cesrgo.KINDS, *kind) {
		return nil, "", "", types.Map{}, types.Version{}, fmt.Errorf("kind not supported")
	}

	raw, err := Marshal(ked, kind)
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

	vs, err := Versify(&proto, &pvrsn, kind, size, nil)
	if err != nil {
		return nil, "", "", types.Map{}, types.Version{}, err
	}

	rawOut := make([]byte, len(raw[:fore])+len(vs)+len(raw[back:]))
	copy(rawOut, raw[:fore])
	copy(rawOut[fore:], vs)
	copy(rawOut[fore+len(vs):], raw[back:])

	ked.Set("v", vs)

	return rawOut, proto, *kind, ked, pvrsn, nil
}

func Versify(proto *types.Proto, pvrsn *types.Version, kind *types.Kind, size types.Size, gvrsn *types.Version) (string, error) {
	if proto == nil {
		protoKeri := cesrgo.Proto_KERI
		proto = &protoKeri
	}

	if pvrsn == nil {
		pvrsn = &cesrgo.VERSION
	}

	if kind == nil {
		kindJson := cesrgo.Kind_JSON
		kind = &kindJson
	}

	if gvrsn == nil {
		gvrsn = pvrsn
	}

	if !slices.Contains(cesrgo.PROTOS, *proto) {
		return "", fmt.Errorf("proto not supported")
	}

	if !slices.Contains(cesrgo.KINDS, *kind) {
		return "", fmt.Errorf("kind not supported")
	}

	if pvrsn.Major < 2 || gvrsn.Major < 2 {
		return "", fmt.Errorf("major versions must be >= 2")
	}

	pvmaj, err := IntToB64(int(pvrsn.Major), 1)
	if err != nil {
		return "", err
	}

	pvmin, err := IntToB64(int(pvrsn.Minor), 2)
	if err != nil {
		return "", err
	}

	gvmaj, err := IntToB64(int(gvrsn.Major), 1)
	if err != nil {
		return "", err
	}

	gvmin, err := IntToB64(int(gvrsn.Minor), 2)
	if err != nil {
		return "", err
	}

	sz, err := IntToB64(int(size), 4)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s%s%s%s%s%s.", *proto, pvmaj, pvmin, gvmaj, gvmin, *kind, sz), nil
}

//nolint:gocritic
func Deversify(v string) (
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

	return Rematch(match)
}

//nolint:gocritic
func Rematch(match [][]byte) (
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
		pmajor, err := B64ToU32(string(match[7]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		pminor, err := B64ToU32(string(match[8]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		gmajor, err := B64ToU32(string(match[9]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		gminor, err := B64ToU32(string(match[10]))
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		kind := types.Kind(match[11])

		sizeInt, err := B64ToU32(string(match[12]))
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
		pmajor, err := HexToUint32(match[2])
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		pminor, err := HexToUint32(match[3])
		if err != nil {
			return "", types.Version{}, "", 0, &types.Version{}, err
		}

		kind := types.Kind(match[4])

		sizeInt, err := HexToUint32(match[5])
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

func Marshal(ked types.Map, kind *types.Kind) (types.Raw, error) {
	if kind == nil {
		kindJson := cesrgo.Kind_JSON
		kind = &kindJson
	}

	var (
		raw types.Raw
		err error
	)

	switch *kind {
	case cesrgo.Kind_JSON:
		raw, err = json.Marshal(ked)
		if err != nil {
			return nil, err
		}
	case cesrgo.Kind_CBOR:
		_, _ = cbor.Marshal(ked)

		// need to implement MarshalCBOR for types.Map
		return nil, fmt.Errorf("cbor not yet supported")
	case cesrgo.Kind_MGPK:
		_, _ = msgpack.Marshal(ked)

		// need to investigate marshaling msgpack for types.Map
		return nil, fmt.Errorf("mgpk not yet supported")
	default:
		return nil, fmt.Errorf("unsupported kind: %s", *kind)
	}

	return raw, nil
}

func Unmarshal(kind types.Kind, raw types.Raw) (types.Map, error) {
	var ked types.Map

	switch kind {
	case cesrgo.Kind_JSON:
		err := json.Unmarshal(raw, &ked)
		if err != nil {
			return types.Map{}, err
		}
	case cesrgo.Kind_CBOR:
		_ = cbor.Unmarshal(raw, &ked)

		return types.Map{}, fmt.Errorf("cbor not yet supported")
	case cesrgo.Kind_MGPK:
		_ = msgpack.Unmarshal(raw, &ked)

		return types.Map{}, fmt.Errorf("mgpk not yet supported")
	default:
		return types.Map{}, fmt.Errorf("unsupported kind: %s", kind)
	}

	return ked, nil
}

func NowISO8601() types.DateTime {
	timeStr := time.Now().UTC().Format("2006-01-02T15:04:05.000000Z07:00")
	if strings.HasSuffix(timeStr, "Z") {
		timeStr = strings.Replace(timeStr, "Z", "+00:00", 1)
	}
	return types.DateTime(timeStr)
}

func LessThanMaxON(num big.Int) bool {
	limit := &big.Int{}
	limit.Exp(big.NewInt(2), big.NewInt(128), nil)

	return num.Cmp(limit) < 0
}
