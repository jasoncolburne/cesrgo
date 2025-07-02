package cesrgo

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"slices"
	"strings"

	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
	"github.com/jasoncolburne/cesrgo/util"
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
	szg, ok := codex.Sizes[m.GetCode()]
	if !ok {
		return "", fmt.Errorf("unknown code: %s", m.GetCode())
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
	code := m.GetCode()
	both, err := m.Both()
	if err != nil {
		return types.Qb2{}, err
	}
	raw := m.GetRaw()

	szg, ok := codex.Sizes[code]
	if !ok {
		return types.Qb2{}, fmt.Errorf("unknown code: %s", code)
	}
	cs := szg.Hs + szg.Ss

	n := int(math.Ceil(float64(cs) * 3 / 4))

	i, err := util.B64ToU64(both)
	if err != nil {
		return types.Qb2{}, err
	}

	shifted := i << (2 * (cs % 4))
	bcode := make([]byte, 8)
	binary.BigEndian.PutUint64(bcode, shifted)
	bcode = bcode[8-n:]

	full := make([]byte, len(bcode)+int(szg.Ls)+len(raw))
	copy(full[:len(bcode)], bcode)
	copy(full[len(bcode):len(bcode)+int(szg.Ls)], slices.Repeat([]byte{0}, int(szg.Ls)))
	copy(full[len(bcode)+int(szg.Ls):], raw)

	bfs := len(full)
	var fs uint32
	if szg.Fs == nil {
		i := int(szg.Hs+szg.Ss) + (len(raw)+int(szg.Ls))*4/3
		if i > 1<<32-1 {
			return types.Qb2{}, fmt.Errorf("size too large")
		}

		//nolint:gosec
		fs = uint32(i)
	} else {
		fs = *szg.Fs
	}

	if bfs%3 != 0 || (bfs*4/3) != int(fs) {
		return types.Qb2{}, fmt.Errorf("invalid full code '%s' with raw size %d (bfs = %d, fs = %d)", both, len(raw), bfs, fs)
	}

	return types.Qb2(full), nil
}

func minfil(m types.Matter) (types.Qb64b, error) {
	code := m.GetCode()
	both, err := m.Both()
	if err != nil {
		return types.Qb64b{}, err
	}
	raw := m.GetRaw()
	rs := len(raw)
	szg, ok := codex.Sizes[code]
	if !ok {
		return types.Qb64b{}, fmt.Errorf("unknown code: %s", code)
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

	sextets, err := util.NabSextets(qb2, 1)
	if err != nil {
		return err
	}

	first := sextets[0]
	hs, ok := codex.Bardage(first)
	if !ok {
		return fmt.Errorf("unknown bard: %x", first)
	}

	bhs := int(math.Ceil(float64(hs) * 3 / 4))
	if len(qb2) < bhs {
		return fmt.Errorf("insufficient material for hard part of code: qb2 size = %d, bhs = %d", len(qb2), bhs)
	}

	hard, err := util.CodeB2ToB64(qb2, int(hs))
	if err != nil {
		return err
	}

	szg, ok := codex.Sizes[types.Code(hard)]
	if !ok {
		return fmt.Errorf("unknown hard: %s", hard)
	}

	cs := szg.Hs + szg.Ss

	bcs := int(math.Ceil(float64(cs) * 3 / 4))
	if len(qb2) < bcs {
		return fmt.Errorf("insufficient material: qb2 size = %d, bcs = %d", len(qb2), bcs)
	}

	both, err := util.CodeB2ToB64(qb2, int(cs))
	if err != nil {
		return err
	}

	soft := both[int(szg.Hs):int(szg.Hs+szg.Ss)]
	xtra := soft[:int(szg.Xs)]
	soft = soft[int(szg.Xs):]

	if xtra != strings.Repeat(Pad, int(szg.Xs)) {
		return fmt.Errorf("invalid prepad extra material: xtra = %s", xtra)
	}

	var fs uint32
	var size uint32
	if szg.Fs == nil {
		if len(qb2) < bcs {
			return fmt.Errorf("insufficient material for code: qb2 size = %d, bcs = %d", len(qb2), bcs)
		}

		i, err := util.B64ToU32(soft)
		if err != nil {
			return err
		}
		fs = i*4 + cs
	} else {
		fs = *szg.Fs
	}

	bfs := int(math.Ceil((float64(fs) * 3) / 4))
	if len(qb2) < bfs {
		return fmt.Errorf("insufficient material: qb2 size = %d, bfs = %d", len(qb2), bfs)
	}

	qb2 = qb2[:bfs]

	ps := cs % 4
	pbs := 2 * ps

	pi := int(qb2[bcs-1 : bcs][0])
	pi &= 2<<pbs - 1
	if pi != 0 {
		return fmt.Errorf("non-zeroed code midpad bits")
	}

	li := util.BytesToInt(qb2[bcs : bcs+int(szg.Ls)])
	if li != 0 {
		return fmt.Errorf("non-zeroed lead midpad bytes")
	}

	raw := qb2[bcs+int(szg.Ls):]

	m.SetCode(types.Code(hard))
	m.SetSize(types.Size(size))
	m.SetRaw(types.Raw(raw))
	if soft != "" {
		m.SetSoft(&soft)
	}

	return nil
}

func mexfil(m types.Matter, qb64 types.Qb64) error {
	if len(qb64) == 0 {
		return fmt.Errorf("qb64 is empty")
	}

	first := qb64[0]
	hs, ok := codex.Hardage(first)
	if !ok {
		return fmt.Errorf("unknown hard: %x", first)
	}

	if len(qb64) < int(hs) {
		return fmt.Errorf("insufficient material for hard part of code: qb64 size = %d, hs = %d", len(qb64), hs)
	}

	hard := qb64[:hs]

	szg, ok := codex.Sizes[types.Code(hard)]
	if !ok {
		return fmt.Errorf("unknown code: %s", hard)
	}

	cs := szg.Hs + szg.Ss
	soft := qb64[hs : hs+int(szg.Ss)]
	xtra := soft[:szg.Xs]
	soft = soft[szg.Xs:]

	if string(xtra) != strings.Repeat(Pad, int(szg.Xs)) {
		return fmt.Errorf("invalid prepad extra material: xtra = %s", xtra)
	}

	var fs uint32
	if szg.Fs == nil {
		i, err := util.B64ToU32(string(soft))
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

	ps := cs % 4
	base := strings.Repeat("A", int(ps)) + string(qb64[int(cs):])
	paw, err := base64.URLEncoding.DecodeString(base)
	if err != nil {
		return err
	}
	raw := paw[int(ps+szg.Ls):]

	// ensure midpad bytes are zero
	pi := util.BytesToInt(paw[:int(ps+szg.Ls)])
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

	length := len(raw)
	if length > 1<<32-1 {
		return fmt.Errorf("size too large")
	}

	m.SetSize(types.Size(length))

	return nil
}

func NewMatter(m types.Matter, opts ...options.MatterOption) error {
	config := &options.MatterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	if config.Code != nil && config.Raw != nil {
		if config.Qb2 != nil || config.Qb64 != nil || config.Qb64b != nil {
			return fmt.Errorf("code and raw cannot be used with qb2, qb64, or qb64b")
		}

		length := len(*config.Raw)
		if length > 1<<32-1 {
			return fmt.Errorf("size too large")
		}

		m.SetCode(*config.Code)
		m.SetRaw(*config.Raw)
		//nolint:gosec
		m.SetSize(types.Size(len(*config.Raw)))
		m.SetSoft(config.Soft)

		return nil
	}

	if config.Qb2 != nil {
		if config.Code != nil || config.Raw != nil || config.Qb64 != nil || config.Qb64b != nil {
			return fmt.Errorf("qb2 cannot be used with code, raw, qb64, or qb64b")
		}

		return mbexfil(m, *config.Qb2)
	}

	if config.Qb64 != nil {
		if config.Code != nil || config.Raw != nil || config.Qb2 != nil || config.Qb64b != nil {
			return fmt.Errorf("qb64 cannot be used with code, raw, qb2, or qb64b")
		}

		return mexfil(m, *config.Qb64)
	}

	if config.Qb64b != nil {
		if config.Code != nil || config.Raw != nil || config.Qb2 != nil || config.Qb64 != nil {
			return fmt.Errorf("qb64b cannot be used with code, raw, qb2, or qb64")
		}

		return mexfil(m, types.Qb64(*config.Qb64b))
	}

	return fmt.Errorf("no inputs provided")
}
