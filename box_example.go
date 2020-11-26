package boxes

// Output:
// ___________________
// |                 |
// |   hello world   |
// |_________________|
// ___________________
// |                 |
// |   hello world   |
// |   hehe          |
// |_________________|

func ExampleBox() {
	box, err := New().Padding(2).Write([]byte("hello world"))
	if err != nil {
		panic(err)
	}
	println(string(box.Render()))

	box, err = box.Write([]byte("\nhehe"))
	if err != nil {
		panic(err)
	}
	println(string(box.Render()))
}
