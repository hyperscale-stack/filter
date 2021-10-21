// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatFilterFromStringValue(t *testing.T) {
	f := NewFloatFilter()

	assertions := []struct {
		value    interface{}
		expected interface{}
	}{
		{
			value:    "123456.955",
			expected: float64(123456.955),
		},
		{
			value:    "123456",
			expected: float64(123456.0),
		},
		{
			value:    555.5,
			expected: 555.5,
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestFloatFilterWithBadValue(t *testing.T) {
	f := NewFloatFilter()

	u, err := f.Filter("ffff")
	assert.Error(t, err)
	assert.Equal(t, "ffff", u)

	u, err = f.Filter(true)
	assert.Error(t, err)
	assert.Equal(t, true, u)
}

func BenchmarkFloatFilter(b *testing.B) {
	f := NewFloatFilter()

	for i := 0; i < b.N; i++ {
		f.Filter("123456789.455")
	}
}
