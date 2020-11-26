package boxes

import (
	"bytes"
)

const (
	row     = "_"
	column  = "|"
	space   = " "
	newLine = "\n"
)

// Box is a container for a Content
type Box struct {
	padding int
	Content *bytes.Buffer
}

func New() *Box {
	return &Box{
		Content: bytes.NewBuffer(nil),
	}
}

func (b *Box) Padding(padding int) *Box {
	b.padding = padding
	return b
}

func (b *Box) Write(content []byte) (*Box, error) {
	_, err := b.Content.Write(content)
	return b, err
}

func (b *Box) Render() []byte {
	content := b.Content.Bytes()
	final := bytes.NewBuffer(nil)
	padding := b.padding
	if padding > 0 {
		padding = 1
	}
	if padding > 1 {
		padding = padding / 2
	}

	var maxLineSize int

	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines {
		if len(line) > maxLineSize {
			maxLineSize = len(line)
		}
	}

	lineSize := maxLineSize + padding
	linepadding := bytes.Repeat([]byte(" "), padding+2)

	firstLine := bytes.Repeat([]byte(row), lineSize+(padding+2)+4)
	lastLine := bytes.Repeat([]byte(row), lineSize+(padding+2)+2)
	emptyLine := bytes.Repeat([]byte(" "), lineSize+(padding+4))

	final.Write(firstLine)
	final.WriteString(newLine)

	for i := 0; i < padding; i++ {
		final.WriteString(column)
		final.Write(emptyLine)
		final.WriteString(column)
		final.WriteString(newLine)
	}

	for _, line := range lines {
		final.WriteString(column)
		final.Write(linepadding)
		final.Write(line)
		final.Write(emptyLine[0 : len(firstLine)-len(line)-len(linepadding)-2])
		final.WriteString(column)
		final.WriteString(newLine)
	}

	for i := 0; i < padding; i++ {
		final.WriteString(column)
		final.Write(lastLine)
		final.WriteString(column)
	}

	return final.Bytes()
}
