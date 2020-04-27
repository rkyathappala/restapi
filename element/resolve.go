package element

import (
  "fmt"
  "strconv"
)

type ID string

var resolve map[ID]*Element
var unUsedID []ID
var _id uint64 = 0

func getNewID() ID {
  tmp := strconv.FormatUint(_id, 16)
  _id++
  return ID(tmp)
}

// Push adds element reference to the filetree stack, resolve
func Push(target *Element) error {
  id := getNewID()

  if resolve[id] == nil {
    resolve[id] = target
    target.id = id
    return nil
  }
  return fmt.Errorf("Push ID: Invalid ID: %v for %s, already in use by %s", id, target.name, resolve[id].name)
}

// Pop removes an element from the filetree stack
func Pop(target *Element) {
  delete(resolve, target.id)
  target.id = ""
}

// Lookup allows access to the name of a
func LookupByID(target ID) (*Element, error) {
  if resolve[target] != nil {
    return resolve[target], nil
  }
  return nil, fmt.Errorf("Lookup: element with ID = %s not found", target)
}
