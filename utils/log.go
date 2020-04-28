package utils

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "time"
  //"path/filepath"
)

var (
  // Log file
	LogFile *os.File

	// Log writer
	Logger *bufio.Writer
)

func init() {
  // Create log system
  logFileName := "restapi.log"

  logPath, err := os.Getwd()
  if err != nil {
    fmt.Println("error making log from working dir: ", err.Error())
    os.Exit(1)
  }
  //logPath = filepath.Dir(logPath)

  if _, err = os.Stat(logPath); !os.IsNotExist(err) {
    if err = os.MkdirAll(logPath, 0666); err != nil {
      fmt.Println("Could not create log folder: %s", err.Error())
      os.Exit(2)
    }
  }

  LogFile, err = os.Create(strings.Join([]string{logPath, logFileName}, "/"))
  if err != nil {
    fmt.Println("Could not create log file: %s", err.Error())
    os.Exit(3)
  }

  Logger = bufio.NewWriter(LogFile)
}


// Log is a custom printf function
func Log(format string, a ...interface{}) (n int, err error) {
  defer Logger.Flush()
  return fmt.Fprintf(Logger, Timestamp(time.Now()) + ": " + format + "\n", a...)
}
