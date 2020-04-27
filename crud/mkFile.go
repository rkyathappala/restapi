///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 5 March 2020
//
// Make file routine for CRUD file managment REST API
///////////////////////////////////////////////////////////////////////////////

package crud

import (
  "fmt"
  //"path"
  //"strings"
  "strconv"
  "net/http"
  //"encoding/json"

  "github.com/gorilla/mux"
  "restapi/element"
)

// MkFile makes a file
func MkFile(w http.ResponseWriter, r *http.Request) {
  // Get the path (and name) for the file to be saved
  params := mux.Vars(r)
  name := params["file"]
  dirID, err := strconv.Atoi(params["parent"])
  if err != nil {
    fmt.Fprintf(w, "MkFile: %s\n", err)
  }
  fmt.Fprintf(w, "MkFile: will make file \"%s\" in directory with ID = %v\n", name, dirID)

  file := new(element.Element)

  // Get the JSON data
  err = file.SetFileData(r)
  if err != nil {
    fmt.Fprintf(w, "MkFile: %s\n", err)
    return
  }

  // Add to filetree

  fmt.Println(outputAsJSON(file))
}
