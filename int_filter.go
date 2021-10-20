// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
	"strconv"
)

type intFilter struct {
}

// NewIntFilter constructor
func NewIntFilter() Filter {
	return &intFilter{}
}

func (f intFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return value, fmt.Errorf("parsing int failed: %w", err)
		}

		return v, nil
	case int, int8, int16, int32, int64:
		return value, nil
	default:
		return value, fmt.Errorf("IntFilter: unsupported type %v", reflect.TypeOf(value))
	}
}
