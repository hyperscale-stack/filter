// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/pkg/errors"
)

var badQueryParameters = []string{
	"utm_source",
	"utm_medium",
	"utm_campaign",
	"utm_term",
	"utm_content",
}

type urlFilter struct {
}

// NewURLFilter constructor
func NewURLFilter() Filter {
	return &urlFilter{}
}

func (f urlFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		u, err := url.Parse(val)
		if err != nil {
			return value, err
		}

		q := u.Query()

		for _, key := range badQueryParameters {
			q.Del(key)
		}

		u.RawQuery = q.Encode()

		return u.String(), nil
	default:
		return value, errors.Wrap(fmt.Errorf("unsupported type %v", reflect.TypeOf(value)), "URLFilter")
	}
}
