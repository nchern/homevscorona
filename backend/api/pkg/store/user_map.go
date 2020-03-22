// generated by go-codegen(https://github.com/nchern/go-codegen)
// You COULD edit this code it you really need it and know what are you doing

package store

import (
	"sync"

	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

type StringUserPtrMapVisitor func(string, *model.User) bool

type StringUserPtrMap interface {
	Each(visitor StringUserPtrMapVisitor)
	Get(key string) (v *model.User, found bool)
	Set(key string, val *model.User)
	Update(src map[string]*model.User) StringUserPtrMap
	Remove(key string) bool
	Clone() StringUserPtrMap
}

type baseStringUserPtrMap struct {
	_map map[string]*model.User
}

func NewStringUserPtrMap() StringUserPtrMap {
	res := &baseStringUserPtrMap{
		_map: map[string]*model.User{},
	}
	return res
}

func NewStringUserPtrMapSyncronized() StringUserPtrMap {
	return &syncStringUserPtrMap{
		inner: NewStringUserPtrMap(),
	}
}

func (m *baseStringUserPtrMap) Get(key string) (v *model.User, found bool) {
	v, found = m._map[key]
	return
}

func (m *baseStringUserPtrMap) Each(visitor StringUserPtrMapVisitor) {
	for k, v := range m._map {
		if !visitor(k, v) {
			return
		}
	}
}

func (m *baseStringUserPtrMap) Set(key string, val *model.User) {
	m._map[key] = val
}

func (m *baseStringUserPtrMap) Update(src map[string]*model.User) StringUserPtrMap {
	for k, v := range src {
		m._map[k] = v
	}
	return m
}

func (m *baseStringUserPtrMap) Remove(key string) bool {
	_, found := m._map[key]
	delete(m._map, key)

	return found
}

func (m *baseStringUserPtrMap) Clone() StringUserPtrMap {
	res := NewStringUserPtrMap()
	for k, v := range m._map {
		res.Set(k, v)
	}

	return res
}

type syncStringUserPtrMap struct {
	inner StringUserPtrMap

	mutex sync.RWMutex
}

func (m *syncStringUserPtrMap) Each(visitor StringUserPtrMapVisitor) {
	m.mutex.RLock()
	m.inner.Each(visitor)
	m.mutex.RUnlock()
}

func (m *syncStringUserPtrMap) Get(key string) (v *model.User, found bool) {
	m.mutex.RLock()
	v, found = m.inner.Get(key)
	m.mutex.RUnlock()
	return
}

func (m *syncStringUserPtrMap) Set(key string, val *model.User) {
	m.mutex.Lock()
	m.inner.Set(key, val)
	m.mutex.Unlock()
}

func (m *syncStringUserPtrMap) Update(src map[string]*model.User) StringUserPtrMap {
	m.mutex.Lock()
	m.inner.Update(src)
	m.mutex.Unlock()

	return m
}

func (m *syncStringUserPtrMap) Remove(key string) bool {
	m.mutex.Lock()
	found := m.inner.Remove(key)
	m.mutex.Unlock()

	return found
}

func (m *syncStringUserPtrMap) Clone() StringUserPtrMap {
	m.mutex.RLock()
	r := m.inner.Clone()
	m.mutex.RUnlock()
	return r
}
