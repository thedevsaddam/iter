iter
---
[![Build Status](https://travis-ci.org/thedevsaddam/iter.svg?branch=master)](https://travis-ci.org/thedevsaddam/iter)
[![Project status](https://img.shields.io/badge/version-v1-green.svg)](https://github.com/thedevsaddam/iter/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/thedevsaddam/iter)](https://goreportcard.com/report/github.com/thedevsaddam/iter)
[![Coverage Status](https://coveralls.io/repos/github/thedevsaddam/iter/badge.svg?branch=master)](https://coveralls.io/github/thedevsaddam/iter)
[![GoDoc](https://godoc.org/github.com/thedevsaddam/iter?status.svg)](https://pkg.go.dev/github.com/thedevsaddam/iter)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](LICENSE.md)

Iter provides functionality like Python's range function to iterate over numbers and letters

### Installation

Install the package using
```go
$ go get github.com/thedevsaddam/iter
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/thedevsaddam/iter"
```

Let's see a quick example:

```go
package main

import (
	"fmt"

	"github.com/thedevsaddam/iter"
)

func main() {
	// sequence: 0-9
	for v := range iter.N(10) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// output: 0 1 2 3 4 5 6 7 8 9

	// sequence: 5-9
	for v := range iter.N(5, 10) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// output: 5 6 7 8 9

	// sequence: 1-9, increment by 2
	for v := range iter.N(5, 10, 2) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// output: 5 7 9

	// sequence: a-e
	for v := range iter.L('a', 'e') {
		fmt.Printf("%s ", string(v))
	}
	fmt.Println()
	// output: a b c d e
}

```


## Bugs and Issues

If you encounter any bugs or issues, feel free to [open an issue at
github](https://github.com/thedevsaddam/iter/issues).

Also, you can shoot me an email to
<mailto:thedevsaddam@gmail.com> for hugs or bugs.

## Contribution
If you are interested to make the package better please send pull requests or create an issue so that others can fix.
[Read the contribution guide here](CONTRIBUTING.md)


## License
The **iter** is an open-source software licensed under the [MIT License](LICENSE.md).