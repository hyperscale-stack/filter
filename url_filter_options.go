// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

// URLOption type.
type URLOption func(*urlFilter)

var utmQueryParameters = []string{
	"utm_source",
	"utm_medium",
	"utm_campaign",
	"utm_term",
	"utm_content",
}

// URLStripUTMParameters remove all utm_* query parameters.
func URLStripUTMParameters() URLOption {
	return func(uf *urlFilter) {
		uf.stripQueryParameters = append(uf.stripQueryParameters, utmQueryParameters...)
	}
}
