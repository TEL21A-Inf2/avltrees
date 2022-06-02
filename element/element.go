package element

// Typ für Elemente mit Key, Höhe und Kindern.
type Element struct {
	key         int
	height      int
	left, right *Element
}

// Konstruktor für Elemente
func NewElement() *Element {
	return &Element{0, 0, nil, nil}
}

// Liefert true, falls das Element leer ist.
func (element Element) IsEmpty() bool {
	return element.left == nil && element.right == nil
}

// Setzt die Daten des Elements.
func (element *Element) SetData(key int) {
	if element.IsEmpty() {
		element.left = NewElement()
		element.right = NewElement()
		element.height = 1
	}
	element.key = key
}

// Fügt ein neues Element zum Baum hinzu.
func (element *Element) Add(key int) {
	// Wenn wir ein Blatt gefunden haben, fügen wir dort ein.
	if element.IsEmpty() {
		element.SetData(key)
		return
	}

	// Rekursiver Abstieg.
	if key < element.key {
		element.left.Add(key)
	} else {
		element.right.Add(key)
	}

	// Update der Höhe.
	if element.left.height > element.right.height {
		element.height = element.left.height + 1
	} else {
		element.height = element.right.height + 1
	}
}

// String-Ausgabe des Elements:
func (element Element) String() string {
	return element.MermaidEdges("")
}
