package utils

import "time"

func Timestamp(t time.Time) string {
  t = t.Round(time.Second)
	return t.Format("2006-01-02 @ 15:04:05")
}
