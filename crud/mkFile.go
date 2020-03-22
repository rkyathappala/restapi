///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 5 March 2020
//
// Make file routine for file managment REST API
///////////////////////////////////////////////////////////////////////////////

package crud

import (
  "encoding/json"
  "restapi/element"
)

func mkFile(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  file := new(element.File)


}
