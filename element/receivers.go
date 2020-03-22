package element

import "fmt"

// Parent returns a pointer to an interface of the parent directory
func (el *Element) Parent() *Element {
  return el.parent.toElement()
}

// Subtree returns all elements inside a directory element
func (el Element) Subtree() ([]Element, error) {
  if el.IsDir() {
    return el.subtree, nil
  }

  return nil, fmt.Errorf("subtree: Cannot get subtree for \"%s\"", el.name)
}
