/*
 * Copyright (c) 2022 Lynn <lynnplus90@gmail.com>
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
	"sync/atomic"
)

var (
	_ SafeMap[int, int] = (*SyncMap[int, int])(nil)
	_ Enumerable[bool]  = (*SyncMap[int, bool])(nil)
)

// SyncMap efficiency is lower than the sync.Map, only suitable for small maps
type SyncMap[K constraints.Basic, V any] struct {
	instance          *sync.Map
	cacheSize         atomic.Int64
	unknownStoreCount atomic.Int64
}

func NewSyncMap[K constraints.Basic, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		instance: &sync.Map{},
	}
}

func (s *SyncMap[K, V]) Get(key K) V {
	temp, _ := s.Load(key)
	return temp
}

func (s *SyncMap[K, V]) Exist(key K) (ok bool) {
	_, ok = s.Load(key)
	return ok
}

func (s *SyncMap[K, V]) Store(key K, value V) {
	if s.verifyNil(value) {
		return
	}
	s.instance.Store(key, value)
	s.unknownStoreCount.Add(1)
}

func (s *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := s.instance.Load(key)
	if !ok {
		return value, false
	}
	data := v.(V)
	return data, ok
}

func (s *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	s.instance.Range(func(k any, v any) bool {
		return f(k.(K), v.(V))
	})
}

func (s *SyncMap[K, V]) Each(f func(key K, value V)) {
	s.Range(func(k K, v V) bool {
		f(k, v)
		return true
	})
}

func (s *SyncMap[K, V]) EachValue(f func(value V)) {
	s.Range(func(k K, v V) bool {
		f(v)
		return true
	})
}

func (s *SyncMap[K, V]) Keys() []K {
	temp := s.Size()
	var keys []K
	if temp > 0 {
		keys = make([]K, 0, temp)
	}
	s.Each(func(k K, v V) {
		keys = append(keys, k)
	})
	return keys
}

func (s *SyncMap[K, V]) Values() []V {
	temp := s.Size()
	var vs []V
	if temp > 0 {
		vs = make([]V, 0, temp)
	}
	s.Each(func(k K, v V) {
		vs = append(vs, v)
	})
	return vs
}

// Size Does not represent the actual size of the current map, just an estimate
func (s *SyncMap[K, V]) Size() int {
	temp := int(s.cacheSize.Load())
	if temp < 0 || s.unknownStoreCount.Load() > int64(temp) {
		temp = s.syncSize()
	}
	return temp
}

func (s *SyncMap[K, V]) Delete(key K) {
	if _, ok := s.instance.LoadAndDelete(key); ok {
		s.cacheSize.Add(-1)
	}
}

func (s *SyncMap[K, V]) DeleteAll() {
	s.Each(func(k K, v V) {
		s.Delete(k)
	})
	s.cacheSize.Store(0)
	s.unknownStoreCount.Store(0)
}

func (s *SyncMap[K, V]) Data() map[K]V {
	temp := s.Size()
	var data map[K]V
	if temp > 0 {
		data = make(map[K]V, temp)
	} else {
		data = map[K]V{}
	}
	s.Each(func(k K, v V) {
		data[k] = v
	})
	return data
}

func (s *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	if s.verifyNil(value) {
		return actual, false
	}
	data, load := s.instance.LoadOrStore(key, value)
	if !load {
		s.cacheSize.Add(1)
	}
	actual = data.(V)
	return actual, load
}

func (s *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	data, exist := s.instance.LoadAndDelete(key)
	if exist {
		s.cacheSize.Add(-1)
	}
	if !exist {
		return nil, false
	}
	value = data.(V)
	return value, exist
}

func (s *SyncMap[K, V]) syncSize() int {
	count := 0
	s.instance.Range(func(key, value any) bool {
		count++
		return true
	})
	s.unknownStoreCount.Store(0)
	s.cacheSize.Store(int64(count))
	return count
}

func (s *SyncMap[K, V]) verifyNil(v V) bool {
	ptr := interface{}(v)
	return ptr == nil
}
