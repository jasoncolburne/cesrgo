package cesrgo

import (
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

func NewCounter(c types.Counter, opts ...options.CounterOption) error {
	config := &options.CounterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	if config.Code != nil {
		c.SetCode(*config.Code)
	}

	if config.Raw != nil {
		c.SetRaw(*config.Raw)
	}

	if config.Count != nil {
		c.SetCount(*config.Count)
	}

	return nil
}
