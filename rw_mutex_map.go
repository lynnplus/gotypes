/*
 * Copyright (c) 2022-2023 Lynn <lynnplus90@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gotypes

import (
	"github.com/lynnplus/gotypes/constraints"
	"sync"
)

var (
	_ SafeMap[int, int] = (*RWMutexMap[int, int])(nil)
	_ Enumerable[bool]  = (*RWMutexMap[int, bool])(nil)
)

// RWMutexMap implements the SafeMap[K,V] interface
type RWMutexMap[K constraints.Basic, V any] struct {
	lock *sync.RWMutex
	bm   map[K]V
}

func NewRWMutexMap[K constraints.Basic, V any]() *RWMutexMap[K, V] {
	return &RWMutexMap[K, V]{
		lock: new(sync.RWMutex),
		bm:   make(map[K]V),
	}
}

func (m *RWMutexMap[K, V]) Get(key K) V {
	val, _ := m.Load(key)
	return val
}

func (m *RWMutexMap[K, V]) Exist(key K) (ok bool) {
	_, ok = m.Load(key)
	return ok
}

func (m *RWMutexMap[K, V]) Store(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.bm[key] = value
}

func (m *RWMutexMap[K, V]) Load(key K) (value V, ok bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	val, ok := m.bm[key]
	return val, ok
}

func (m *RWMutexMap[K, V]) Range(f func(key K, value V) bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for k, v := range m.bm {
		if !f(k, v) {
			break
		}
	}
}

func (m *RWMutexMap[K, V]) Each(f func(key K, value V)) {
	m.Range(func(k1 K, v1 V) bool {
		f(k1, v1)
		return true
	})
}

func (m *RWMutexMap[K, V]) EachValue(f func(value V)) {
	m.Range(func(_ K, v1 V) bool {
		f(v1)
		return true
	})
}

func (m *RWMutexMap[K, V]) Keys() []K {
	m.lock.RLock()
	defer m.lock.RUnlock()
	r := make([]K, len(m.bm))
	index := 0
	for k := range m.bm {
		r[index] = k
		index++
	}
	return r
}

func (m *RWMutexMap[K, V]) Values() []V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	r := make([]V, len(m.bm))
	index := 0
	for _, v := range m.bm {
		r[index] = v
		index++
	}
	return r
}

func (m *RWMutexMap[K, V]) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.bm)
}

func (m *RWMutexMap[K, V]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, key)
}

func (m *RWMutexMap[K, V]) DeleteAll() {
	m.lock.Lock()
	defer m.lock.Unlock()
	for k := range m.bm {
		delete(m.bm, k)
	}
}

func (m *RWMutexMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	temp, ok := m.Load(key)
	if ok {
		return temp, true
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	temp, ok = m.bm[key]
	if ok {
		return temp, true
	}
	m.bm[key] = value
	return value, false
}

func (m *RWMutexMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	temp, ok := m.bm[key]
	delete(m.bm, key)
	return temp, ok
}

func (m *RWMutexMap[K, V]) Data() map[K]V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	r := make(map[K]V, len(m.bm))
	for k, v := range m.bm {
		r[k] = v
	}
	return r
}
