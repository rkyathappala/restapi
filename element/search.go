///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 5 March 2020
//
// Search REST API for a file/directory
///////////////////////////////////////////////////////////////////////////////

package element

import (
  "fmt"
)

// SearchFor will find and return the address of an Element whose name
// matches the search term
func (origin *Element) SearchFor(target string) (e *Element, err error) {
  if origin.IsDir() {
    subs, err := origin.Subtree()
    if err != nil {
      return nil, fmt.Errorf("search: %s", err)
    }

    for _, el := range subs {
      if el.Name() == target {
        return &el, nil
      } else if el.IsDir() {
        go el.SearchFor(target)
      }
    }

    return nil, fmt.Errorf("search: target not found")
  }

  return nil, fmt.Errorf("search: origin of search is not a directory")
}
