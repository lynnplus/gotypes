/*
 * Copyright (c) 2023 Lynn <lynnplus90@gmail.com>
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

package gsync

import (
	"sync"
)

type Pool[T Resetter] struct {
	pool sync.Pool
}

func NewPool[T Resetter](newFunc func() T) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{New: func() any {
			if newFunc == nil {
				return nil
			}
			return newFunc()
		}},
	}
}

func (p *Pool[T]) Put(x T) {
	if any(x) == nil {
		return
	}
	x.Reset()
	p.pool.Put(x)
}

func (p *Pool[T]) Get() T {
	v, _ := p.pool.Get().(T)
	return v
}

type Resetter interface {
	Reset()
}
