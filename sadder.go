package cesrgo

import (
	"fmt"
	"strings"

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

//nolint:gocritic
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
