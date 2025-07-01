package cesrgo

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"

	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type Sad interface {
	GetKed() types.Map
	SetKed(ked types.Map)

	GetProto() types.Proto
	SetProto(proto types.Proto)

	GetKind() types.Kind
	SetKind(kind types.Kind)

	GetVersion() types.Version
	SetVersion(version types.Version)
}

type sad struct {
	ked     types.Map
	proto   types.Proto
	kind    types.Kind
	version types.Version
}

type Sadder struct {
	matter
	sad
	saider *Saider
}

func (s *sad) GetKed() types.Map {
	return s.ked
}

func (s *sad) SetKed(ked types.Map) {
	s.ked = ked
}

func (s *sad) GetProto() types.Proto {
	return s.proto
}

func (s *sad) SetProto(proto types.Proto) {
	s.proto = proto
}

func (s *sad) GetKind() types.Kind {
	return s.kind
}

func (s *sad) SetKind(kind types.Kind) {
	s.kind = kind
}

func (s *sad) GetVersion() types.Version {
	return s.version
}

func (s *sad) SetVersion(version types.Version) {
	s.version = version
}

//nolint:lll
var (
	VER1FULLSPAN = 17
	VER1TERM     = byte("_"[0])
	VEREX1       = "([A-Z]{4})([0-9a-f])([0-9a-f])([A-Z]{4})([0-9a-f]{6})_"

	VER2FULLSPAN = 19
	VER2TERM     = byte("."[0])
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

func hexToUint32(hex []byte) (uint32, error) {
	val, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return 0, err
	}

	return uint32(val), nil
}

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

func (s *Sadder) inhale(raw types.Raw) error {
	proto, pvrsn, kind, size, _, err := smell(raw)
	if err != nil {
		return err
	}

	if pvrsn.Major != VERSION.Major && pvrsn.Minor != VERSION.Minor {
		return fmt.Errorf("version mismatch")
	}

	ked, err := unmarshal(kind, raw)
	if err != nil {
		return err
	}

	s.SetKed(ked)
	s.SetProto(proto)
	s.SetVersion(pvrsn)
	s.SetKind(kind)

	err = NewMatter(s, options.WithCode(s.code), options.WithRaw(raw))
	if err != nil {
		return err
	}

	if s.size != size {
		*s = Sadder{}
		return fmt.Errorf("size mismatch")
	}

	return nil
}

func (s *Sadder) exhale(ked types.Map, kind *types.Kind) (
	types.Raw,
	types.Proto,
	types.Kind,
	types.Map,
	types.Version,
	error,
) {
	return sizeify(ked, kind, nil)
}

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
	copy(rawOut[fore:], []byte(vs))
	copy(rawOut[fore+len(vs):], raw[back:])

	ked.Set("vs", vs)

	return rawOut, proto, *kind, ked, pvrsn, nil
}

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

	offsets := re.FindIndex([]byte(v))
	if offsets == nil {
		return "", types.Version{}, "", 0, &types.Version{}, fmt.Errorf("version string not found")
	}

	fore := offsets[0]
	back := offsets[1]

	full := []byte(v[fore:back])

	return rematch(full, match)
}

func rematch(full []byte, match [][]byte) (
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

	max := length - len(s)
	for i := 0; i < max; i++ {
		s = "A" + s
	}

	return s, nil
}

func NewSadder(
	ked *types.Map,
	kind *types.Kind,
	saidify bool,
	opts ...options.MatterOption,
) (*Sadder, error) {
	config := &options.MatterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	var code = config.Code
	if code == nil {
		codeStr := codex.Blake3_256
		code = &codeStr
	}

	if kind == nil {
		kindStr := Kind_JSON
		kind = &kindStr
	}

	s := &Sadder{}

	raw := config.Raw
	if raw != nil {
		if ked != nil {
			return nil, fmt.Errorf("both raw and ked cannot be provided")
		}

		if err := s.inhale(*raw); err != nil {
			return nil, err
		}
	} else if ked != nil {
		if raw != nil {
			return nil, fmt.Errorf("both raw and ked cannot be provided")
		}

		szg, err := codex.GetSizage(*code)
		if err != nil {
			return nil, err
		}

		kedCopy := ked.Clone()
		if saidify {
			_, ok := kedCopy.Set("d", strings.Repeat("#", int(*szg.Fs)))
			if !ok {
				return nil, fmt.Errorf("failed to set d")
			}
		}

		exhaledRaw, proto, kind, ked, pvrsn, err := s.exhale(kedCopy, kind)
		if err != nil {
			return nil, err
		}

		s.SetKed(ked)
		s.SetProto(proto)
		s.SetVersion(pvrsn)
		s.SetKind(kind)

		err = NewMatter(s, options.WithCode(*code), options.WithRaw(exhaledRaw))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("raw or ked must be provided")
	}

	if saidify {
		saider, err := NewSaider(&s.ked, nil, nil, options.WithCode(*code))
		if err != nil {
			return nil, err
		}

		s.saider = saider

		ked := s.GetKed()
		qb64, err := saider.Qb64()
		if err != nil {
			return nil, err
		}
		ked.Set("d", qb64)
	}

	return s, nil
}
