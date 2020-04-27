///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 30 January 2020
//
// Logic for file managment REST API
///////////////////////////////////////////////////////////////////////////////

package main

import (
  //"encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  //"strconv"
  "os"
  "os/signal"
  "syscall"

  //"restapi/utils"
  "restapi/crud"
)

var (
	signalChannel chan os.Signal
	// Build information
	Version string
	Build   string
)

// Entry point for application
func main() {
  // Hanldle runtime errors
  signalChannel = make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	go signalHander()

	// Setup function used to recover from panics.
	defer func() {
		if r := recover(); r != nil {
			//mlog.Error("%s failed with panic %v", config.Configuration.ServiceName, r)
      fmt.Println("restapi.exe failed with panic %v", r)
		}
	}()

  // Init router
  router := mux.NewRouter()

  // Route handlers & endpoints
  router.HandleFunc("/make/file/{file}", crud.MkFile).Methods("POST")
  //router.HandleFunc("/make/file/{parent}/{fileName}", crud.MkFile).Methods("POST")    --> should not need parent info for creation
  //router.HandleFunc("/make/dir/{parent}/{dirName}", crud.MkFile).Methods("POST")    --> should not need parent info for creation
  log.Fatal(http.ListenAndServe(":8000", router))
}

// Helper method used to handle abort and termination signals.
func signalHander() {
  sig := <-signalChannel

  switch sig {
  case os.Interrupt, syscall.SIGABRT, syscall.SIGTERM:
    os.Exit(0)
  default:
    fmt.Println("Unknown signal caught %v\n", sig)
  }
}
