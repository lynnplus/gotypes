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

import (
	"fmt"
	"strconv"
)

func ExampleConvertTo() {
	data := map[string]int{}

	data["a"] = 1
	data["b"] = 2
	data["c"] = 3

	r := ConvertTo[int](GoMap[string, int](data), func(v int) string {
		return strconv.Itoa(v)
	})
	fmt.Println(r)
	// Output: [1 2 3]
}
