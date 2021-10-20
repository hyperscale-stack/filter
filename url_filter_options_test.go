// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLStripUTMParameters(t *testing.T) {
	uf := &urlFilter{
		stripQueryParameters: []string{},
	}

	assert.Equal(t, []string{}, uf.stripQueryParameters)

	URLStripUTMParameters()(uf)

	assert.Equal(t, utmQueryParameters, uf.stripQueryParameters)
}
