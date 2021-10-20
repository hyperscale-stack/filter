// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
	"strings"
)

type stringTrimFilter struct {
}

// NewStringTrimFilter constructor.
func NewStringTrimFilter() Filter {
	return &stringTrimFilter{}
}

func (f stringTrimFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		return strings.TrimSpace(val), nil
	default:
		return value, fmt.Errorf("StringTrimFilter: unsupported type %v", reflect.TypeOf(value))
	}
}
