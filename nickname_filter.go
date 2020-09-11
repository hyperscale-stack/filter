// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gosimple/slug"
	"github.com/pkg/errors"
)

type nicknameFilter struct {
}

// NewNicknameFilter constructor
func NewNicknameFilter() Filter {
	return &nicknameFilter{}
}

func (f nicknameFilter) Filter(value Value) (Value, error) {
	switch val := value.(type) {
	case string:
		nickname := slug.Make(val)
		nickname = strings.Replace(nickname, "-", "", -1)
		nickname = strings.Replace(nickname, ".", "", -1)
		nickname = strings.Replace(nickname, "_", "", -1)

		return nickname, nil
	default:
		return value, errors.Wrap(fmt.Errorf("unsupported type %v", reflect.TypeOf(value)), "NicknameFilter")
	}
}
