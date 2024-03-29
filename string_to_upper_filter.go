// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
	"strings"
)

type stringToUpperFilter struct {
}

// NewStringToUpperFilter constructor.
func NewStringToUpperFilter() Filter {
	return &stringToUpperFilter{}
}

func (f stringToUpperFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		return strings.ToUpper(val), nil
	default:
		return value, fmt.Errorf("StringToUpperFilter: unsupported type %v", reflect.TypeOf(value))
	}
}
