package cesrgo

import (
	"fmt"
	"strings"

	"github.com/jasoncolburne/cesrgo/common"
	"github.com/jasoncolburne/cesrgo/common/util"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type Sad interface {
	GetCode() types.Code
	SetCode(code types.Code)

	GetSize() types.Size
	SetSize(size types.Size)

	GetRaw() types.Raw
	SetRaw(raw types.Raw)

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
	code types.Code
	size types.Size
	raw  types.Raw

	ked     types.Map
	proto   types.Proto
	kind    types.Kind
	version types.Version
}

type Sadder struct {
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

func (s *sad) GetCode() types.Code {
	return s.code
}

func (s *sad) SetCode(code types.Code) {
	s.code = code
}

func (s *sad) GetSize() types.Size {
	return s.size
}

func (s *sad) SetSize(size types.Size) {
	s.size = size
}

func (s *sad) GetRaw() types.Raw {
	return s.raw
}

func (s *sad) SetRaw(raw types.Raw) {
	s.raw = raw
}

func (s *Sadder) inhale(raw types.Raw) error {
	proto, pvrsn, kind, size, _, err := util.Smell(raw)
	if err != nil {
		return err
	}

	if pvrsn.Major != common.VERSION.Major && pvrsn.Minor != common.VERSION.Minor {
		return fmt.Errorf("version mismatch")
	}

	ked, err := util.Unmarshal(kind, raw)
	if err != nil {
		return err
	}

	said, ok := ked.Get("d")
	if !ok {
		return fmt.Errorf("d not found")
	}

	saidStr, ok := said.(string)
	if !ok {
		return fmt.Errorf("d is not a string")
	}

	saider, err := NewSaider(nil, nil, nil, options.WithQb64(types.Qb64(saidStr)))
	if err != nil {
		return err
	}

	s.SetKed(ked)
	s.SetProto(proto)
	s.SetVersion(pvrsn)
	s.SetKind(kind)

	s.SetCode(saider.GetCode())
	//nolint:gosec
	s.SetSize(types.Size(len(raw)))
	s.SetRaw(raw)

	s.saider = saider

	if s.GetSize() != size {
		*s = Sadder{}
		return fmt.Errorf("size mismatch")
	}

	return nil
}

//nolint:gocritic
func (s *Sadder) exhale(ked types.Map, kind *types.Kind) (
	types.Raw,
	types.Proto,
	types.Kind,
	types.Map,
	types.Version,
	error,
) {
	return util.Sizeify(ked, kind, nil)
}

func NewSadder(
	code *types.Code,
	raw *types.Raw,
	ked *types.Map,
	kind *types.Kind,
	saidify bool,
) (*Sadder, error) {
	if code == nil {
		codeStr := codex.Blake3_256
		code = &codeStr
	}

	if kind == nil {
		kindStr := common.Kind_JSON
		kind = &kindStr
	}

	s := &Sadder{}

	if raw != nil {
		if ked != nil {
			return nil, fmt.Errorf("both raw and ked cannot be provided")
		}

		if err := s.inhale(*raw); err != nil {
			return nil, err
		}
	} else if ked != nil {
		szg, ok := codex.Sizes[*code]
		if !ok {
			return nil, fmt.Errorf("unknown code: %s", *code)
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

		s.SetCode(*code)
		//nolint:gosec
		s.SetSize(types.Size(len(exhaledRaw)))
		s.SetRaw(exhaledRaw)
	} else {
		return nil, fmt.Errorf("raw or ked must be provided")
	}

	if saidify {
		saider, err := NewSaider(&s.ked, nil, nil, options.WithCode(*code))
		if err != nil {
			return nil, err
		}

		if s.saider != nil {
			parsedQb64, err := s.saider.Qb64()
			if err != nil {
				return nil, err
			}

			derivedQb64, err := saider.Qb64()
			if err != nil {
				return nil, err
			}

			if parsedQb64 != derivedQb64 {
				return nil, fmt.Errorf("saider mismatch")
			}
		} else {
			s.saider = saider

			ked := s.GetKed()
			qb64, err := saider.Qb64()
			if err != nil {
				return nil, err
			}
			ked.Set("d", qb64)
		}
	}

	return s, nil
}
