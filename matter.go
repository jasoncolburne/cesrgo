package cesrgo

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"slices"
	"strings"

	codex "github.com/jasoncolburne/cesrgo/matter"
	tables "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

const (
	Pad = "_"
)

type matter struct {
	code types.Code
	size types.Size
	raw  types.Raw
	soft *string
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

func (m *matter) Hard() string {
	return string(m.code)
}

func (m *matter) SetSoft(soft *string) {
	m.soft = soft
}

func (m *matter) GetSoft() string {
	if m.soft == nil {
		return ""
	}

	return *m.soft
}

func (m *matter) Both() (string, error) {
	szg, err := codex.GetSizage(m.GetCode())
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s%s", m.Hard(), strings.Repeat(Pad, int(szg.Xs)), m.GetSoft()), nil
}

func (m *matter) Qb2() (types.Qb2, error) {
	return mbinfil(m)
}

func (m *matter) Qb64() (types.Qb64, error) {
	qb64b, err := minfil(m)
	if err != nil {
		return types.Qb64(""), err
	}

	return types.Qb64(qb64b), nil
}

func (m *matter) Qb64b() (types.Qb64b, error) {
	return minfil(m)
}

func mbinfil(m types.Matter) (types.Qb2, error) {
	_ = m

	return types.Qb2{}, nil
}

func minfil(m types.Matter) (types.Qb64b, error) {
	code := m.GetCode()
	both, err := m.Both()
	if err != nil {
		return types.Qb64b{}, err
	}
	raw := m.GetRaw()
	rs := len(raw)
	szg, err := codex.GetSizage(code)
	if err != nil {
		return types.Qb64b{}, err
	}

	cs := szg.Hs + szg.Ss

	if int(cs) != len(both) {
		return types.Qb64b{}, fmt.Errorf("both length mismatch: cs = %d, both = '%s'", cs, both)
	}

	var full string
	if szg.Fs == nil {
		if (int(szg.Ls)+rs)%3 != 0 || cs%4 != 0 {
			return types.Qb64b{}, fmt.Errorf("invalid full code '%s' with variable size rs = %d", both, rs)
		}

		bytes := make([]byte, int(szg.Ls)+rs)

		copy(bytes[:int(szg.Ls)], slices.Repeat([]byte{0}, int(szg.Ls)))
		copy(bytes[int(szg.Ls):], raw)

		full = both + base64.URLEncoding.EncodeToString(bytes)
	} else {
		ps := (3 - ((rs + int(szg.Ls)) % 3)) % 3
		if ps != int(cs)%4 {
			return types.Qb64b{}, fmt.Errorf("invalid full code '%s' with fixed size rs = %d", both, rs)
		}

		bytes := make([]byte, ps+int(szg.Ls)+rs)

		copy(bytes[:ps+int(szg.Ls)], slices.Repeat([]byte{0}, ps+int(szg.Ls)))
		copy(bytes[ps+int(szg.Ls):], raw)

		full = both + base64.URLEncoding.EncodeToString(bytes)[ps:]
	}

	if (len(full)%4 != 0) || (szg.Fs != nil && len(full) != int(*szg.Fs)) {
		return types.Qb64b{}, fmt.Errorf("invalid full size given code '%s' with rs = %d", both, rs)
	}

	return types.Qb64b(full), nil
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

	first := qb64[:1]
	hs, err := tables.Hardage(string(first))
	if err != nil {
		return err
	}

	if len(qb64) < int(hs) {
		return fmt.Errorf("insufficient material for hard part of code: qb64 size = %d, hs = %d", len(qb64), hs)
	}

	hard := qb64[:hs]

	szg, err := tables.GetSizage(types.Code(hard))
	if err != nil {
		return err
	}

	cs := szg.Hs + szg.Ss
	soft := qb64[int(hs):int(hs+szg.Ss)]
	xtra := soft[:int(szg.Xs)]
	soft = soft[int(szg.Xs):]

	if string(xtra) != strings.Repeat(Pad, int(szg.Xs)) {
		return fmt.Errorf("invalid prepad extra material: xtra = %s", xtra)
	}

	var fs uint32
	if szg.Fs == nil {
		i, err := b64ToU32(string(soft))
		if err != nil {
			return err
		}
		fs = i*4 + cs
	} else {
		fs = *szg.Fs
	}

	if len(qb64) < int(fs) {
		return fmt.Errorf("insufficient material: qb64 size = %d, fs = %d", len(qb64), fs)
	}

	qb64 = qb64[:fs]

	ps := cs % 4                                                  // net prepad bytes to ensure 24 bit align when encodeB64
	base := strings.Repeat("A", int(ps)) + string(qb64[int(cs):]) // prepad ps 'A's to  B64 of (lead + raw)
	paw, err := base64.URLEncoding.DecodeString(base)             // now should have ps + ls leading sextexts of zeros
	if err != nil {
		return err
	}
	raw := paw[int(ps+szg.Ls):] // remove prepad midpat bytes to invert back to raw
	// ensure midpad bytes are zero
	bytes := make([]byte, 8)
	copy(bytes[int(8-ps-szg.Ls):], paw[:int(ps+szg.Ls)])
	pi := binary.BigEndian.Uint64(bytes)
	if pi != 0 {
		return fmt.Errorf("nonzero midpad bytes=0x%x", pi)
	}

	if len(raw) != ((len(qb64)-int(cs))*3/4)-int(szg.Ls) {
		return fmt.Errorf("improperly qualified material: qb64 = %s", qb64)
	}

	m.SetCode(types.Code(hard))
	if len(soft) > 0 {
		softStr := string(soft)
		m.SetSoft(&softStr)
	}
	m.SetRaw(types.Raw(raw))
	m.SetSize(types.Size(len(raw)))

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

		length := len(*options.Raw)
		if length > 1<<32-1 {
			return fmt.Errorf("size too large")
		}

		m.SetCode(*options.Code)
		m.SetRaw(*options.Raw)
		//nolint:gosec
		m.SetSize(types.Size(len(*options.Raw)))
		m.SetSoft(options.Soft)

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
