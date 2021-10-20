// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntFilterFromStringValue(t *testing.T) {
	f := NewIntFilter()

	assertions := []struct {
		value    interface{}
		expected interface{}
	}{
		{
			value:    "123456",
			expected: int64(123456),
		},
		{
			value:    555,
			expected: 555,
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestIntFilterWithBadValue(t *testing.T) {
	f := NewIntFilter()

	u, err := f.Filter("ffff")
	assert.Error(t, err)
	assert.Equal(t, "ffff", u)

	u, err = f.Filter(true)
	assert.Error(t, err)
	assert.Equal(t, true, u)
}
