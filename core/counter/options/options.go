package options

import (
	"github.com/jasoncolburne/cesrgo/core/types"
)

type CounterOptions struct {
	Code  *types.Code
	Count *types.Count
	Qb2   *types.Qb2
	Qb64  *types.Qb64
	Qb64b *types.Qb64b
}

type CounterOption func(options *CounterOptions)

func WithCode(code types.Code) CounterOption {
	return func(options *CounterOptions) {
		options.Code = &code
	}
}

func WithCount(count types.Count) CounterOption {
	return func(options *CounterOptions) {
		options.Count = &count
	}
}

func WithQb2(qb2 types.Qb2) CounterOption {
	return func(options *CounterOptions) {
		options.Qb2 = &qb2
	}
}

func WithQb64(qb64 types.Qb64) CounterOption {
	return func(options *CounterOptions) {
		options.Qb64 = &qb64
	}
}

func WithQb64b(qb64b types.Qb64b) CounterOption {
	return func(options *CounterOptions) {
		options.Qb64b = &qb64b
	}
}
