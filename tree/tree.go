package tree

import "github.com/tel21a-inf2/avltrees/element"

// Container-Typ für einen Baum.
type Tree struct {
	root *element.Element
}

// Konstruktor für einen Baum.
func NewTree() *Tree {
	return &Tree{element.NewElement()}
}

// Fügt ein Element hinzu.
func (t *Tree) Add(key int) {
	t.root.Add(key)
}

// String-Ausgabe.
func (t Tree) String() string {
	return element.TreeAsMermaid(t.root)
}
