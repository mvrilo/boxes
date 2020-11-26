# boxes

[![GoDoc](https://godoc.org/github.com/mvrilo/boxes?status.svg)](https://godoc.org/github.com/mvrilo/boxes)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvrilo/boxes)](https://goreportcard.com/report/github.com/mvrilo/boxes)

Boxes is a simple ascii box builder.

## Example

```go
// Output:
// ___________________
// |                 |
// |   hello world   |
// |_________________|

package main

import "github.com/mvrilo/boxes"

func main() {
	box, err := boxes.New().Padding(2).Write([]byte("hello world"))
	if err != nil {
		panic(err)
	}
	println(string(box.Render()))
}
```

## License

MIT

## Author

Murilo Santana <<mvrilo@gmail.com>>
