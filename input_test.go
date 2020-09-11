// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputFilter(t *testing.T) {
	i := NewInputFilter(map[string][]Filter{
		"url": {
			NewURLFilter(),
		},
	})

	value, err := i.Filter(map[string]interface{}{
		"url":  "https://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
		"name": "Title",
	})
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"url":  "https://www.google.fr/",
		"name": "Title",
	}, value)
}

func TestInputFilterWithBadValue(t *testing.T) {
	i := NewInputFilter(map[string][]Filter{
		"url": {
			NewURLFilter(),
		},
	})

	value, err := i.Filter(map[string]interface{}{
		"url":  "134://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
		"name": "Title",
	})
	assert.Error(t, err)

	assert.Equal(t, map[string]interface{}{
		"url":  "134://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
		"name": "Title",
	}, value)
}
