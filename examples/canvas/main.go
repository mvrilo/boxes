package main

import "github.com/mvrilo/boxes"

func main() {
	canvas := boxes.NewCanvas()

	box, _ := boxes.New().Padding(3).WriteString("hello world")
	canvas.AddBox(box)

	box, _ = boxes.New().Padding(3).WriteString("boxes is a simple\nascii box builder\n\n:)")
	canvas.AddBox(box)

	println(string(canvas.HorizontalRender()))
}
