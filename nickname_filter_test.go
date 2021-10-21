// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNicknameFilter(t *testing.T) {
	f := NewNicknameFilter()

	assertions := []struct {
		value    string
		expected string
	}{
		{
			value:    "euskadi31",
			expected: "euskadi31",
		},
		{
			value:    "Euskadi31",
			expected: "euskadi31",
		},
		{
			value:    "Euskadi 31",
			expected: "euskadi31",
		},
		{
			value:    "Euskadi.31",
			expected: "euskadi31",
		},
		{
			value:    "Euskadi_31",
			expected: "euskadi31",
		},
		{
			value:    "Éàéù",
			expected: "eaeu",
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestNicknameFilterWithBadValue(t *testing.T) {
	f := NewNicknameFilter()

	u, err := f.Filter(1345)
	assert.Error(t, err)
	assert.Equal(t, 1345, u)
}

func BenchmarkNicknameFilter(b *testing.B) {
	f := NewNicknameFilter()

	for i := 0; i < b.N; i++ {
		f.Filter("Euskadi 31")
	}
}
