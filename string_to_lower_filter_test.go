// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToLowerFilter(t *testing.T) {
	f := NewStringToLowerFilter()

	assertions := []struct {
		value    string
		expected string
	}{
		{
			value:    "Test Title",
			expected: "test title",
		},
		{
			value:    "TÉst",
			expected: "tést",
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestStringToLowerFilterWithBadValue(t *testing.T) {
	f := NewStringToLowerFilter()

	u, err := f.Filter(12345)
	assert.Error(t, err)
	assert.Equal(t, 12345, u)
}

func BenchmarkStringToLowerFilter(b *testing.B) {
	f := NewStringToLowerFilter()

	for i := 0; i < b.N; i++ {
		f.Filter("Test Title")
	}
}
