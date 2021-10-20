// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
	"strconv"
)

type boolFilter struct {
}

// NewBoolFilter constructor.
func NewBoolFilter() Filter {
	return &boolFilter{}
}

func (f boolFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		v, err := strconv.ParseBool(val)
		if err != nil {
			return value, fmt.Errorf("parsing bool failed: %w", err)
		}

		return v, nil
	case bool:
		return value, nil
	default:
		return value, fmt.Errorf("BoolFilter: unsupported type %v", reflect.TypeOf(value))
	}
}
