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

package geom

import (
	"fmt"
	. "github.com/lynnplus/gotypes/constraints"
	"math"
)

type Vector3[T Number] struct {
	X, Y, Z T
}

type Point3[T Number] struct {
	Vector3[T]
}

func (v Vector3[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", v.X, v.Y, v.Z)
}

func (p Point3[T]) Value() Vector3[T] {
	return p.Vector3
}

// Add returns the standard vector sum of v and ov.
func (v Vector3[T]) Add(ov Vector3[T]) Vector3[T] {
	return Vector3[T]{v.X + ov.X, v.Y + ov.Y, v.Z + ov.Z}
}

// Abs returns the vector with non-negative components.
func (v Vector3[T]) Abs() Vector3[T] {
	return Vector3[T]{Abs(v.X), Abs(v.Y), Abs(v.Z)}
}

// Norm returns the vector's norm.
func (v Vector3[T]) Norm() float64 {
	return math.Sqrt(float64(v.Dot(v)))
}

// Norm2 returns the square of the norm.
func (v Vector3[T]) Norm2() T {
	return v.Dot(v)
}

// Mul returns the standard scalar product of v and m.
func (v Vector3[T]) Mul(m float64) Vector3[float64] {
	return Vector3[float64]{m * float64(v.X), m * float64(v.Y), m * float64(v.Z)}
}

// Dot returns the standard dot product of v and ov.
func (v Vector3[T]) Dot(ov Vector3[T]) T {
	return v.X*ov.X + v.Y*ov.Y + v.Z*ov.Z
}

// Sub returns the standard vector difference of v and ov.
func (v Vector3[T]) Sub(ov Vector3[T]) Vector3[T] {
	return Vector3[T]{v.X - ov.X, v.Y - ov.Y, v.Z - ov.Z}
}

// Cross returns the standard cross product of v and ov.
func (v Vector3[T]) Cross(ov Vector3[T]) Vector3[T] {
	return Vector3[T]{
		v.Y*ov.Z - v.Z*ov.Y,
		v.Z*ov.X - v.X*ov.Z,
		v.X*ov.Y - v.Y*ov.X,
	}
}

// Point3 return Point3 struct data
func (v Vector3[T]) Point3() Point3[T] {
	return Point3[T]{v}
}

// Distance returns the Euclidean distance between v and ov.
func (v Vector3[T]) Distance(ov Vector3[T]) float64 {
	return v.Sub(ov).Norm()
}

// Normalize returns a unit vector in the same direction as v.
func (v Vector3[T]) Normalize() Vector3[float64] {
	n2 := v.Norm2()
	if n2 == 0 {
		return Vector3[float64]{0, 0, 0}
	}
	return v.Mul(1 / math.Sqrt(float64(n2)))
}

func Vec3[T Number](x, y, z T) Vector3[T] {
	return Vector3[T]{x, y, z}
}

func Pt3[T Number](x, y, z T) Point3[T] {
	return Point3[T]{Vec3(x, y, z)}
}
