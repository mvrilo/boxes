package main

import "github.com/mvrilo/boxes"

func main() {
	box, _ := boxes.New().Padding(3).WriteString("hey there!")
	println(string(box.Render()))
}
