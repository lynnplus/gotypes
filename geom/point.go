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
	"math"
)

type Vector[T Number] struct {
	X, Y T
}

type Point[T Number] struct {
	Vector[T]
}

func (p *Point[T]) Value() Vector[T] {
	return p.Vector
}

func (v Vector[T]) String() string {
	return fmt.Sprintf("(%.12f, %.12f)", v.X, v.Y)
}

func (v Vector[T]) Plus(ov Vector[T]) Vector[T] {
	return Vector[T]{v.X + ov.X, v.Y + ov.Y}
}

// Dot returns the dot product between v and ov.
func (v Vector[T]) Dot(ov Vector[T]) T {
	return v.X*ov.X + v.Y*ov.Y
}

// Sub returns the difference of p and ov.
func (v Vector[T]) Sub(ov Vector[T]) Vector[T] {
	return Vector[T]{v.X - ov.X, v.Y - ov.Y}
}

//func (v Vector[T]) Div(ov Vector[T]) Vector[T] {
//	return
//}

// Mul returns the scalar product of v and m.
func (v Vector[T]) Mul(m float64) Vector[float64] {
	return Vector[float64]{m * float64(v.X), m * float64(v.Y)}
}

// Norm returns the vector's norm.
func (v Vector[T]) Norm() float64 {
	//hypotenuse
	return math.Hypot(float64(v.X), float64(v.Y))
}

// Cross returns the cross product of v and ov.
func (v Vector[T]) Cross(ov Vector[T]) T {
	return v.X*ov.Y - v.Y*ov.X
}

//func (v Vector[T]) Rotate(rad float64) Vector[T]{
//	return Vec[T](v.X*math.Cos(rad)-v.Y*math.Sin(rad),v.X*math.Sin(rad)+v.Y*math.Cos(rad))
//}

// Point return Point struct data
func (v Vector[T]) Point() Point[T] {
	return Point[T]{v}
}

func Vec[T Number](x, y T) Vector[T] {
	return Vector[T]{x, y}
}

func Pt[T Number](x, y T) Point[T] {
	return Point[T]{Vec(x, y)}
}
