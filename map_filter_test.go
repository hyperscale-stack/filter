// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFilter(t *testing.T) {
	f := NewMapFilter(
		MapFilterForKey("language", NewStringToLowerFilter()),
		MapFilterForKey("country", NewStringToUpperFilter()),
	)

	assertions := []struct {
		value    map[string]Value
		expected map[string]Value
	}{
		{
			value: map[string]Value{
				"country":  "fr",
				"language": "FR",
			},
			expected: map[string]Value{
				"country":  "FR",
				"language": "fr",
			},
		},
	}

	for _, assertion := range assertions {
		u, err := f.Filter(assertion.value)
		assert.NoError(t, err)

		assert.Equal(t, assertion.expected, u)
	}
}

func TestMapFilterWithBadValue(t *testing.T) {
	f := NewMapFilter(
		MapFilterForKey("language", NewStringToLowerFilter()),
		MapFilterForKey("country", NewStringToUpperFilter()),
	)

	u, err := f.Filter(1345)
	assert.Error(t, err)
	assert.Equal(t, 1345, u)

	u, err = f.Filter(nil)
	assert.Error(t, err)
	assert.Nil(t, u)

	u, err = f.Filter(map[string]int{"language": 1223})
	assert.Error(t, err)
	assert.Equal(t, map[string]int{"language": 1223}, u)
}

func BenchmarkMapFilter(b *testing.B) {
	f := NewMapFilter(
		MapFilterForKey("language", NewStringToLowerFilter()),
		MapFilterForKey("country", NewStringToUpperFilter()),
	)

	for i := 0; i < b.N; i++ {
		f.Filter(map[string]string{
			"country":  "fr",
			"language": "FR",
		})
	}
}
