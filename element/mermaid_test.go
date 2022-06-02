package element

import "fmt"

func ExampleElement_AllMermaidEdges() {
	e1 := testTree1()

	fmt.Println(e1.AllMermaidEdges("**"))
	fmt.Println(e1.AllMermaidEdges("+++"))
	fmt.Println(e1.AllMermaidEdges(""))

	//Output:
	//**42 --> 23
	//**42 --> 77
	//**23 --> 15
	//**77 --> 60
	//**77 --> 98
	//
	//+++42 --> 23
	//+++42 --> 77
	//+++23 --> 15
	//+++77 --> 60
	//+++77 --> 98
	//
	//42 --> 23
	//42 --> 77
	//23 --> 15
	//77 --> 60
	//77 --> 98
}

func ExampleTreeAsMermaid() {
	fmt.Println(TreeAsMermaid(testTree1()))

	//Output:
	// graph TD
	//   42 --> 23
	//   42 --> 77
	//   23 --> 15
	//   77 --> 60
	//   77 --> 98
}
