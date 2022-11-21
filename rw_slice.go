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

import "sync"

var (
	_ SafeArray[int] = (*RWSlice[int])(nil)
)

type RWSlice[V any] struct {
	lock *sync.RWMutex
	bm   []V
}

func NewRWSlice[V any]() *RWSlice[V] {
	return &RWSlice[V]{
		lock: new(sync.RWMutex),
		bm:   []V{},
	}
}

func (rw RWSlice[V]) Add(src ...V) {
	rw.bm = append(rw.bm, src...)
}

func (rw RWSlice[V]) Remove(index int) {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	if index == 0 {
		rw.bm = rw.bm[1:]
	} else if index == len(rw.bm)-1 {
		rw.bm = rw.bm[:index]
	} else {
		j := 0
		for i, v := range rw.bm {
			if i != index {
				rw.bm[j] = v
				j++
			}
		}
		rw.bm = rw.bm[:j]
	}
}

func (rw RWSlice[V]) Range(f func(index int, value V) bool) {
	rw.lock.RLock()
	defer rw.lock.RUnlock()
	for i, v := range rw.bm {
		if !f(i, v) {
			break
		}
	}
}

func (rw RWSlice[V]) Each(f func(index int, value V)) {
	rw.Range(func(i int, v1 V) bool {
		f(i, v1)
		return true
	})
}

func (rw RWSlice[V]) Data() []V {
	rw.lock.RLock()
	defer rw.lock.RUnlock()

	cp := make([]V, len(rw.bm))
	for i, v := range rw.bm {
		cp[i] = v
	}
	return cp
}

func (rw RWSlice[V]) Set(index int, v V) {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.bm[index] = v
}

func (rw RWSlice[V]) Get(index int) V {
	rw.lock.RLock()
	defer rw.lock.RUnlock()
	return rw.bm[index]
}

func (rw RWSlice[V]) Length() int {
	rw.lock.RLock()
	defer rw.lock.RUnlock()
	return len(rw.bm)
}

func (rw RWSlice[V]) Capacity() int {
	rw.lock.RLock()
	defer rw.lock.RUnlock()
	return cap(rw.bm)
}
