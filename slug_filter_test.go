// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugFilter(t *testing.T) {
	f := NewSlugFilter()

	assertions := []struct {
		value    string
		expected string
	}{
		{
			value:    "My super cool title",
			expected: "my-super-cool-title",
		},
		{
			value:    "Insolite : Apple lance une chiffonnette à... 25 euros",
			expected: "insolite-apple-lance-une-chiffonnette-a-25-euros",
		},
		{
			value:    "Sophie & Axel",
			expected: "sophie-and-axel",
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestSlugFilterWithLanguage(t *testing.T) {
	f := NewSlugFilter(SlugLanguage("fr"))

	assertions := []struct {
		value    string
		expected string
	}{
		{
			value:    "My super cool title",
			expected: "my-super-cool-title",
		},
		{
			value:    "Insolite : Apple lance une chiffonnette à... 25 euros",
			expected: "insolite-apple-lance-une-chiffonnette-a-25-euros",
		},
		{
			value:    "Sophie & Axel",
			expected: "sophie-et-axel",
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestSlugFilterWithBadValue(t *testing.T) {
	f := NewSlugFilter()

	u, err := f.Filter(1345)
	assert.Error(t, err)
	assert.Equal(t, 1345, u)
}

func BenchmarkSlugFilter(b *testing.B) {
	f := NewSlugFilter()

	for i := 0; i < b.N; i++ {
		f.Filter("Test Title")
	}
}
