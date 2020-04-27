package element

import (
  "fmt"
  "time"
  "net/http"
  "encoding/json"

  "github.com/golang/gddo/httputil/header"
  "restapi/utils"
)

// SetCreated sets the timestamp for when the element was created
func (el *Element) SetCreated() {
  el.created = utils.Timestamp(time.Now())
  el.edited = utils.Timestamp(time.Now())
}

// SetEdited sets the timestamp for when the element was created
func (el *Element) SetEdited() {
  el.edited = utils.Timestamp(time.Now())
}

// SetFileData sets file name and content
func (el *Element) SetFileData(r *http.Request) error {
  // Ensure client is sending data by type JSON
  if r.Header.Get("Content-Type") != "" {
      value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
      if value != "application/json" {
          return fmt.Errorf("GetFileData: Unsupported media type: Content-Type not set to \"application/json\"")
      }
  }

  file := new(file)
  err := json.NewDecoder(r.Body).Decode(file)
  if err != nil {
    return fmt.Errorf("GetFileData: %s", err)
  }

  *el = *(file.toElement())
  el.SetCreated()

  return nil
}

// SetDirData sets directory name and subtree
func (el *Element) SetDirData(r *http.Request) error {
  // Ensure client is sending data by type JSON
  if r.Header.Get("Content-Type") != "" {
      value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
      if value != "application/json" {
          return fmt.Errorf("SetFileData: Unsupported media type: Content-Type not set to \"application/json\"")
      }
  }

  dir := new(dir)
  err := json.NewDecoder(r.Body).Decode(dir)
  if err != nil {
    return fmt.Errorf("SetDirData: %s", err)
  }

  *el = *(dir.toElement())

  return nil
}

// SetParent sets the parent field for file or Directory
func (el *Element) SetParent(p *Element) error {
  el.parent = p.id

  return nil
}
