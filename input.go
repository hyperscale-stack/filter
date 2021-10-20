// Copyright 2019 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package filter

// InputValue type.
type InputValue map[string]interface{}

// InputFilter interface.
type InputFilter interface {
	Filter(input map[string]interface{}) (map[string]interface{}, error)
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
			return val, err
		}
	}

	return val, nil
}

func (f inputFilter) Filter(input map[string]interface{}) (map[string]interface{}, error) {
	for key, val := range input {
		val, err := f.filterField(key, val)
		if err != nil {
			return input, err
		}

		input[key] = val
	}

	return input, nil
}
