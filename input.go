// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

import (
	"fmt"
	"net/url"
)

// InputValue type.
type InputValue map[string]interface{}

// InputFilter interface.
type InputFilter interface {
	FilterMap(input map[string]interface{}) (map[string]interface{}, error)
	FilterValues(input url.Values) (url.Values, error)
}

type inputFilter struct {
	filters map[string][]Filter
}

// NewInputFilter constructor.
func NewInputFilter(filters map[string][]Filter) InputFilter {
	return &inputFilter{
		filters: filters,
	}
}

func (f inputFilter) filterField(key string, value Value) (Value, error) {
	filters, ok := f.filters[key]
	if !ok {
		return value, nil
	}

	val := value

	var err error

	for _, filter := range filters {
		val, err = filter.Filter(val)
		if err != nil {
			return val, fmt.Errorf("apply filter: %w", err)
		}
	}

	return val, nil
}

func (f inputFilter) FilterMap(input map[string]interface{}) (map[string]interface{}, error) {
	for key, val := range input {
		val, err := f.filterField(key, val)
		if err != nil {
			return input, fmt.Errorf("apply filter: %w", err)
		}

		input[key] = val
	}

	return input, nil
}

func (f inputFilter) filterFieldValues(key string, values []Value) ([]Value, error) {
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
				return retvals, fmt.Errorf("apply filter: %w", err)
			}

			retvals[i] = val
		}
	}

	return retvals, nil
}

func (f inputFilter) castSliceStringToSliceValue(values []string) []Value {
	retvals := make([]Value, len(values))

	for i, val := range values {
		retvals[i] = val
	}

	return retvals
}

func (f inputFilter) castSliceValueToSliceString(values []Value) []string {
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

func (f inputFilter) FilterValues(values url.Values) (url.Values, error) {
	for key, vals := range values {
		vals, err := f.filterFieldValues(key, f.castSliceStringToSliceValue(vals))
		if err != nil {
			return values, fmt.Errorf("apply filter: %w", err)
		}

		values[key] = f.castSliceValueToSliceString(vals)
	}

	return values, nil
}
