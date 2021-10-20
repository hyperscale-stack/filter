// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringTrimFilter(t *testing.T) {
	f := NewStringTrimFilter()

	assertions := []struct {
		value    string
		expected string
	}{
		{
			value:    "     Test Title       ",
			expected: "Test Title",
		},
		{
			value:    "  TÉst    ",
			expected: "TÉst",
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestStringTrimFilterWithBadValue(t *testing.T) {
	f := NewStringTrimFilter()

	u, err := f.Filter(12345)
	assert.Error(t, err)
	assert.Equal(t, 12345, u)
}
