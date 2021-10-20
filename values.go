// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"net/url"
)

// ValuesFilter interface.
type ValuesFilter interface {
	Filter(input url.Values) (url.Values, error)
}

type valuesFilter struct {
	filters map[string][]Filter
}

// NewValuesFilter constructor.
func NewValuesFilter(filters map[string][]Filter) ValuesFilter {
	return &valuesFilter{
		filters: filters,
	}
}

func (f valuesFilter) filterField(key string, values []Value) ([]Value, error) {
	filters, ok := f.filters[key]
	if !ok {
		return values, nil
	}

	retvals := make([]Value, len(values))

	for i, value := range values {
		val := value

		var err error

		for _, filter := range filters {
			val, err = filter.Filter(val)
			if err != nil {
				return retvals, err
			}

			retvals[i] = val
		}
	}

	return retvals, nil
}

func (f valuesFilter) castSliceStringToSliceValue(values []string) []Value {
	retvals := make([]Value, len(values))

	for i, val := range values {
		retvals[i] = val
	}

	return retvals
}

func (f valuesFilter) castSliceValueToSliceString(values []Value) []string {
	retvals := make([]string, len(values))

	for i, val := range values {
		switch v := val.(type) {
		case string:
			retvals[i] = v
		default:
			retvals[i] = fmt.Sprintf("%v", val)
		}
	}

	return retvals
}

func (f valuesFilter) Filter(values url.Values) (url.Values, error) {
	for key, vals := range values {
		vals, err := f.filterField(key, f.castSliceStringToSliceValue(vals))
		if err != nil {
			return values, err
		}

		values[key] = f.castSliceValueToSliceString(vals)
	}

	return values, nil
}
