package element

// Empty returns true if element object has not been populated
func (el *Element) Empty() bool {
  return len(el.content) == 0 && len(el.subtree) == 0
}

// IsFile retunrs true if element object has content
func (el *Element) IsFile() bool {
  return len(el.content) > 0 && len(el.subtree) == 0
}

// IsDir returns true if element object has other elements in its subtree
func (el *Element) IsDir() bool {
  return len(el.subtree) > 0 && len(el.content) == 0
}
