// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gosimple/slug"
)

var nicknameDefaultStripChars = []string{
	"-",
	".",
	"_",
}

type nicknameFilter struct {
	stripChars []string
}

// NewNicknameFilter constructor
func NewNicknameFilter() Filter {
	return &nicknameFilter{
		stripChars: nicknameDefaultStripChars,
	}
}

func (f nicknameFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		nickname := slug.Make(val)

		for _, c := range f.stripChars {
			nickname = strings.Replace(nickname, c, "", -1)
		}

		return nickname, nil
	default:
		return value, fmt.Errorf("NicknameFilter: unsupported type %v", reflect.TypeOf(value))
	}
}
