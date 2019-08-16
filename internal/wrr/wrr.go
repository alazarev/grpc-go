/*
 *
 * Copyright 2019 gRPC authors.
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

package wrr

// WRR defines an interface that implements weighted round robin.
type WRR interface {
	// Add adds an item with weight to the WRR set. If the item is already in the set weight will be added.
	Add(item interface{}, weight int64)
	// Next returns the next picked item.
	//
	// Next needs to be thread safe.
	Next() interface{}

	// UpdateOrAdd changes weight of item to provided value. Adds an item if it doesn't exist.
	UpdateOrAdd(item interface{}, weight int64)

	// Remove removes item from the WRR set.
	Remove(item interface{})
}