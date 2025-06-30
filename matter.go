package cesrgo

import (
	"fmt"

	tables "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type matter struct {
	code types.Code
	size types.Size
	raw  types.Raw
}

func (m *matter) SetCode(code types.Code) {
	m.code = code
}

func (m *matter) GetCode() types.Code {
	return m.code
}

func (m *matter) SetRaw(raw types.Raw) {
	m.raw = raw
}

func (m *matter) GetRaw() types.Raw {
	return m.raw
}

func (m *matter) SetSize(size types.Size) {
	m.size = size
}

func (m *matter) GetSize() types.Size {
	return m.size
}

func mbexfil(m types.Matter, qb2 types.Qb2) error {
	if len(qb2) == 0 {
		return fmt.Errorf("qb2 is empty")
	}

	sextets, err := nabSextets(qb2, 1)
	if err != nil {
		return err
	}

	first := sextets[0]
	hs, err := tables.Bardage(first)
	if err != nil {
		return err
	}

	bhs := (hs*3 + 3) / 4
	if len(qb2) < int(bhs) {
		return fmt.Errorf("insufficient material for hard part of code: qb2 size = %d, bhs = %d", len(qb2), bhs)
	}

	hard, err := codeB2ToB64(qb2, int(hs))
	if err != nil {
		return err
	}

	szg, err := tables.GetSizage(types.Code(hard))
	if err != nil {
		return err
	}

	cs := szg.Hs + szg.Ss

	hard, err = codeB2ToB64(qb2, int(hs))
	if err != nil {
		return err
	}

	szg, err = tables.GetSizage(types.Code(hard))
	if err != nil {
		return err
	}

	var fs uint32
	var size uint32
	bcs := ((cs + 1) * 3) / 4

	if szg.Fs == nil {
		if cs%4 != 0 {
			return fmt.Errorf("code size not multiple of 4 for variable length material: cs = %d", cs)
		}

		if len(qb2) < int(bcs) {
			return fmt.Errorf("insufficient material for code: qb2 size = %d, bcs = %d", len(qb2), bcs)
		}

		both, err := codeB2ToB64(qb2, int(cs))
		if err != nil {
			return err
		}

		size, err = b64ToU32(both[szg.Hs:cs])
		if err != nil {
			return err
		}
		fs = (size*4 + cs)
	} else {
		fs = *szg.Fs
	}

	bfs := ((fs + 1) * 3) / 4
	if len(qb2) < int(bfs) {
		return fmt.Errorf("insufficient material: qb2 size = %d, bfs = %d", len(qb2), bfs)
	}

	trim := qb2[:bfs]
	ps := cs % 4

	var pbs uint32
	if ps != 0 {
		pbs = 2 * ps
	} else {
		pbs = 2 * szg.Ls
	}

	if ps != 0 {
		pi := trim[bcs-1]
		if pi&(2<<pbs-1) != 0 {
			return fmt.Errorf("non-zeroed pad bits")
		}
	} else {
		for _, value := range trim[bcs+szg.Ls : bcs+szg.Ls+szg.Ls] {
			if value != 0 {
				if szg.Ls == 1 {
					return fmt.Errorf("non-zeroed lead byte")
				}
				return fmt.Errorf("non-zeroed lead bytes")
			}
		}
	}

	raw := trim[bcs+szg.Ls:]
	if len(raw) != len(trim)-int(bcs)-int(szg.Ls) {
		return fmt.Errorf("improperly qualified material: qb2 = %v", qb2)
	}

	m.SetCode(types.Code(hard))
	m.SetSize(types.Size(size))
	m.SetRaw(types.Raw(raw))

	return nil
}

func mexfil(m types.Matter, qb64 types.Qb64) error {
	if len(qb64) == 0 {
		return fmt.Errorf("qb64 is empty")
	}

	_ = m

	return nil
}

func NewMatter(m types.Matter, opts ...options.MatterOption) error {
	options := &options.MatterOptions{}

	for _, opt := range opts {
		opt(options)
	}

	if options.Code != nil && options.Raw != nil {
		if options.Qb2 != nil || options.Qb64 != nil || options.Qb64b != nil {
			return fmt.Errorf("code and raw cannot be used with qb2, qb64, or qb64b")
		}

		m := &matter{}

		m.SetCode(*options.Code)
		m.SetRaw(*options.Raw)
		m.SetSize(types.Size(len(*options.Raw)))

		return nil
	}

	if options.Qb2 != nil {
		if options.Code != nil || options.Raw != nil || options.Qb64 != nil || options.Qb64b != nil {
			return fmt.Errorf("qb2 cannot be used with code, raw, qb64, or qb64b")
		}

		return mbexfil(m, *options.Qb2)
	}

	if options.Qb64 != nil {
		if options.Code != nil || options.Raw != nil || options.Qb2 != nil || options.Qb64b != nil {
			return fmt.Errorf("qb64 cannot be used with code, raw, qb2, or qb64b")
		}

		return mexfil(m, *options.Qb64)
	}

	if options.Qb64b != nil {
		if options.Code != nil || options.Raw != nil || options.Qb2 != nil || options.Qb64 != nil {
			return fmt.Errorf("qb64b cannot be used with code, raw, qb2, or qb64")
		}

		return mexfil(m, types.Qb64(*options.Qb64b))
	}

	return fmt.Errorf("no inputs provided")
}
