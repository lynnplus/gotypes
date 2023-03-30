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

type SafeArray[V any] interface {
	Add(src ...V)
	Remove(index int)
	Set(index int, v V)
	Get(index int) V
	Length() int
	Capacity() int
	Range(f func(index int, value V) bool)
	Each(f func(index int, value V))
	Data() []V
}

type Container interface {
	Empty() bool
	Size() int
	RemoveAll()
}

type Array[V any] interface {
	Container
	Enumerable2[int, V]

	Add(src ...V)
	Remove(index int)
	Get(index int) (V, bool)
	Set(index int, v V)
	Values() []V
	IndexOf(value V) int
}

type Enumerable2[K any, V any] interface {
	// Each calls the given function once for each element, passing that element's index(or key) and value.
	Each(f func(index K, value V))
	// Range passes each element of the list to the given function and
	// break loop if the function returns false.
	Range(f func(index K, value V) bool)
	// Every passes each element of the list to the given function and
	// returns true if the function returns true for all elements.
	Every(f func(index K, value V) bool) bool
	// Some passes each element of the container to the given function and
	// returns true if the function ever returns true for any element.
	Some(f func(index K, value V) bool) bool
}
