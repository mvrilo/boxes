package boxes

import (
	"bytes"
	"math"

	"github.com/rs/xid"
)

const (
	bottomCorner = "'"
	topCorner    = "."
	row          = "-"
	column       = "|"
	space        = " "
	newLine      = "\n"
)

// Box is a container for a content
type Box struct {
	id      []byte
	padding int
	content *bytes.Buffer
}

func New() *Box {
	return &Box{
		id:      xid.New().Bytes(),
		content: bytes.NewBuffer(nil),
	}
}

func (b *Box) Padding(padding int) *Box {
	if padding < 1 {
		b.padding = 1
	}
	b.padding = padding
	return b
}

func (b *Box) Write(content []byte) (*Box, error) {
	_, err := b.content.Write(content)
	return b, err
}

func (b *Box) WriteString(content string) (*Box, error) {
	return b.Write([]byte(content))
}

func (b *Box) Bytes() []byte {
	return b.content.Bytes()
}

func (b *Box) Render() []byte {
	var lineSize int

	content := b.content.Bytes()
	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines {
		if len(line) > lineSize {
			lineSize = len(line)
		}
	}

	padding := b.padding
	lineSize = lineSize + (padding * 2)
	fullRow := bytes.Repeat([]byte(row), lineSize)
	emptyRow := bytes.Repeat([]byte(" "), lineSize)

	final := bytes.NewBuffer(nil)
	final.Write([]byte(topCorner))
	final.Write(fullRow)
	final.Write([]byte(topCorner))
	final.WriteString(newLine)

	linePadding := b.padding
	if linePadding > 1 {
		linePadding = int(math.Ceil(float64(linePadding) * 0.3))
	}

	for i := 0; i < linePadding; i++ {
		final.WriteString(column)
		final.Write(emptyRow)
		final.WriteString(column)
		final.WriteString(newLine)
	}

	for _, line := range lines {
		final.WriteString(column)
		final.Write(emptyRow[0:b.padding])
		final.Write(line)
		final.Write(emptyRow[0 : len(emptyRow)-len(line)-(padding)])
		final.WriteString(column)
		final.WriteString(newLine)
	}

	for i := 0; i < linePadding; i++ {
		final.WriteString(column)
		final.Write(emptyRow)
		final.WriteString(column)
		final.WriteString(newLine)
	}

	final.Write([]byte(bottomCorner))
	final.Write(fullRow)
	final.Write([]byte(bottomCorner))

	return final.Bytes()
}
