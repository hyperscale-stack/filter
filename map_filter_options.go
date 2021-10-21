// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

type MapOption func(*mapFilter)

func MapFilterForKey(key interface{}, filters ...Filter) MapOption {
	return func(mf *mapFilter) {
		mf.def[key] = filters
	}
}
