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

func NewMap() Map {
	return Map(*orderedmap.New[string, any]())
}

func (m Map) _map() orderedmap.OrderedMap[string, any] {
	return orderedmap.OrderedMap[string, any](m)
}

func (m Map) Set(key string, value any) (any, bool) {
	om := m._map()
	return om.Set(key, value)
}

func (m Map) Get(key string) (any, bool) {
	om := m._map()
	return om.Get(key)
}

func (m Map) Delete(key string) (any, bool) {
	om := m._map()
	return om.Delete(key)
}

func (m Map) MarshalJSON() ([]byte, error) {
	om := m._map()
	return om.MarshalJSON()
}

func (m Map) UnmarshalJSON(data []byte) error {
	om := m._map()
	return om.UnmarshalJSON(data)
}
