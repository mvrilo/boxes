package main

import "github.com/mvrilo/boxes"

func main() {
	canvas := boxes.NewCanvas()

	box, _ := boxes.New().Padding(3).WriteString("hello world,\n")
	canvas.AddBox(box)

	box, _ = boxes.New().Padding(3).WriteString("from boxes\n")
	canvas.AddBox(box)

	println(string(canvas.Render()))
}
