package types

type (
	Matter interface {
		SetCode(code Code)
		GetCode() Code

		SetRaw(raw Raw)
		GetRaw() Raw

		SetSize(size Size)
		GetSize() Size

		Qb2() Qb2
		Qb64() Qb64
		Qb64b() Qb64b
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

		Qb2() Qb2
		Qb64() Qb64
		Qb64b() Qb64b
	}
)
