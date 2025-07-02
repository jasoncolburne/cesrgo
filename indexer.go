package cesrgo

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"slices"
	"strings"

	codex "github.com/jasoncolburne/cesrgo/indexer"
	"github.com/jasoncolburne/cesrgo/indexer/options"
	"github.com/jasoncolburne/cesrgo/types"
	"github.com/jasoncolburne/cesrgo/util"
)

type indexer struct {
	code  types.Code
	raw   types.Raw
	index types.Index
	ondex *types.Ondex
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

func (i *indexer) SetOndex(ondex *types.Ondex) {
	i.ondex = ondex
}

func (i *indexer) GetOndex() *types.Ondex {
	return i.ondex
}

func (i *indexer) Qb2() (types.Qb2, error) {
	return ibinfil(i)
}

func (i *indexer) Qb64() (types.Qb64, error) {
	return iinfil(i)
}

func (i *indexer) Qb64b() (types.Qb64b, error) {
	qb64, err := iinfil(i)
	if err != nil {
		return types.Qb64b{}, err
	}

	return types.Qb64b(qb64), nil
}

func ibexfil(i types.Indexer, qb2 types.Qb2) error {
	if len(qb2) == 0 {
		return fmt.Errorf("qb2 is empty")
	}

	first, err := util.NabSextets(qb2, 1)
	if err != nil {
		return err
	}

	hs, ok := codex.Bardage(first[0])
	if !ok {
		return fmt.Errorf("unknown bard: %x", first[0])
	}

	bhs := int(math.Ceil(float64(hs) * 3 / 4))
	if len(qb2) < bhs {
		return fmt.Errorf("insufficient material for hard part of code: qb2 size = %d, bhs = %d", len(qb2), bhs)
	}

	hard, err := util.CodeB2ToB64(qb2, hs)
	if err != nil {
		return err
	}

	szg, ok := codex.Sizes[types.Code(hard)]
	if !ok {
		return fmt.Errorf("unknown sizage: %s", hard)
	}
	cs := szg.Hs + szg.Ss
	ms := szg.Ss - szg.Os
	bcs := ((cs + 1) * 3) / 4

	if len(qb2) < int(bcs) {
		return fmt.Errorf("insufficient material for code: qb2 size = %d, bcs = %d", len(qb2), bcs)
	}

	both, err := util.CodeB2ToB64(qb2, int(cs))
	if err != nil {
		return err
	}

	index, err := util.B64ToU32(both[hs : hs+int(ms)])
	if err != nil {
		return err
	}

	var ondex *types.Ondex
	if util.ValidateCode(types.Code(hard), codex.ValidCurrentSigCodes) {
		if szg.Os != 0 {
			odx, err := util.B64ToU32(both[hs+int(ms) : hs+int(ms)+int(szg.Os)])
			if err != nil {
				return err
			}

			if odx != 0 {
				return fmt.Errorf("invalid ondex = '%d' for code = '%s'", odx, hard)
			}
		}
	} else if szg.Os != 0 {
		odx, err := util.B64ToU32(both[hs+int(ms) : hs+int(ms)+int(szg.Os)])
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
		pi := bytes[0]
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
		i.SetOndex(ondex)
	}

	return nil
}

func ibinfil(i types.Indexer) (types.Qb2, error) {
	code := i.GetCode()
	index := i.GetIndex()
	ondex := i.GetOndex()
	raw := i.GetRaw()

	ps := (3 - (len(raw) % 3)) % 3
	szg, ok := codex.Sizes[code]
	if !ok {
		return types.Qb2{}, fmt.Errorf("no sizage for code: %s", code)
	}

	cs := szg.Hs + szg.Ss
	ms := szg.Ss - szg.Os

	if index > (1<<(6*szg.Ss) - 1) {
		return types.Qb2{}, fmt.Errorf("invalid index=%d for code=%s", index, code)
	}

	if ondex != nil && *ondex > (1<<(6*szg.Os)-1) {
		return types.Qb2{}, fmt.Errorf("invalid ondex=%d for os=%d and code=%s", ondex, szg.Os, code)
	}

	var fs uint32
	if szg.Fs == nil {
		if cs%4 != 0 {
			return types.Qb2{}, fmt.Errorf("whole code size not multiple of 4 for variable length material. cs = %d", cs)
		}
		if szg.Os != 0 {
			return types.Qb2{}, fmt.Errorf("non-zero other index size for variable length material. os = %d", szg.Os)
		}
		fs = (uint32(index) * 4) + cs
	} else {
		fs = *szg.Fs
	}

	odx := 0
	if ondex != nil {
		odx = int(*ondex)
	}

	indexB64, err := util.IntToB64(int(index), int(ms))
	if err != nil {
		return types.Qb2{}, err
	}

	ondexB64, err := util.IntToB64(odx, int(szg.Os))
	if err != nil {
		return types.Qb2{}, err
	}

	both := string(code) + indexB64 + ondexB64

	if len(both) != int(cs) {
		return types.Qb2{}, fmt.Errorf("mismatch code size = %d with table = %d", cs, len(both))
	}

	if (int(cs) % 4) != ps-int(szg.Ls) {
		return types.Qb2{}, fmt.Errorf("invalid code=%s for converted raw pad size=%d", both, ps)
	}

	bothU16, err := util.B64ToU16(both)
	if err != nil {
		return types.Qb2{}, err
	}

	bcode := binary.BigEndian.AppendUint16([]byte{}, bothU16<<(2*(ps-int(szg.Ls))))
	full := make([]byte, len(bcode)+int(szg.Ls)+len(raw))

	copy(full[:len(bcode)], bcode)
	copy(full[len(bcode):len(bcode)+int(szg.Ls)], slices.Repeat([]byte{0}, int(szg.Ls)))
	copy(full[len(bcode)+int(szg.Ls):], raw)

	bfs := len(full)
	if bfs%3 != 0 || bfs*4/3 != int(fs) {
		return types.Qb2{}, fmt.Errorf("invalid code=%s (%s) for raw size=%d", code, both, len(raw))
	}

	return types.Qb2(full), nil
}

func iinfil(i types.Indexer) (types.Qb64, error) {
	code := i.GetCode()
	index := i.GetIndex()
	ondex := i.GetOndex()
	raw := i.GetRaw()

	ps := (3 - (len(raw) % 3)) % 3
	szg, ok := codex.Sizes[code]
	if !ok {
		return "", fmt.Errorf("unknown sizage for code: %s", code)
	}

	cs := szg.Hs + szg.Ss
	ms := szg.Ss - szg.Os

	var fs uint32
	if szg.Fs == nil {
		if cs%4 != 0 {
			return "", fmt.Errorf("whole code size not multiple of 4 for variable length material. cs=%d", cs)
		}
		if szg.Os != 0 {
			return "", fmt.Errorf("non-zero other index size for variable length material. os=%d", szg.Os)
		}
		fs = (uint32(index) * 4) + cs
	} else {
		fs = *szg.Fs
	}

	if index > (1<<(6*ms) - 1) {
		return "", fmt.Errorf("invalid index=%d for code=%s", index, code)
	}

	if szg.Os != 0 && *ondex > (1<<(6*szg.Os)-1) {
		return "", fmt.Errorf("invalid ondex=%d for os=%d and code=%s", ondex, szg.Os, code)
	}

	// both is hard code + converted index + converted ondex
	odx := 0
	if ondex != nil {
		odx = int(*ondex)
	}

	indexB64, err := util.IntToB64(int(index), int(ms))
	if err != nil {
		return "", err
	}

	ondexB64, err := util.IntToB64(odx, int(szg.Os))
	if err != nil {
		return "", err
	}

	both := fmt.Sprintf("%s%s%s", code, indexB64, ondexB64)

	if len(both) != int(cs) {
		return "", fmt.Errorf("mismatch code size = %d with table = %d", cs, len(both))
	}

	if (int(cs) % 4) != ps-int(szg.Ls) {
		return "", fmt.Errorf("invalid code=%s for converted raw pad size=%d", both, ps)
	}
	bytes := make([]byte, ps+len(raw))

	copy(bytes[:ps], slices.Repeat([]byte{0}, ps))
	copy(bytes[ps:], raw)
	full := both + base64.URLEncoding.EncodeToString(bytes)[ps-int(szg.Ls):]

	if len(full) != int(fs) {
		return "", fmt.Errorf("invalid code=%s for raw size=%d", both, len(raw))
	}

	return types.Qb64(full), nil
}

func iexfil(i types.Indexer, qb64 types.Qb64) error {
	if len(qb64) == 0 {
		return fmt.Errorf("qb64 is empty")
	}

	first := qb64[0]
	hs, ok := codex.Hardage(first)
	if !ok {
		return fmt.Errorf("unknown hard: %x", first)
	}

	hard := qb64[:hs]
	szg, ok := codex.Sizes[types.Code(hard)]
	if !ok {
		return fmt.Errorf("unknown sizage: %s", hard)
	}

	cs := szg.Hs + szg.Ss
	ms := szg.Ss - szg.Os

	if len(qb64) < int(cs) {
		return fmt.Errorf("insufficient material for code: qb64 size = %d, cs = %d", len(qb64), cs)
	}

	indexB64 := qb64[hs : hs+int(ms)]
	index, err := util.B64ToU32(string(indexB64))
	if err != nil {
		return err
	}

	ondexB64 := qb64[hs+int(ms) : hs+int(ms)+int(szg.Os)]

	var ondex *types.Ondex
	if slices.Contains(codex.ValidCurrentSigCodes, types.Code(hard)) {
		if szg.Os != 0 {
			odx, err := util.B64ToU32(string(ondexB64))
			if err != nil {
				return err
			}

			if odx != 0 {
				return fmt.Errorf("invalid ondex = '%d' for code = '%s'", odx, hard)
			}

			_ondex := types.Ondex(odx)
			ondex = &_ondex
		}
	} else if szg.Os != 0 {
		odx, err := util.B64ToU32(string(ondexB64))
		if err != nil {
			return err
		}

		_ondex := types.Ondex(odx)
		ondex = &_ondex
	} else {
		_ondex := types.Ondex(index)
		ondex = &_ondex
	}

	var fs uint32
	if szg.Fs == nil {
		if cs%4 != 0 {
			return fmt.Errorf("code size not multiple of 4 for variable length material: cs = %d", cs)
		}

		if szg.Os != 0 {
			return fmt.Errorf("non-zero other index size for variable length material: os = %d", szg.Os)
		}

		fs = (index * 4) + cs
	} else {
		fs = *szg.Fs
	}

	if len(qb64) < int(fs) {
		return fmt.Errorf("insufficient material: qb64 size = %d, fs = %d", len(qb64), fs)
	}

	qb64 = qb64[:fs]

	ps := cs % 4
	var pbs uint32
	if ps != 0 {
		pbs = 2 * ps
	} else {
		pbs = 2 * szg.Ls
	}

	var raw types.Raw
	if ps != 0 {
		base := strings.Repeat("A", int(ps)) + string(qb64[cs:])
		paw, err := base64.URLEncoding.DecodeString(base)
		if err != nil {
			return err
		}

		pi := util.BytesToInt(paw[:ps])
		if pi&(1<<pbs-1) != 0 {
			return fmt.Errorf("non-zeroed pad bits: %x", pi&(1<<pbs-1))
		}

		raw = paw[ps:]
	} else {
		base := string(qb64[cs:])
		paw, err := base64.URLEncoding.DecodeString(base)
		if err != nil {
			return err
		}

		li := util.BytesToInt(paw[:szg.Ls])
		if li != 0 {
			return fmt.Errorf("non-zeroed lead byte: %x", li)
		}

		raw = paw[szg.Ls:]
	}

	if len(raw) != (len(qb64)-int(cs))*3/4 {
		return fmt.Errorf("improperly qualified material: qb64 = %s", qb64)
	}

	i.SetCode(types.Code(hard))
	i.SetIndex(types.Index(index))
	i.SetOndex(ondex)
	i.SetRaw(raw)

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

		szg, ok := codex.Sizes[*config.Code]
		if !ok {
			return fmt.Errorf("unknown sizage for code: %s", *config.Code)
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

		if util.ValidateCode(*config.Code, codex.ValidCurrentSigCodes) && config.Ondex != nil {
			return fmt.Errorf("non-nil ondex %d for code %s", *config.Ondex, *config.Code)
		}

		if util.ValidateCode(*config.Code, codex.ValidBothSigCodes) {
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
		i.SetRaw((*config.Raw)[:rize])
		i.SetIndex(*config.Index)
		i.SetOndex(config.Ondex)

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
