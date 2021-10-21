// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
)

type mapFilter struct {
	def map[interface{}][]Filter
}

// NewMapFilter constructor.
func NewMapFilter(opts ...MapOption) Filter {
	f := &mapFilter{
		def: map[interface{}][]Filter{},
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f mapFilter) Filter(value Value) (Value, error) {
	s := reflect.ValueOf(value)
	if s.Kind() != reflect.Map {
		return value, fmt.Errorf("value is not a map type: %v", s)
	}
	/*
		if s.IsNil() {
			return value, nil
		}
	*/

	data := make(map[string]Value, s.Len())

	for _, key := range s.MapKeys() {
		v := s.MapIndex(key)

		data[key.String()] = v.Interface()

		filters, ok := f.def[key.Interface()]
		if ok {
			for _, filter := range filters {
				val, err := filter.Filter(data[key.String()])
				if err != nil {
					return value, err
				}

				data[key.String()] = val
			}
		}
	}

	return data, nil
}
