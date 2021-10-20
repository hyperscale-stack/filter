// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

// SlugOption type.
type SlugOption func(*slugFilter)

// SlugLanguage config language for SlugFilter.
func SlugLanguage(language string) SlugOption {
	return func(f *slugFilter) {
		f.language = language
	}
}
