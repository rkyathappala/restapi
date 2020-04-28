package element

import (
  "fmt"
)

// toDir converts elements types to direcory objects
func (el *Element) toDir() (*dir, error) {
  if el.IsDir() {
    d := new(dir)
    d.Name = el.name
    d.Subtree = el.subtree
    return d, nil
  }

  return nil, fmt.Errorf("convert: \"%s\" connot be converted to directory", el.name)
}

// ToFile converts element objects to file objects
func (el *Element) toFile() (*file, error) {
  if el.IsFile() {
    f := new(file)
    f.Name = el.name
    f.Content = el.content
    return f, nil
  }

  return nil, fmt.Errorf("toFile: \"%s\" connot be converted to file", el.name)
}

func (d *dir) toElement() *Element {
  el := new(Element)
  el.name = d.Name
  el.parent = d.Parent
  el.subtree = d.Subtree
  return el
}

func (f *file) toElement() *Element {
  el := new(Element)
  el.name = f.Name
  el.parent = f.Parent
  el.content = f.Content
  return el
}
