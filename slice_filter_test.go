// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceFilter(t *testing.T) {
	f := NewSliceFilter(
		NewStringToLowerFilter(),
	)

	assertions := []struct {
		value    []string
		expected []Value
	}{
		{
			value:    []string{"FR", "EN"},
			expected: []Value{"fr", "en"},
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestSliceFilterWithBadValue(t *testing.T) {
	f := NewSliceFilter(
		NewStringToLowerFilter(),
	)

	u, err := f.Filter(1345)
	assert.Error(t, err)
	assert.Equal(t, 1345, u)

	u, err = f.Filter(nil)
	assert.Error(t, err)
	assert.Nil(t, u)

	u, err = f.Filter([]int{1345})
	assert.Error(t, err)
	assert.Equal(t, []int{1345}, u)
}

func BenchmarkSliceFilter(b *testing.B) {
	f := NewSliceFilter(
		NewStringToLowerFilter(),
	)

	for i := 0; i < b.N; i++ {
		f.Filter([]string{"FR", "EN"})
	}
}
