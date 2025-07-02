package cesrgo

import (
	"fmt"

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
	return types.Qb2{}, nil
}

func cinfil(c types.Counter) (types.Qb64, error) {
	return types.Qb64(""), nil
}

func cexfil(c types.Counter, qb64 types.Qb64) error {
	return nil
}

func cbexfil(c types.Counter, qb2 types.Qb2) error {
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
