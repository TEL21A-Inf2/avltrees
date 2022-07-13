package element

import "fmt"

func ExampleRotateLeft() {
	t1 := NewElement()
	t1.Add(42)
	t1.Add(77)
	t1.Add(85)

	fmt.Println(t1.AllMermaidEdges(""))
	t1 = RotateLeft(t1)
	fmt.Println(t1.AllMermaidEdges(""))

	// Output:
	// 42 --> 77
	// 77 --> 85
	//
	// 77 --> 42
	// 77 --> 85
}

func ExampleRotateRight() {
	t1 := NewElement()
	t1.Add(42)
	t1.Add(23)
	t1.Add(15)

	fmt.Println(t1.AllMermaidEdges(""))
	t1 = RotateRight(t1)
	fmt.Println(t1.AllMermaidEdges(""))

	// Output:
	// 42 --> 23
	// 23 --> 15
	//
	// 23 --> 15
	// 23 --> 42
}

func ExampleRotateLeftRight() {
	t1 := NewElement()
	t1.Add(42)
	t1.Add(23)
	t1.Add(77)
	t1.Add(60)
	t1.Add(65)

	fmt.Println(t1.AllMermaidEdges(""))
	t1 = RotateLeftRight(t1)
	fmt.Println(t1.AllMermaidEdges(""))

	// Output:
	// 42 --> 23
	// 42 --> 77
	// 77 --> 60
	// 60 --> 65
	//
	// 60 --> 42
	// 60 --> 77
	// 42 --> 23
	// 77 --> 65
}

func ExampleRotateRightLeft() {
	t1 := NewElement()
	t1.Add(42)
	t1.Add(23)
	t1.Add(77)
	t1.Add(30)
	t1.Add(25)

	fmt.Println(t1.AllMermaidEdges(""))
	t1 = RotateLeftRight(t1)
	fmt.Println(t1.AllMermaidEdges(""))

	// Output:
	// 42 --> 23
	// 42 --> 77
	// 23 --> 30
	// 30 --> 25
	//
	// 30 --> 23
	// 30 --> 42
	// 23 --> 25
	// 42 --> 77
}
