package cesrgo

import (
	"fmt"
	"math"

	"github.com/jasoncolburne/cesrgo/common"
	"github.com/jasoncolburne/cesrgo/common/util"
	codex "github.com/jasoncolburne/cesrgo/counter"
	"github.com/jasoncolburne/cesrgo/counter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type counter struct {
	code  types.Code
	raw   types.Raw
	count types.Count
}

func (c *counter) SetCode(code types.Code) {
	c.code = code
}

func (c *counter) GetCode() types.Code {
	return c.code
}

func (c *counter) SetRaw(raw types.Raw) {
	c.raw = raw
}

func (c *counter) GetRaw() types.Raw {
	return c.raw
}

func (c *counter) SetCount(count types.Count) {
	c.count = count
}

func (c *counter) GetCount() types.Count {
	return c.count
}

func (c *counter) Qb2() (types.Qb2, error) {
	return cbinfil(c)
}

func (c *counter) Qb64() (types.Qb64, error) {
	return cinfil(c)
}

func (c *counter) Qb64b() (types.Qb64b, error) {
	qb64, err := cinfil(c)
	if err != nil {
		return types.Qb64b{}, err
	}

	return types.Qb64b(qb64), nil
}

func cbinfil(c types.Counter) (types.Qb2, error) {
	qb64, err := cinfil(c)
	if err != nil {
		return types.Qb2{}, err
	}

	b2, err := util.CodeB64ToB2(string(qb64))
	if err != nil {
		return types.Qb2{}, err
	}

	return types.Qb2(b2), nil
}

func cinfil(c types.Counter) (types.Qb64, error) {
	code := c.GetCode()
	count := c.GetCount()

	szg, ok := codex.Sizes[common.VERSION_2_0.Major][code]
	if !ok {
		return types.Qb64(""), fmt.Errorf("unknown code: %s", code)
	}

	if count < 0 || count > (1<<(6*szg.Ss)-1) {
		return types.Qb64(""), fmt.Errorf("invalid count=%d for code=%s", count, code)
	}

	countB64, err := util.IntToB64(int(count), int(szg.Ss))
	if err != nil {
		return types.Qb64(""), err
	}

	both := fmt.Sprintf("%s%s", code, countB64)

	if len(both)%4 != 0 {
		return types.Qb64(""), fmt.Errorf("invalid size=%d of %s not a multiple of 4", len(both), both)
	}

	return types.Qb64(both), nil
}

func cexfil(c types.Counter, qb64 types.Qb64) error {
	if len(qb64) < 2 {
		return fmt.Errorf("empty material, need more characters")
	}

	first := string(qb64[:2])
	if first[0] == '_' {
		return fmt.Errorf("unexpected op code start while extracting Counter.")
	}

	hs, ok := codex.Hards[first]
	if !ok {
		return fmt.Errorf("uneunsupported code start first=%s", first)
	}

	if len(qb64) < hs {
		return fmt.Errorf("need more characters")
	}

	hard := qb64[:hs]

	szg, ok := codex.Sizes[common.VERSION_2_0.Major][types.Code(hard)]
	if !ok {
		return fmt.Errorf("unsupported code=%s", hard)
	}

	if len(qb64) < int(szg.Fs) {
		return fmt.Errorf("need more characters")
	}

	count := qb64[hs:szg.Fs]
	countInt, err := util.B64ToU32(string(count))
	if err != nil {
		return err
	}

	c.SetCode(types.Code(hard))
	c.SetCount(types.Count(countInt))

	return nil
}

func cbexfil(c types.Counter, qb2 types.Qb2) error {
	if len(qb2) < 2 {
		return fmt.Errorf("empty material, need more bytes")
	}

	first, err := util.NabSextets(qb2, 2)
	if err != nil {
		return err
	}

	if first[0] == 0xfc {
		return fmt.Errorf("unexpected op code start while extracting Counter.")
	}

	hs, ok := codex.Bards[string(first)]
	if !ok {
		return fmt.Errorf("unsupported code start sextet=%s", first)
	}

	bhs := int(math.Ceil(float64(hs) * 3 / 4))
	if len(qb2) < bhs {
		return fmt.Errorf("need more bytes")
	}

	hard, err := util.CodeB2ToB64(qb2, int(hs))
	if err != nil {
		return err
	}

	szg, ok := codex.Sizes[common.VERSION_2_0.Major][types.Code(hard)]
	if !ok {
		return fmt.Errorf("unsupported code=%s", hard)
	}

	bcs := int(math.Ceil(float64(szg.Fs) * 3 / 4))
	if len(qb2) < bcs {
		return fmt.Errorf("need more bytes")
	}

	both, err := util.CodeB2ToB64(qb2, int(szg.Fs))
	if err != nil {
		return err
	}

	count, err := util.B64ToU32(string(both[hs:szg.Fs]))
	if err != nil {
		return err
	}

	c.SetCode(types.Code(hard))
	c.SetCount(types.Count(count))

	return nil
}

func NewCounter(c types.Counter, opts ...options.CounterOption) error {
	config := &options.CounterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	if config.Code != nil && config.Raw != nil && config.Count != nil {
		if config.Qb2 != nil || config.Qb64 != nil || config.Qb64b != nil {
			return fmt.Errorf("qb2, qb64, or qb64b cannot be used with code, raw, and count")
		}

		c.SetCode(*config.Code)
		c.SetRaw(*config.Raw)
		c.SetCount(*config.Count)
	}

	if config.Qb2 != nil {
		if config.Code != nil || config.Raw != nil || config.Count != nil || config.Qb64 != nil || config.Qb64b != nil {
			return fmt.Errorf("qb2 cannot be used with code, raw, count, qb64, or qb64b")
		}

		return cbexfil(c, *config.Qb2)
	}

	if config.Qb64 != nil {
		if config.Code != nil || config.Raw != nil || config.Count != nil || config.Qb2 != nil || config.Qb64b != nil {
			return fmt.Errorf("qb64 cannot be used with code, raw, count, qb2, or qb64b")
		}

		return cexfil(c, *config.Qb64)
	}

	if config.Qb64b != nil {
		if config.Code != nil || config.Raw != nil || config.Count != nil || config.Qb2 != nil || config.Qb64 != nil {
			return fmt.Errorf("qb64b cannot be used with code, raw, count, qb2, or qb64")
		}

		return cexfil(c, types.Qb64(*config.Qb64b))
	}

	return fmt.Errorf("no valid options provided")
}
