package mttools

import (
	"iter"

	"github.com/elliotchance/orderedmap/v3"
)

// Ordered, concurrency safe list of key->value elements
type Values struct {
	Lockable

	list *orderedmap.OrderedMap[string, any]
}

func NewValues() Values {
	return Values{
		list: orderedmap.NewOrderedMap[string, any](),
	}
}

// Values with initial capacity given
func NewValuesC(capacity int) Values {
	return Values{
		list: orderedmap.NewOrderedMapWithCapacity[string, any](capacity),
	}
}

func (m *Values) CopyFrom(another *Values) {
	m.Lock()
	m.list = another.list.Copy()
	m.Unlock()
}

// Sets value
func (m *Values) Set(key any, v any) *Values {
	m.Lock()
	m.list.Set(AnyToString(key), v)
	m.Unlock()
	return m
}

// Gets value by key, returns ok=false if there is no such key.
func (m *Values) GetOk(key any) (v any, ok bool) {
	m.Lock()
	v, ok = m.list.Get(AnyToString(key))
	m.Unlock()
	return v, ok
}

// Gets value by key, returns nil if there is no such key.
// Use GetOk() if you need distinguish nil values from key absence.
func (m *Values) Get(key any) any {
	if v, ok := m.GetOk(key); ok {
		return v
	}

	return nil
}

// Returns true if value for key was set.
func (m *Values) Has(key any) bool {
	return m.list.Has(AnyToString(key))
}

func (m *Values) GetNamesIterator() iter.Seq[string] {
	return m.list.Keys()
}

func (m *Values) GetIterator() iter.Seq2[string, any] {
	return m.list.AllFromFront()
}
