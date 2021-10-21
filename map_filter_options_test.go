// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFilterForKey(t *testing.T) {
	f := &mapFilter{
		def: map[interface{}][]Filter{},
	}

	assert.Equal(t, 0, len(f.def))

	MapFilterForKey("foo", NewStringTrimFilter(), NewStringToUpperFilter())(f)

	assert.Contains(t, f.def, "foo")

	assert.Equal(t, 2, len(f.def["foo"]))
}
