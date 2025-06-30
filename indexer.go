package cesrgo

import (
	"fmt"

	tables "github.com/jasoncolburne/cesrgo/indexer"
	"github.com/jasoncolburne/cesrgo/indexer/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type indexer struct {
	code  types.Code
	raw   types.Raw
	index types.Index
	ondex types.Ondex
}

func (i *indexer) SetCode(code types.Code) {
	i.code = code
}

func (i *indexer) GetCode() types.Code {
	return i.code
}

func (i *indexer) SetRaw(raw types.Raw) {
	i.raw = raw
}

func (i *indexer) GetRaw() types.Raw {
	return i.raw
}

func (i *indexer) SetIndex(index types.Index) {
	i.index = index
}

func (i *indexer) GetIndex() types.Index {
	return i.index
}

func (i *indexer) SetOndex(ondex types.Ondex) {
	i.ondex = ondex
}

func (i *indexer) GetOndex() types.Ondex {
	return i.ondex
}

func ibexfil(i types.Indexer, qb2 types.Qb2) error {
	if len(qb2) == 0 {
		return fmt.Errorf("qb2 is empty")
	}

	first, err := nabSextets(qb2, 1)
	if err != nil {
		return err
	}

	hs, err := tables.GetBardage(first[0])
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
	ms := szg.Ss - szg.Os
	bcs := ((cs + 1) * 3) / 4

	if len(qb2) < int(bcs) {
		return fmt.Errorf("insufficient material for code: qb2 size = %d, bcs = %d", len(qb2), bcs)
	}

	both, err := codeB2ToB64(qb2, int(cs))
	if err != nil {
		return err
	}

	index, err := b64ToU32(both[hs : hs+ms])
	if err != nil {
		return err
	}

	var ondex *types.Ondex
	if validateCode(types.Code(hard), tables.ValidCurrentSigCodes) {
		if szg.Os != 0 {
			odx, err := b64ToU32(both[hs+ms : hs+(ms+szg.Os)])
			if err != nil {
				return err
			}

			if odx != 0 {
				return fmt.Errorf("invalid ondex = '%d' for code = '%s'", odx, hard)
			}
		}
	} else if szg.Os != 0 {
		odx, err := b64ToU32(both[hs+ms : hs+(ms+szg.Os)])
		if err != nil {
			return err
		}
		o := types.Ondex(odx)
		ondex = &o
	} else {
		odx := types.Ondex(index)
		ondex = &odx
	}

	var fs uint32
	if szg.Fs == nil {
		if cs%4 != 0 {
			// unreachable unless sizages are broken
			return fmt.Errorf("code size not multiple of 4 for variable length material: cs = %d", cs)
		}

		if szg.Os != 0 {
			// unreachable using current tables
			return fmt.Errorf("non-zero other index size for variable length material: os = %d", szg.Os)
		}

		fs = (index * 4) + cs
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
		bytes := [1]byte{trim[bcs-1]}
		pi := uint8(bytes[0])
		if pi&(1<<pbs-1) != 0 {
			return fmt.Errorf("non-zeroed pad bits")
		}
	} else {
		for _, value := range trim[bcs+szg.Ls:] {
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
		// unreachable. rust prevents this by the definition of `raw` above.
		return fmt.Errorf("improperly qualified material: qb2 = %v", qb2)
	}

	i.SetCode(types.Code(hard))
	i.SetRaw(types.Raw(raw))
	i.SetIndex(types.Index(index))
	if ondex != nil {
		i.SetOndex(*ondex)
	}

	return nil
}

func iexfil(i types.Indexer, qb64 types.Qb64) error {
	if len(qb64) == 0 {
		return fmt.Errorf("qb64 is empty")
	}

	_ = i

	return nil
}

func NewIndexer(i types.Indexer, opts ...options.IndexerOption) error {
	config := &options.IndexerOptions{}
	for _, opt := range opts {
		opt(config)
	}

	if config.Code != nil && config.Raw != nil && config.Index != nil {
		if config.Qb2 != nil || config.Qb64 != nil || config.Qb64b != nil {
			return fmt.Errorf("qb2, qb64, or qb64b cannot be used with code and raw")
		}

		szg, err := tables.GetSizage(*config.Code)
		if err != nil {
			return err
		}

		cs := szg.Hs + szg.Ss
		ms := szg.Ss - szg.Os

		if *config.Index > (1<<(6*ms) - 1) {
			return fmt.Errorf("invalid index %d for code %s", *config.Index, *config.Code)
		}

		if config.Ondex != nil {
			if szg.Os > 0 && *config.Ondex > (1<<(6*szg.Os)-1) {
				return fmt.Errorf("invalid ondex %d for code %s", *config.Ondex, *config.Code)
			}
		}

		if validateCode(*config.Code, tables.ValidCurrentSigCodes) && config.Ondex != nil {
			return fmt.Errorf("non-nil ondex %d for code %s", *config.Ondex, *config.Code)
		}

		if validateCode(*config.Code, tables.ValidBothSigCodes) {
			ondex := types.Ondex(*config.Index)
			if config.Ondex == nil {
				config.Ondex = &ondex
			}
		} else if config.Ondex != nil {
			if uint32(*config.Ondex) != uint32(*config.Index) && szg.Os == 0 {
				return fmt.Errorf("non matching ondex %d and index %d for code %s", *config.Ondex, *config.Index, *config.Code)
			}
		}

		// compute fs from index
		fs := szg.Fs
		if fs == nil {
			if cs%4 != 0 {
				return fmt.Errorf("whole code size not multiple of 4 for variable length material. cs = %d", cs)
			}
			if szg.Os != 0 {
				return fmt.Errorf("non-zero other index size for variable length material. os = %d", szg.Os)
			}

			_fs := (uint32(*config.Index) * 4) + cs
			fs = &_fs
		}

		rize := (*fs - cs) * 3 / 4
		if len(*config.Raw) < int(rize) {
			return fmt.Errorf("insufficient raw material: raw size = %d, rize = %d", len(*config.Raw), rize)
		}

		i.SetCode(*config.Code)
		i.SetRaw(types.Raw((*config.Raw)[:rize]))
		i.SetIndex(*config.Index)
		if config.Ondex != nil {
			i.SetOndex(*config.Ondex)
		}

		return nil
	}

	if config.Qb2 != nil {
		if config.Code != nil || config.Raw != nil || config.Qb64 != nil || config.Qb64b != nil {
			return fmt.Errorf("qb2 cannot be used with code, raw, qb64, or qb64b")
		}

		return ibexfil(i, *config.Qb2)
	}

	if config.Qb64 != nil {
		if config.Code != nil || config.Raw != nil || config.Qb2 != nil || config.Qb64b != nil {
			return fmt.Errorf("qb64 cannot be used with code, raw, qb2, or qb64b")
		}

		return iexfil(i, *config.Qb64)
	}

	if config.Qb64b != nil {
		if config.Code != nil || config.Raw != nil || config.Qb2 != nil || config.Qb64 != nil {
			return fmt.Errorf("qb64b cannot be used with code, raw, qb2, or qb64")
		}

		return iexfil(i, types.Qb64(*config.Qb64b))
	}

	return fmt.Errorf("no inputs provided")
}
