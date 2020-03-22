package element

import (
  "fmt"
)

// toDir converts elements types to direcory types
func (el *Element) toDir() (*dir, error) {
  if len(el.content) == 0 {
    d := new(dir)
    d.name = el.name
    d.parent = el.parent
    d.subtree = el.subtree
    return d, nil
  }

  return nil, fmt.Errorf("convert: \"%s\" connot be converted to directory", el.name)
}

func (el *Element) toFile() (*file, error) {
  if len(el.subtree) == 0 {
    f := new(file)
    f.name = el.name
    f.parent = el.parent
    f.content = el.content
    return f, nil
  }

  return nil, fmt.Errorf("convert: \"%s\" connot be converted to file", el.name)
}

func (d *dir) toElement() *Element {
  el := new(Element)
  el.name = d.name
  el.parent = d.parent
  el.subtree = d.subtree
  return el
}
