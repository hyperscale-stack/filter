// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

type stringToLowerFilter struct {
}

// NewStringToLowerFilter constructor
func NewStringToLowerFilter() Filter {
	return &stringToLowerFilter{}
}

func (f stringToLowerFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		return strings.ToLower(val), nil
	default:
		return value, errors.Wrap(fmt.Errorf("unsupported type %v", reflect.TypeOf(value)), "StringToLowerFilter")
	}
}
