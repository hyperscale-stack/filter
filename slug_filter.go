// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"

	"github.com/gosimple/slug"
)

type slugFilter struct {
	language string
}

// NewSlugFilter constructor
func NewSlugFilter(opts ...SlugOption) Filter {
	f := &slugFilter{
		language: "en",
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f slugFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		return slug.MakeLang(val, f.language), nil
	default:
		return value, fmt.Errorf("SlugFilter: unsupported type %v", reflect.TypeOf(value))
	}
}
