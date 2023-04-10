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

package geom

import (
	"fmt"
	. "github.com/lynnplus/gotypes/constraints"
)

type Point[T Number] struct {
	X, Y T
}

func (p Point[T]) String() string {
	return fmt.Sprintf("(%v,%v)", p.X, p.Y)
}

func (p Point[T]) Add(q Point[T]) Point[T] {
	return Point[T]{p.X + q.X, p.Y + q.Y}
}

func (p Point[T]) ToSize() Size[T] {
	return Size[T]{p.X, p.Y}
}

func Pt[T Number](x, y T) Point[T] {
	return Point[T]{x, y}
}
