package element

import "fmt"

func ExampleElement() {
	fmt.Println(testTree1().AllMermaidEdges(""))
	fmt.Println(testTree2().AllMermaidEdges("")) // Keine Ausgabe erwartet.
	fmt.Println(testTree3().AllMermaidEdges(""))

	// Output:
	// 42 --> 23
	// 42 --> 77
	// 23 --> 15
	// 77 --> 60
	// 77 --> 98
	//
	// 1 --> 2
	// 2 --> 3
	// 3 --> 4
	// 4 --> 5
	// 5 --> 6
}

func ExampleElement_height() {
	fmt.Println(testTree1().height)
	fmt.Println(testTree2().height)
	fmt.Println(testTree3().height)

	// Output:
	// 3
	// 0
	// 6
}
