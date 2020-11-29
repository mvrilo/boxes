# boxes

[![GoDoc](https://godoc.org/github.com/mvrilo/boxes?status.svg)](https://godoc.org/github.com/mvrilo/boxes)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvrilo/boxes?_=1)](https://goreportcard.com/report/github.com/mvrilo/boxes?_=1)

Boxes is a simple ascii box builder.

## Example

```go
package main

import "github.com/mvrilo/boxes"

func main() {
	box, _ := boxes.New().Padding(3).WriteString("hey there!")
	println(string(box.Render()))
}
```

```
$ go run examples/box/main.go
.----------------.
|                |
|   hey there!   |
|                |
'----------------'
```


## License

MIT

## Author

Murilo Santana <<mvrilo@gmail.com>>
