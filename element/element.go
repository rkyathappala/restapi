///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 30 January 2020
//
// Interface for element type - building block for REST API
///////////////////////////////////////////////////////////////////////////////

package element

// Element interface
// Able to interact with both file/directory structures
type Elem interface {
  // bools.go
  Empty() bool
  IsFile() bool
  IsDir() bool

  // receivers.go
  Name() string
  Parent() *Element
  Subtree() []Element
}

// Element struct
// Polymorphic type to keep files and directories properly seperate, while
// making searching easier
type Element struct {
  parent *dir
  name string

  // file properties
  content string

  // directory properties
  subtree []Element
}

// File struct
// Able to be addressed (name & ID), and contains Content
type file struct {
  parent *dir
  name string
  content string
}

// Directory struct
// Able to be addressed, and contian multiple files & folders
type dir struct {
  parent *dir
  name string
  subtree []Element
}
