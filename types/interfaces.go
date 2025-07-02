package types

type (
	Matter interface {
		SetCode(code Code)
		GetCode() Code

		SetRaw(raw Raw)
		GetRaw() Raw

		SetSize(size Size)
		GetSize() Size

		Hard() string
		SetSoft(soft *string)
		GetSoft() string
		Both() (string, error)

		Qb2() (Qb2, error)
		Qb64() (Qb64, error)
		Qb64b() (Qb64b, error)
	}

	Indexer interface {
		SetCode(code Code)
		GetCode() Code

		SetRaw(raw Raw)
		GetRaw() Raw

		SetIndex(index Index)
		GetIndex() Index

		SetOndex(ondex *Ondex)
		GetOndex() *Ondex

		Qb2() (Qb2, error)
		Qb64() (Qb64, error)
		Qb64b() (Qb64b, error)
	}

	Counter interface {
		SetCode(code Code)
		GetCode() Code

		SetRaw(raw Raw)
		GetRaw() Raw

		SetCount(count Count)
		GetCount() Count

		Qb2() (Qb2, error)
		Qb64() (Qb64, error)
		Qb64b() (Qb64b, error)
	}
)
