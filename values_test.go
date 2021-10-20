// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValuesFilter(t *testing.T) {
	i := NewValuesFilter(map[string][]Filter{
		"url": {
			NewURLFilter(URLStripUTMParameters()),
		},
		"size": {
			NewIntFilter(),
		},
	})

	values := url.Values{}
	values.Set("url", "https://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4")
	values.Set("name", "Title")
	values.Set("size", "1024")

	value, err := i.Filter(values)
	assert.NoError(t, err)

	assert.Equal(t, url.Values{
		"url":  []string{"https://www.google.fr/"},
		"name": []string{"Title"},
		"size": []string{"1024"},
	}, value)
}

func TestValuesFilterWithBadValue(t *testing.T) {
	i := NewValuesFilter(map[string][]Filter{
		"url": {
			NewURLFilter(URLStripUTMParameters()),
		},
	})

	values := url.Values{}
	values.Set("url", "134://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4")
	values.Set("name", "Title")

	value, err := i.Filter(values)
	assert.Error(t, err)

	assert.Equal(t, url.Values{
		"url":  []string{"134://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4"},
		"name": []string{"Title"},
	}, value)
}
