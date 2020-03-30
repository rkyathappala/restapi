///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 30 January 2020
//
// Logic for file managment REST API
///////////////////////////////////////////////////////////////////////////////

package main

import (
  //"encoding/json"
  //"fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  //"strconv"

  //"restapi/utils"
  "restapi/crud"
)

// Entry point for application
func main() {
  // Init router
  router := mux.NewRouter()

  // Route handlers & endpoints
  router.HandleFunc("/make/{file}", crud.MkFile).Methods("POST")
  router.HandleFunc("/make/{parent}/{file}", crud.MkFile).Methods("POST")
  log.Fatal(http.ListenAndServe(":8000", router))
}
