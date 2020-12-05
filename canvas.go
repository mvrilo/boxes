package boxes

import "bytes"

type Canvas struct {
	Boxes []*Box
}

func NewCanvas() *Canvas {
	return &Canvas{}
}

func (c *Canvas) AddBox(box *Box) {
	c.Boxes = append(c.Boxes, box)
}

func (c *Canvas) Render() (data []byte) {
	for i, box := range c.Boxes {
		if i > 0 {
			data = append(data, byte('\n'))
		}
		data = append(data, box.Render()...)
	}
	return
}

func (c *Canvas) HorizontalRender() (final []byte) {
	for i, box := range c.Boxes {
		data := box.Render()

		if i == 0 {
			final = append(final, data...)
			continue
		}

		finalLines := bytes.Split(final, []byte("\n"))
		boxLines := bytes.Split(data, []byte("\n"))
		width := len(finalLines[0])

		rows := len(finalLines)
		if len(boxLines) > rows {
			rows = len(boxLines)
		}

		for j := 0; j <= rows; j++ {
			if j >= len(finalLines) {
				line := bytes.Repeat([]byte(" "), width)
				finalLines = append(finalLines, line)
			}

			if j >= len(boxLines) {
				line := bytes.Repeat([]byte(" "), len(boxLines[j-1]))
				boxLines = append(boxLines, line)
			}

			finalLines[j] = append(finalLines[j], ' ')
			finalLines[j] = append(finalLines[j], boxLines[j]...)
		}

		final = bytes.Join(finalLines, []byte("\n"))
	}

	return
}
