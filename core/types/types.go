package types

import (
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type (
	Map  orderedmap.OrderedMap[string, any]
	List []any

	Code string
	Size uint32
	Raw  []byte

	Proto   string
	Kind    string
	Version struct {
		Major uint32
		Minor uint32
	}
	Versage struct {
		Proto Proto
		Pvrsn Version
		Gvrsn *Version
	}

	Index uint32
	Ondex uint32

	Count uint32
	Ilk   string
	Trait string
	Tier  string
	Field string

	DateTime string

	Qb64  string
	Qb64b []byte
	Qb2   []byte
)

func NewMap(keys []string, values []any) Map {
	pairs := []orderedmap.Pair[string, any]{}

	if len(keys) != len(values) {
		panic("keys and values must have the same length")
	}

	for i, key := range keys {
		value := values[i]
		pairs = append(pairs, orderedmap.Pair[string, any]{Key: key, Value: value})
	}

	return Map(*orderedmap.New[string, any](orderedmap.WithInitialData(pairs...)))
}

func (m Map) _map() orderedmap.OrderedMap[string, any] {
	return orderedmap.OrderedMap[string, any](m)
}

func (m Map) Clone() Map {
	newMap := NewMap(nil, nil)
	om := m._map()
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		value := pair.Value
		if toClone, ok := value.(Map); ok {
			value = toClone.Clone()
		}

		newMap.Set(pair.Key, value)
	}
	return newMap
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
