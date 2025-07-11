package options

import "github.com/jasoncolburne/cesrgo/core/types"

type MatterOptions struct {
	Code  *types.Code
	Soft  *string
	Raw   *types.Raw
	Qb2   *types.Qb2
	Qb64  *types.Qb64
	Qb64b *types.Qb64b
}

type MatterOption func(options *MatterOptions)

func WithCode(code types.Code) MatterOption {
	return func(options *MatterOptions) {
		options.Code = &code
	}
}

func WithSoft(soft string) MatterOption {
	return func(options *MatterOptions) {
		options.Soft = &soft
	}
}

func WithRaw(raw types.Raw) MatterOption {
	return func(options *MatterOptions) {
		options.Raw = &raw
	}
}

func WithQb2(qb2 types.Qb2) MatterOption {
	return func(options *MatterOptions) {
		options.Qb2 = &qb2
	}
}

func WithQb64(qb64 types.Qb64) MatterOption {
	return func(options *MatterOptions) {
		options.Qb64 = &qb64
	}
}

func WithQb64b(qb64b types.Qb64b) MatterOption {
	return func(options *MatterOptions) {
		options.Qb64b = &qb64b
	}
}
