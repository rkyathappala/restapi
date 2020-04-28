///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 5 March 2020
//
// Make file routine for CRUD file managment REST API
///////////////////////////////////////////////////////////////////////////////

package crud

import (
  //"fmt"
  //"path"
  //"strings"
  "net/http"
  //"encoding/json"

  //"github.com/gorilla/mux"
  "restapi/element"
  "restapi/utils"
)

// MkFile makes a file
func MkFile(w http.ResponseWriter, r *http.Request) {
  // Get the path (and name) for the file to be saved
  //params := mux.Vars(r)
  //name := params["file"]

  file := new(element.Element)

  // Get the JSON data
  err := file.SetFileData(r)
  if err != nil {
    utils.Log("MkFile: %s\n", err)
    return
  }

  // Add to filetree

  utils.Log("File created:\n" + file.OutputJSON())
}
