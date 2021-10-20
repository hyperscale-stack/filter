Hyperscale Filter [![Last release](https://img.shields.io/github/release/hyperscale-stack/filter.svg)](https://github.com/hyperscale-stack/filter/releases/latest) [![Documentation](https://godoc.org/github.com/hyperscale-stack/filter?status.svg)](https://godoc.org/github.com/hyperscale-stack/filter)
====================

[![Go Report Card](https://goreportcard.com/badge/github.com/hyperscale-stack/filter)](https://goreportcard.com/report/github.com/hyperscale-stack/filter)

| Branch  | Status | Coverage |
|---------|--------|----------|
| master  | [![Build Status](https://github.com/hyperscale-stack/filter/workflows/Go/badge.svg?branch=master)](https://github.com/hyperscale-stack/filter/actions?query=workflow%3AGo) | [![Coveralls](https://img.shields.io/coveralls/hyperscale-stack/filter/master.svg)](https://coveralls.io/github/hyperscale-stack/filter?branch=master) |

The Hyperscale Filter library provides a set of commonly needed data filters. It also provides a simple filter chaining mechanism by which multiple filters may be applied to a single datum in a user-defined order. 

## Example

Filter by `map[string]interface{}`

```go
package main

import (
    "fmt"
    "github.com/hyperscale-stack/filter"
)

func main() {

    i := NewInputFilter(map[string][]Filter{
		"email": {
			NewStringToLowerFilter(),
		},
	})

	value, err := i.Filter(map[string]interface{}{
		"email":  "STEVE@APPLE.COM",
    })
    // return 
    // map[string]interface{}{
	//     "email":  "steve@apple.com",
    // }
}

```


Filter by `url.Values`

```go
package main

import (
    "fmt"
    "github.com/hyperscale-stack/filter"
)

func main() {

    i := NewValuesFilter(map[string][]Filter{
		"email": {
			NewStringToLowerFilter(),
		},
	})

    values := url.Values{}
    values.Set("email", "STEVE@APPLE.COM")

	value, err := i.Filter(values)
    // return 
    // url.Values{
	//     "email":  []string{"steve@apple.com"},
    // }
}

```


## License

Hyperscale Filter is licensed under [the MIT license](LICENSE.md).
