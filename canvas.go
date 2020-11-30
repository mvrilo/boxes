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

		lines := bytes.Split(final, []byte("\n"))
		nextLines := bytes.Split(data, []byte("\n"))
		count := len(lines[0])

		for i, line := range nextLines {
			if i >= len(lines) {
				lines = append(lines, bytes.Repeat([]byte(" "), count))
			}

			lines[i] = append(lines[i], ' ')
			lines[i] = append(lines[i], line...)
		}

		final = bytes.Join(lines, []byte("\n"))
	}

	return
}
