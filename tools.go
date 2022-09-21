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

type Sizer interface {
	Size() int
}

type Enumerable[V any] interface {
	Sizer
	EachValue(f func(value V))
}

type EnumerableWithKey[K comparable, V any] interface {
	Sizer
	Each(func(key K, value V))
}

func ConvertTo[V any, R any](source Enumerable[V], convert func(V) R) []R {
	r := make([]R, 0, 10)
	source.EachValue(func(value V) {
		r = append(r, convert(value))
	})
	return r
}

func ConvertToWithKey[K comparable, V any, R any](source EnumerableWithKey[K, V], convert func(K, V) R) []R {
	r := make([]R, 0, 10)
	source.Each(func(key K, value V) {
		r = append(r, convert(key, value))
	})
	return r
}

func ConvertToWithIndex[V any, R any](source Enumerable[V], convert func(int, V) R) []R {
	r := make([]R, 0, 10)
	index := 0
	source.EachValue(func(value V) {
		r = append(r, convert(index, value))
		index++
	})
	return r
}
