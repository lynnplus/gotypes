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

import . "github.com/lynnplus/gotypes/constraints"

type Size[T Number] struct {
	Width, Height T
}

type Rectangle[T Number] struct {
	Min, Max Point[T]
}

func (rt Rectangle[T]) String() string {
	return rt.Min.String() + "-" + rt.Max.String()
}

func (rt Rectangle[T]) Size() Size[T] {
	return Size[T]{
		rt.Max.X - rt.Min.X,
		rt.Max.Y - rt.Min.Y,
	}
}

func Rect[T Number](x0, y0, x1, y1 T) Rectangle[T] {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rectangle[T]{Point[T]{x0, y0}, Point[T]{x1, y1}}
}

type RotatedRect[T Number] struct {
	Center Point[T]
	Size   Size[T]
	Angle  float64
}

type Range[T Number] struct {
	Start, End T
}
