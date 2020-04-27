///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 30 January 2020
//
// Interface for element type - building block for REST API
///////////////////////////////////////////////////////////////////////////////

package element

// Element struct
// Polymorphic type to keep files and directories properly seperate, while
// making searching easier
type Element struct {
  parent ID
  id ID
  name string
  created string
  edited string

  // file properties
  content string

  // directory properties
  subtree []Element
}

// File struct
// Able to be addressed (name & ID), and contains Content
type file struct {
  Name string     `json:"name"`
  Content string  `json:"content"`
}

// Directory struct
// Able to be addressed, and contian multiple files & folders
type dir struct {
  Name string       `json:"name"`
  Subtree []Element `json:"subtree"`
}
