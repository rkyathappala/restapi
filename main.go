///////////////////////////////////////////////////////////////////////////////
// Rahul Kyathappala
// 30 January 2020
//
// Logic for file managment REST API
///////////////////////////////////////////////////////////////////////////////

package main

import (
  "encoding/json"
  //"fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "strconv"
)

//////
// File struct
// Able to be addressed (name & ID), and contains Content
//////
type File struct {
  ID       string  `json:"id"`
  DestID   string  `json:"dest_id"`
  Name     string  `json:"name"`
  Content  string  `json:"content"`
}

//////
// Directory struct
// Able to be addressed, and contian multiple files & folders
//////
type Dir struct {
  ID    string    `json:"id"`
  DestID   string  `json:"dest_id"`
  Name  string    `json:"name"`
  Sub  []Dir     `json:"sub"`
  Files []File    `json:"files"`
}

// Init ID counters ( ID > 1000 => file)
var folderID int = 1000
var fileID int = 500000

// Init home directory
var home = Dir {
  ID: strconv.Itoa(folderID),
  DestID: "0",
  Name: "home",
  Sub: make([]Dir, 0),
  Files: make([]File, 0)}

// Get all folders & files
func getFiles (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(home)
}

// Get 1 file based on ID
func getFile (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  // Get ID
  params := mux.Vars(r) // Get params

  // Check the destination ID
  id, err := strconv.Atoi(params["id"])
  if err == nil {
    if id < 500000 {
      log.Fatal("Invalid destination ID")
    }
  } else {
    log.Fatal(err)
  }

  //Loop through main files to find ID
  for _, item := range home.Files {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      break
    }
  }

}

// Get 1 folder based on ID
func getFolder (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  // Get ID
  params := mux.Vars(r) // Get params

  // Check the destination ID
  id, err := strconv.Atoi(params["id"])
  if err == nil {
    if id > 500000 {
      log.Fatal("Invalid destination ID")
    }
  } else {
    log.Fatal(err)
  }

  //Loop through subdirectories to find ID
  for _, item := range home.Sub {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      break
    }
  }
}

// Create new File
func makeFile (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var file File
  found := false

  // De-JSON incoming info
  _ = json.NewDecoder(r.Body).Decode(&file)

  // Set file ID
  fileID++
  file.ID = strconv.Itoa(fileID)

  // Check the destination ID
  id, err := strconv.Atoi(file.DestID)
  if err != nil {
    log.Fatal(err)
  } else if id > 500000 {
    log.Fatal("Invalid destination ID")
  }

  // If adding to home directory
  if id == 1000 {
    home.Files = append(home.Files, file)
    found = true
  } else {
    //Loop through subdirectories to find destination ID
    for i, item := range home.Sub {
      if item.ID == strconv.Itoa(id) {
        home.Sub[i].Files = append(home.Sub[i].Files, file)
        found = true
        break
      }
    }
  }

  // If given destination ID DNE - Error
  if !found {
    log.Fatal("Invalid destination ID")
  }

  json.NewEncoder(w).Encode(file)
}

// Create new folder
func makeFolder (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var folder Dir
  found := false

  // De-JSON incomming info
  _ = json.NewDecoder(r.Body).Decode(&folder)

  //Set ID
  folderID++
  folder.ID = strconv.Itoa(folderID)

  // Check the destination ID - Error
  id, err := strconv.Atoi(folder.DestID)
  if err != nil {
    log.Fatal(err)
  } else if id > 500000 {
    log.Fatal("Invalid destination ID")
  }

  // If adding to home directory
  if id == 1000 {
    home.Sub = append(home.Sub, folder)
    found = true
  } else {
    //Loop through subdirectories to find destination ID
    for i, item := range home.Sub {
      if item.ID == strconv.Itoa(id) {
        home.Sub[i].Sub = append(home.Sub[i].Sub, folder)
        found = true
        break
      }
    }
  }

  // If given destination ID DNE
  if !found {
    log.Fatal("Invalid destination ID")
  }

  json.NewEncoder(w).Encode(folder)
}


// Update file based on ID
func updateFile (w http.ResponseWriter, r *http.Request) {
  // Delete, then create
  w.Header().Set("Content-Type", "application/json")

  // Get ID
  params := mux.Vars(r) // Get params

  //Loop through to find ID
  for i, item := range home.Sub {
    if item.ID == params["id"] {
      // Delete old version
      home.Sub = append(home.Sub[:i], home.Sub[i+1:]...)

      // C
      var file File

      //De-JSON input
      _ = json.NewDecoder(r.Body).Decode(&file)

      // Reset file ID
      file.ID = params["id"]

      // Append File back to the destination ID (update file and/or move it)
      if file.DestID != strconv.Itoa(1000) {
        for j, folder := range home.Sub {
          if folder.ID == file.DestID {
            home.Sub[j].Files = append(home.Sub[j].Files, file)
          }
        }
      }
      json.NewEncoder(w).Encode(file)
      break
    }
  }
}

// Delete 1 folder/file based on ID
func deleteByID (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  // Get ID
  params := mux.Vars(r) // Get params

  // Illegal (and impossible) to delete 'home' directory
  //Loop through subdirectories to find ID
  for i, item := range home.Sub {
    if item.ID == params["id"] {
      // Slice it out
      home.Sub = append(home.Sub[:i], home.Sub[i+1:]...)
      break
    }
  }

  json.NewEncoder(w).Encode(home)
}

//////
// Entry point for application
/////
func main() {
  // Init router
  router := mux.NewRouter()

  // Test data
  fileID++
  home.Files = append(home.Files, File {
    ID: strconv.Itoa(fileID),
    DestID: "1000",
    Name: "myFile.cpp",
    Content: "#include <iostream> ..."  })

  fileID++
  home.Files = append(home.Files, File {
    ID: strconv.Itoa(fileID),
    DestID: "1000",
    Name: "myFile2.cpp",
    Content: "#include <stdio.h> ..."  })

  folderID++
  home.Sub = append(home.Sub, Dir {
    ID: strconv.Itoa(folderID),
    DestID: "1000",
    Name: "myDir",
    Sub: make([]Dir, 0),
    Files: make([]File, 0)  })

  // Route handlers & endpoints
  router.HandleFunc("/api", getFiles).Methods("GET")
  router.HandleFunc("/api/getFile/{id}", getFile).Methods("GET")
  router.HandleFunc("/api/getFolder/{id}", getFolder).Methods("GET")
  router.HandleFunc("/api/makeFile", makeFile).Methods("POST")
  router.HandleFunc("/api/makeFolder", makeFolder).Methods("POST")
  router.HandleFunc("/api/updateDile/{id}", updateFile).Methods("PUT")
  router.HandleFunc("/api/delete/{id}", deleteByID).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}
