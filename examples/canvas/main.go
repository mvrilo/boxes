package main

import "github.com/mvrilo/boxes"

func main() {
	canvas := boxes.NewCanvas()

	box, _ := boxes.New().Padding(3).WriteString("boxes is a simple\nascii box builder")
	canvas.AddBox(box)

	box, _ = boxes.New().Padding(3).WriteString(":)")
	canvas.AddBox(box)

	box, _ = boxes.New().Padding(3).WriteString("1\n2\n3\n4\n5")
	canvas.AddBox(box)

	println(string(canvas.HorizontalRender()))
}
