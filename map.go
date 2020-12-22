package syncmap

import (
	"sync"
)

type Map struct {
	m sync.Map
}

func (m *Map) Length() *int64 {
	var counter int64

	m.m.Range(func(key, value interface{}) bool {
		counter++

		return true
	})

	return &counter
}

func (m *Map) Range(f func(key, value interface{}) bool) {
	m.m.Range(f)
}

func (m *Map) Delete(key interface{}) {
	m.m.Delete(key)
}

func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	value, ok = m.m.Load(key)

	return
}

func (m *Map) Store(key interface{}, value interface{}) {
	m.m.Store(key, value)
}
