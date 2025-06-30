package types

type (
	Matter interface {
		SetCode(code Code)
		GetCode() Code

		SetRaw(raw Raw)
		GetRaw() Raw

		SetSize(size Size)
		GetSize() Size
	}

	Indexer interface {
		SetCode(code Code)
		GetCode() Code

		SetRaw(raw Raw)
		GetRaw() Raw

		SetIndex(index Index)
		GetIndex() Index

		SetOndex(ondex Ondex)
		GetOndex() Ondex
	}
)
