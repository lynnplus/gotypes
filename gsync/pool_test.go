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
	"testing"
)

type Item int

func (t Item) Reset() {
}

func BenchmarkPool(b *testing.B) {
	var p sync.Pool
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkGsyncPool(b *testing.B) {
	p := NewPool[Item](nil)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkPoolOverflow(b *testing.B) {
	var p sync.Pool
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for b := 0; b < 100; b++ {
				p.Put(1)
			}
			for b := 0; b < 100; b++ {
				p.Get()
			}
		}
	})
}

func BenchmarkGsyncPoolOverflow(b *testing.B) {
	p := NewPool[Item](nil)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for b := 0; b < 100; b++ {
				p.Put(1)
			}
			for b := 0; b < 100; b++ {
				p.Get()
			}
		}
	})
}
