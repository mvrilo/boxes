package main

import "github.com/mvrilo/boxes"

func main() {
	canvas := boxes.NewCanvas()

	box, _ := boxes.New().Padding(3).WriteString("hello world\n hehehe")
	canvas.AddBox(box)

	box, _ = boxes.New().Padding(3).WriteString("from boxes\n\nyeeha\n:)")
	canvas.AddBox(box)

	println(string(canvas.Render()))

	println(string(canvas.HorizontalRender()))
}
