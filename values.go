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

func (m *Values) CopyFrom(another Values) {
	//clear by creating new empty map
	m.list = make(map[string]any, 0)
	maps.Copy(m.list, another.list)
}

func (m *Values) Set(key string, v any) *Values {
	m.list[key] = v
	return m
}

func (m *Values) GetOk(key string) (v any, ok bool) {
	v, ok = m.list[key]
	return v, ok
}

func (m *Values) Get(key string) any {
	if v, ok := m.GetOk(key); ok {
		return v
	}

	return nil
}

func (m *Values) GetNamesIterator() iter.Seq[string] {
	return maps.Keys(m.list)
}
