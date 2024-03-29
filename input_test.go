// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputFilter(t *testing.T) {
	i := NewInputFilter(map[string][]Filter{
		"url": {
			NewURLFilter(URLStripUTMParameters()),
		},
		"size": {
			NewIntFilter(),
		},
	})

	{
		value, err := i.FilterMap(map[string]interface{}{
			"url":  "https://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
			"name": "Title",
			"size": "1024",
		})
		assert.NoError(t, err)

		assert.Equal(t, map[string]interface{}{
			"url":  "https://www.google.fr/",
			"name": "Title",
			"size": int64(1024),
		}, value)
	}

	{
		values := url.Values{}
		values.Set("url", "https://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4")
		values.Set("name", "Title")
		values.Set("size", "1024")

		value, err := i.FilterValues(values)
		assert.NoError(t, err)

		assert.Equal(t, url.Values{
			"url":  []string{"https://www.google.fr/"},
			"name": []string{"Title"},
			"size": []string{"1024"},
		}, value)
	}
}

func TestInputFilterWithBadValue(t *testing.T) {
	i := NewInputFilter(map[string][]Filter{
		"url": {
			NewURLFilter(URLStripUTMParameters()),
		},
	})

	{
		value, err := i.FilterMap(map[string]interface{}{
			"url":  "134://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
			"name": "Title",
		})
		assert.Error(t, err)

		assert.Equal(t, map[string]interface{}{
			"url":  "134://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
			"name": "Title",
		}, value)
	}

	{
		values := url.Values{}
		values.Set("url", "134://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4")
		values.Set("name", "Title")

		value, err := i.FilterValues(values)
		assert.Error(t, err)

		assert.Equal(t, url.Values{
			"url":  []string{"134://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4"},
			"name": []string{"Title"},
		}, value)
	}
}

func BenchmarkInputFilterMap(b *testing.B) {
	f := NewInputFilter(map[string][]Filter{
		"url": {
			NewURLFilter(URLStripUTMParameters()),
		},
		"size": {
			NewIntFilter(),
		},
	})

	input := map[string]interface{}{
		"url":  "https://www.google.fr/?utm_source=test&utm_medium=test1&utm_campaign=test2&utm_term=test3&utm_content=test4",
		"name": "Title",
		"size": "1024",
	}

	for i := 0; i < b.N; i++ {
		f.FilterMap(input)
	}
}

func BenchmarkInputFilterValues(b *testing.B) {
	f := NewInputFilter(map[string][]Filter{
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

	for i := 0; i < b.N; i++ {
		f.FilterValues(values)
	}
}
