package helpers

import (
  "time"
)

func FormatISO(t time.Time) string {
  return t.Format("2006-01-02")
}
