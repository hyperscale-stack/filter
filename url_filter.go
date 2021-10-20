// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"net/url"
	"reflect"
)

type urlFilter struct {
	stripQueryParameters []string
}

// NewURLFilter constructor
func NewURLFilter(opts ...URLOption) Filter {
	f := &urlFilter{
		stripQueryParameters: []string{},
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f urlFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		u, err := url.Parse(val)
		if err != nil {
			return value, fmt.Errorf("parsing url failed: %w", err)
		}

		q := u.Query()

		for _, key := range f.stripQueryParameters {
			q.Del(key)
		}

		u.RawQuery = q.Encode()

		return u.String(), nil
	default:
		return value, fmt.Errorf("URLFilter: unsupported type %v", reflect.TypeOf(value))
	}
}
