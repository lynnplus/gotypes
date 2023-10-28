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

var (
	_ Array[int] = (*LinkedList[int])(nil)
)

type listElement[T comparable] struct {
	value T
	prev  *listElement[T]
	next  *listElement[T]
}

type LinkedList[T comparable] struct {
	first *listElement[T]
	last  *listElement[T]
	size  int
}

// NewLinkedList return an LinkedList
func NewLinkedList[T comparable](values ...T) *LinkedList[T] {
	list := &LinkedList[T]{}
	list.Add(values...)
	return list
}

func (list *LinkedList[T]) Add(values ...T) {
	for _, value := range values {
		newElement := &listElement[T]{value: value, prev: list.last}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

func (list *LinkedList[T]) Set(index int, v T) {
	if !list.checkInRange(index) {
		if index == list.size {
			list.Add(v)
		}
		return
	}
	var temp *listElement[T]
	if index > list.size-index {
		temp = list.last
		for i := list.size - 1; i != index; i, temp = i-1, temp.prev {
		}
	} else {
		temp = list.first
		for i := 0; i != index; i, temp = i+1, temp.next {
		}
	}
	temp.value = v
}

func (list *LinkedList[T]) Get(index int) (T, bool) {
	if !list.checkInRange(index) {
		return *new(T), false
	}

	if list.size-index < index {
		element := list.last
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
		}
		return element.value, true
	}
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

func (list *LinkedList[T]) Remove(index int) {
	if !list.checkInRange(index) {
		return
	}

	if list.size == 1 {
		list.RemoveAll()
		return
	}

	var temp *listElement[T]
	if index > list.size-index {
		temp = list.last
		for i := list.size - 1; i != index; i, temp = i-1, temp.prev {
		}
	} else {
		temp = list.first
		for i := 0; i != index; i, temp = i+1, temp.next {
		}
	}
	if temp == list.first {
		list.first = temp.next
	}
	if temp == list.last {
		list.last = temp.prev
	}
	if temp.prev != nil {
		temp.prev.next = temp.next
	}
	if temp.next != nil {
		temp.next.prev = temp.prev
	}
	temp = nil
	list.size--
}

func (list *LinkedList[T]) RemoveAll() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *LinkedList[T]) IndexOf(value T) int {
	if list.size == 0 {
		return -1
	}
	index := -1
	list.Range(func(i int, v T) bool {
		if v == value {
			index = i
			return false
		}
		return true
	})
	return index
}

func (list *LinkedList[T]) Range(f func(index int, value T) bool) {
	for i, ele := 0, list.first; ele != nil; i, ele = i+1, ele.next {
		if !f(i, ele.value) {
			break
		}
	}
}

func (list *LinkedList[T]) ReverseRange(f func(index int, value T) bool) {
	for i, ele := list.size-1, list.last; ele != nil; i, ele = i-1, ele.prev {
		if !f(i, ele.value) {
			break
		}
	}
}

func (list *LinkedList[T]) Each(f func(index int, value T)) {
	list.Range(func(i int, v1 T) bool {
		f(i, v1)
		return true
	})
}

func (list *LinkedList[T]) Every(f func(index int, value T) bool) bool {
	ok := true
	list.Range(func(i int, v1 T) bool {
		ok = f(i, v1)
		return ok
	})
	return ok
}

func (list *LinkedList[T]) Some(f func(index int, value T) bool) bool {
	ok := true
	list.Range(func(i int, v1 T) bool {
		ok = f(i, v1)
		return !ok
	})
	return ok
}

func (list *LinkedList[T]) Values() []T {
	values := make([]T, list.size, list.size)
	list.Each(func(index int, value T) {
		values[index] = value
	})
	return values
}

func (list *LinkedList[T]) Empty() bool {
	return list.size == 0
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) checkInRange(index int) bool {
	return index >= 0 && index < list.size
}
