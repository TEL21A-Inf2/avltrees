package tree

func testTree1() *Tree {
	result := NewTree()
	result.Add(42)
	result.Add(23)
	result.Add(15)
	result.Add(77)
	result.Add(98)
	result.Add(60)

	return result
}

func testTree2() *Tree {
	result := NewTree()
	return result
}

func testTree3() *Tree {
	result := NewTree()

	result.Add((1))
	result.Add((2))
	result.Add((3))
	result.Add((4))
	result.Add((5))
	result.Add((6))

	return result
}
