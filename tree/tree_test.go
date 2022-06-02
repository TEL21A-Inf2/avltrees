package tree

import "fmt"

func ExampleElement() {

	fmt.Println(testTree1())
	fmt.Println(testTree2())
	fmt.Println(testTree3())

	// Output:
	// graph TD
	//   42 --> 23
	//   42 --> 77
	//   23 --> 15
	//   77 --> 60
	//   77 --> 98
	//
	// graph TD
	//
	// graph TD
	//   1 --> 2
	//   2 --> 3
	//   3 --> 4
	//   4 --> 5
	//   5 --> 6
}
