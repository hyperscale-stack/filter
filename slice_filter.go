// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
)

type sliceFilter struct {
	filters []Filter
}

// NewSliceFilter constructor.
func NewSliceFilter(filters ...Filter) Filter {
	return &sliceFilter{
		filters: filters,
	}
}

func (f sliceFilter) Filter(value Value) (Value, error) {
	s := reflect.ValueOf(value)
	if s.Kind() != reflect.Slice {
		return value, fmt.Errorf("value is not a slice type: %v", s)
	}
	/*
		if s.IsNil() {
			return value, nil
		}
	*/
	items := make([]Value, s.Len())

	for i := 0; i < s.Len(); i++ {
		items[i] = s.Index(i).Interface()
	}

	for i, val := range items {
		for _, ftr := range f.filters {
			v, err := ftr.Filter(val)
			if err != nil {
				return value, err
			}

			val = v
		}

		items[i] = val
	}

	return items, nil
}
