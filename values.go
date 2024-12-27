package mttools

import (
	"iter"
	"maps"
)

type Values struct {
	list map[string]any
}

func NewValues() Values {
	return Values{
		list: make(map[string]any, 0),
	}
}

func (m *Values) CopyFrom(another *Values) {
	//clear this one by creating new empty map
	m.list = make(map[string]any, 0)
	maps.Copy(m.list, another.list)
}

// Sets value
func (m *Values) Set(key string, v any) *Values {
	m.list[key] = v
	return m
}

// Gets value by key, returns ok=false if there is no such key.
func (m *Values) GetOk(key string) (v any, ok bool) {
	v, ok = m.list[key]
	return v, ok
}

// Gets value by key, returns nil if there is no such key.
// Use GetOk() if you need distinguish nil values from key absence.
func (m *Values) Get(key string) any {
	if v, ok := m.list[key]; ok {
		return v
	}

	return nil
}

// Returns true if value for key was set.
func (m *Values) Has(key string) bool {
	_, ok := m.list[key]
	return ok
}

func (m *Values) GetNamesIterator() iter.Seq[string] {
	return maps.Keys(m.list)
}
