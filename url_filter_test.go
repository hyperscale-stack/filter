// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLFilter(t *testing.T) {
	f := NewURLFilter(URLStripUTMParameters())

	assertions := []struct {
		value    string
		expected string
	}{
		{
			value:    "https://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
			expected: "https://www.google.fr/",
		},
		{
			value:    "https://www.google.fr/?foo=bar&utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
			expected: "https://www.google.fr/?foo=bar",
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestURLFilterWithBadValue(t *testing.T) {
	f := NewURLFilter()

	u, err := f.Filter(1235)
	assert.Equal(t, 1235, u)
	assert.Error(t, err)

	u, err = f.Filter("134://foo")
	assert.Equal(t, "134://foo", u)
	assert.Error(t, err)
}
