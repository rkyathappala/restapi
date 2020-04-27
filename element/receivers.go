package element

import "fmt"

// Name returns name of a filetree element
func (el *Element) Name() string {
  return el.name
}

// Parent returns a pointer to an interface of the parent directory
func (el *Element) Parent() *Element {
  par, _ := LookupByID(el.parent)  
  return par
}

// Subtree returns all elements inside a directory element
func (el Element) Subtree() ([]Element, error) {
  if el.IsDir() {
    return el.subtree, nil
  }

  return nil, fmt.Errorf("subtree: Cannot get subtree for \"%s\"", el.name)
}

// Content returns the content of a file element
func (el Element) Content() (string, error) {
  if el.IsFile() {
    return el.content, nil
  }

  return "", fmt.Errorf("content: Cannot get content for \"%s\"", el.name)
}
