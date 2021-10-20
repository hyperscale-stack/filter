// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
	"strconv"
)

type floatFilter struct {
}

// NewFloatFilter constructor.
func NewFloatFilter() Filter {
	return &floatFilter{}
}

func (f floatFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return value, fmt.Errorf("parsing float failed: %w", err)
		}

		return v, nil
	case float32, float64:
		return value, nil
	default:
		return value, fmt.Errorf("FloatFilter: unsupported type %v", reflect.TypeOf(value))
	}
}
