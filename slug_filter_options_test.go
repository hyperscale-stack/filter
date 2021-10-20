// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugLanguage(t *testing.T) {
	f := &slugFilter{
		language: "en",
	}

	SlugLanguage("fr")(f)

	assert.Equal(t, "fr", f.language)
}
