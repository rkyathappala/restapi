package crud

import (
  "strings"

  "restapi/element"
)

// PrintJSON will print element data to JSON format
func outputAsJSON(el *element.Element) (output string) {
  output = "{\n\t\"name\": " + el.Name() + "\n"
  if el.IsDir() {
    output += "\t\"subtree\": ["
    children, _ := el.Subtree()
    for _, child := range children {
      if child.Name() != children[0].Name() {
        output = strings.Join([]string{output, child.Name()}, ", ")
      } else {
        output += child.Name()
      }
    }
    output += "]\n"
  } else if el.IsFile() {
    cont, _ := el.Content()
    output += "\t\"content\": " + cont + "\n"
  }
  output += "}"

  return output
}
