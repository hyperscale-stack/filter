// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

// Value type.
type Value interface{}

// Filter interface.
type Filter interface {
	Filter(value Value) (Value, error)
}
