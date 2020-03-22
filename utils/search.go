///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 5 March 2020
//
// Search REST API for a file/directory
///////////////////////////////////////////////////////////////////////////////

package utils

import (
  "fmt"
  "restapi/element"
)

// SearchFor will find and return the address of an Element whose name
// matches the search term
func (origin *element.Element) SearchFor(target string) (e *element.Element, err Error) {
  if origin.IsDir() {
    for _, el := range Elem.Subtree(origin) {
      if Elem.Name(el) == target {
        return el, nil
      } else if origin.IsDir(el) {
        go el.SearchFor(target)
      }
    }

    return nil, fmt.Errorf("search: target not found")
  }

  return nil, fmt.Errorf("search: origin of search is not a directory")
}
