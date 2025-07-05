package options

import "github.com/jasoncolburne/cesrgo/core/types"

type IndexerOptions struct {
	Code  *types.Code
	Raw   *types.Raw
	Index *types.Index
	Ondex *types.Ondex
	Qb2   *types.Qb2
	Qb64  *types.Qb64
	Qb64b *types.Qb64b
}

type IndexerOption func(options *IndexerOptions)

func WithCode(code types.Code) IndexerOption {
	return func(options *IndexerOptions) {
		options.Code = &code
	}
}

func WithRaw(raw types.Raw) IndexerOption {
	return func(options *IndexerOptions) {
		options.Raw = &raw
	}
}

func WithIndex(index types.Index) IndexerOption {
	return func(options *IndexerOptions) {
		options.Index = &index
	}
}

func WithOndex(ondex types.Ondex) IndexerOption {
	return func(options *IndexerOptions) {
		options.Ondex = &ondex
	}
}

func WithQb2(qb2 types.Qb2) IndexerOption {
	return func(options *IndexerOptions) {
		options.Qb2 = &qb2
	}
}

func WithQb64(qb64 types.Qb64) IndexerOption {
	return func(options *IndexerOptions) {
		options.Qb64 = &qb64
	}
}

func WithQb64b(qb64b types.Qb64b) IndexerOption {
	return func(options *IndexerOptions) {
		options.Qb64b = &qb64b
	}
}
