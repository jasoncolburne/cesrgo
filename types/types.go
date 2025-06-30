package types

import (
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type (
	Code string
	Size uint32
	Raw  []byte

	Proto   string
	Kind    string
	Version struct {
		Major uint32
		Minor uint32
	}

	Index uint32
	Ondex uint32

	Qb64  string
	Qb64b []byte
	Qb2   []byte

	Map  orderedmap.OrderedMap[string, any]
	List []any
)

func (m Map) Map() orderedmap.OrderedMap[string, any] {
	return orderedmap.OrderedMap[string, any](m)
}
