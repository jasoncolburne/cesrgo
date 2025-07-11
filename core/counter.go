package cesr

import (
	"fmt"
	"math"
	"math/big"

	"github.com/jasoncolburne/cesrgo"
	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/counter"
	"github.com/jasoncolburne/cesrgo/core/counter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Counter struct {
	code  types.Code
	count types.Count
}

func (c *Counter) SetCode(code types.Code) {
	c.code = code
}

func (c *Counter) GetCode() types.Code {
	return c.code
}

func (c *Counter) SetCount(count types.Count) {
	c.count = count
}

func (c *Counter) GetCount() types.Count {
	return c.count
}

func (c *Counter) Qb2() (types.Qb2, error) {
	return cbinfil(c)
}

func (c *Counter) Qb64() (types.Qb64, error) {
	return cinfil(c)
}

func (c *Counter) Qb64b() (types.Qb64b, error) {
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

	b2, err := common.CodeB64ToB2(string(qb64))
	if err != nil {
		return types.Qb2{}, err
	}

	return types.Qb2(b2), nil
}

func cinfil(c types.Counter) (types.Qb64, error) {
	code := c.GetCode()
	count := c.GetCount()

	szg, ok := codex.Sizes[cesrgo.VERSION.Major][code]
	if !ok {
		return types.Qb64(""), fmt.Errorf("unknown code: %s", code)
	}

	if count > (1<<(6*szg.Ss) - 1) {
		return types.Qb64(""), fmt.Errorf("invalid count=%d for code=%s", count, code)
	}

	countB64, err := common.BigIntToB64(big.NewInt(int64(count)), int(szg.Ss))
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
		return fmt.Errorf("unexpected op code start while extracting counter")
	}

	hs, ok := codex.Hardage(first)
	if !ok {
		return fmt.Errorf("unsupported code start first=%s", first)
	}

	if len(qb64) < hs {
		return fmt.Errorf("need more characters")
	}

	hard := qb64[:hs]

	szg, ok := codex.Sizes[cesrgo.VERSION.Major][types.Code(hard)]
	if !ok {
		return fmt.Errorf("unsupported code=%s", hard)
	}

	if len(qb64) < int(szg.Fs) {
		return fmt.Errorf("need more characters")
	}

	countStr := string(qb64[hs:szg.Fs])
	// u64 is safe here because the maximum length is 5 b64 octets
	count, err := common.B64ToU64(countStr)
	if err != nil {
		return err
	}

	c.SetCode(types.Code(hard))
	//nolint:gosec
	c.SetCount(types.Count(count))

	return nil
}

func cbexfil(c types.Counter, qb2 types.Qb2) error {
	if len(qb2) < 2 {
		return fmt.Errorf("empty material, need more bytes")
	}

	first, err := common.NabSextets(qb2, 2)
	if err != nil {
		return err
	}

	if first[0] == 0xfc {
		return fmt.Errorf("unexpected op code start while extracting counter")
	}

	b := [2]byte{first[0], first[1]}
	hs, ok := codex.Bardage(b)
	if !ok {
		return fmt.Errorf("unsupported code start sextet=%s", first)
	}

	bhs := int(math.Ceil(float64(hs) * 3 / 4))
	if len(qb2) < bhs {
		return fmt.Errorf("need more bytes")
	}

	hard, err := common.CodeB2ToB64(qb2, hs)
	if err != nil {
		return err
	}

	szg, ok := codex.Sizes[cesrgo.VERSION.Major][types.Code(hard)]
	if !ok {
		return fmt.Errorf("unsupported code=%s", hard)
	}

	bcs := int(math.Ceil(float64(szg.Fs) * 3 / 4))
	if len(qb2) < bcs {
		return fmt.Errorf("need more bytes")
	}

	both, err := common.CodeB2ToB64(qb2, int(szg.Fs))
	if err != nil {
		return err
	}

	// u64 is safe here, max length is 5 b64 octets
	count, err := common.B64ToU64(both[hs:szg.Fs])
	if err != nil {
		return err
	}

	c.SetCode(types.Code(hard))
	//nolint:gosec
	c.SetCount(types.Count(count))

	return nil
}

func NewCounter(opts ...options.CounterOption) (*Counter, error) {
	config := &options.CounterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	c := &Counter{}

	if config.Code != nil && config.Count != nil {
		if config.Qb2 != nil || config.Qb64 != nil || config.Qb64b != nil {
			return c, fmt.Errorf("qb2, qb64, or qb64b cannot be used with code, raw, and count")
		}

		c.SetCode(*config.Code)
		c.SetCount(*config.Count)

		return c, nil
	}

	if config.Qb2 != nil {
		if config.Code != nil || config.Count != nil || config.Qb64 != nil || config.Qb64b != nil {
			return c, fmt.Errorf("qb2 cannot be used with code, raw, count, qb64, or qb64b")
		}

		if err := cbexfil(c, *config.Qb2); err != nil {
			return c, err
		}

		return c, nil
	}

	if config.Qb64 != nil {
		if config.Code != nil || config.Count != nil || config.Qb2 != nil || config.Qb64b != nil {
			return c, fmt.Errorf("qb64 cannot be used with code, raw, count, qb2, or qb64b")
		}

		if err := cexfil(c, *config.Qb64); err != nil {
			return c, err
		}

		return c, nil
	}

	if config.Qb64b != nil {
		if config.Code != nil || config.Count != nil || config.Qb2 != nil || config.Qb64 != nil {
			return c, fmt.Errorf("qb64b cannot be used with code, raw, count, qb2, or qb64")
		}

		if err := cexfil(c, types.Qb64(*config.Qb64b)); err != nil {
			return c, err
		}

		return c, nil
	}

	return c, fmt.Errorf("no valid options provided")
}
