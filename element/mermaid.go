package element

import "fmt"

const edgeFormatString = "%s%d --> %d\n"

// Liefert eine Mermaid-Repräsentation der Kanten
// des Elements zu seinen Kindern.
func (element Element) MermaidEdges(padding string) string {
	result := ""
	if !element.left.IsEmpty() {
		result += fmt.Sprintf(edgeFormatString, padding, element.key, element.left.key)
	}
	if !element.right.IsEmpty() {
		result += fmt.Sprintf(edgeFormatString, padding, element.key, element.right.key)
	}
	return result
}

// Liefert alle Mermaid-Kanten-Repräsentationen des Baumes.
func (element Element) AllMermaidEdges(padding string) string {
	if element.IsEmpty() {
		return ""
	}
	return element.MermaidEdges(padding) +
		element.left.AllMermaidEdges(padding) +
		element.right.AllMermaidEdges(padding)
}

// Liefert eine komplette Mermaid-Repräsentation des Baumes.
func TreeAsMermaid(element *Element) string {
	return fmt.Sprintf("graph TD\n%s", element.AllMermaidEdges("  "))
}
