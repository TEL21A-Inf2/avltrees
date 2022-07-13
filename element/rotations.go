package element

// Erwartet ein Wurzelelement und führt eine einfache Linksrotation damit aus.
// Liefert die neue Wurzel zurück.
func RotateLeft(element *Element) *Element {
	A := element
	B := element.right
	C := B.left

	A.right = C
	B.left = A
	return B
}

// Erwartet ein Wurzelelement und führt eine einfache Rechtsrotation damit aus.
// Liefert die neue Wurzel zurück.
func RotateRight(element *Element) *Element {
	A := element
	B := A.left
	C := B.right

	A.left = C
	B.right = A
	return B
}

// Erwartet ein Wurzelelement und führt eine Links-Rechts-Rotation damit aus.
// Liefert die neue Wurzel zurück.
func RotateLeftRight(element *Element) *Element {
	A := element
	B := element.left

	A.left = RotateLeft(B)
	A = RotateRight(A)
	return A
}

// Erwartet ein Wurzelelement und führt eine Rechts-Links-Rotation damit aus.
// Liefert die neue Wurzel zurück.
func RotateRightLeft(element *Element) *Element {
	// TODO
	return nil
}

// Berechnet den Balancefaktor des Knotens.
func BalanceFactor(element *Element) int {
	return element.right.height - element.left.height
}
