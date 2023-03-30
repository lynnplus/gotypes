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

import "github.com/lynnplus/gotypes/constraints"

type Map[K constraints.Basic, V any] interface {
	Get(key K) V
	Exist(key K) (ok bool)

	Store(key K, value V)
	Load(key K) (value V, ok bool)

	Range(f func(key K, value V) bool)
	Each(f func(key K, value V))

	Keys() []K
	Values() []V

	Size() int

	Delete(key K)
	DeleteAll()

	Data() map[K]V
}

type SafeMap[K constraints.Basic, V any] interface {
	Map[K, V]

	LoadOrStore(key K, value V) (actual V, loaded bool)
	LoadAndDelete(key K) (value V, loaded bool)
}

var _ Enumerable[int] = (*GoMap[string, int])(nil)

type GoMap[K comparable, V any] map[K]V

func (g GoMap[K, V]) Each(f func(key K, value V)) {
	for k, v := range g {
		f(k, v)
	}
}

func (g GoMap[K, V]) Size() int {
	return len(g)
}

func (g GoMap[K, V]) EachValue(f func(value V)) {
	g.Each(func(k K, v V) {
		f(v)
	})
}
