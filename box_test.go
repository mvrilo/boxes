package boxes

import (
	"bytes"
	"testing"
)

func TestPadding(t *testing.T) {
	box := New()
	if box.padding != 0 {
		t.Errorf("got %d, expected %d", box.padding, 0)
	}

	box.Padding(3)
	if box.padding != 3 {
		t.Errorf("got %d, expected %d", box.padding, 3)
	}
}

func TestWrite(t *testing.T) {
	box := New()
	box.Padding(3)

	content := box.Content.Bytes()
	n := len(content)
	if n != 0 {
		t.Errorf("got %d, expected %d", n, 0)
	}

	_, err := box.Write([]byte("hi"))
	if err != nil {
		t.Error(err)
	}

	content = box.Content.Bytes()
	if !bytes.Equal(content, []byte("hi")) {
		t.Errorf("got %s, expected %v", content, []byte("hi"))
	}
}
