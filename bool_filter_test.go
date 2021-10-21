// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolFilterFromStringValue(t *testing.T) {
	f := NewBoolFilter()

	assertions := []struct {
		value    interface{}
		expected interface{}
	}{
		{
			value:    "1",
			expected: true,
		},
		{
			value:    "0",
			expected: false,
		},
		{
			value:    "true",
			expected: true,
		},
		{
			value:    "false",
			expected: false,
		},
		{
			value:    true,
			expected: true,
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestBoolFilterWithBadValue(t *testing.T) {
	f := NewBoolFilter()

	u, err := f.Filter("ffff")
	assert.Error(t, err)
	assert.Equal(t, "ffff", u)

	u, err = f.Filter(5.5)
	assert.Error(t, err)
	assert.Equal(t, 5.5, u)
}

func BenchmarkBoolFilter(b *testing.B) {
	f := NewBoolFilter()

	for i := 0; i < b.N; i++ {
		f.Filter("true")
	}
}
