package boxes

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
