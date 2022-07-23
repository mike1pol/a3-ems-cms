package models

import (
	"fmt"
	"math"
	"time"
)

type OdbFilter struct {
	Name    string
	SDate   string
	EDate   string
	UnGroup bool
}

type Odb struct {
	Id    int64
	Name  string
	Date  string
	First time.Time
	Last  time.Time
	Time  float64
}

type OdbMed struct {
	Odb
	Personal PersonalJson
}

func (o Odb) GetTime() string {
	h := math.Floor(o.Time / 60)
	m := math.Floor(o.Time - (h * 60))
	return fmt.Sprintf("%.0fh %.0fm", h, m)
}
func (o Odb) FirstSeen() string {
	if o.First.IsZero() {
		return ""
	}
	return o.First.Format("02.01.2006 15:04")
}
func (o Odb) LastSeen() string {
	return o.Last.Format("02.01.2006 15:04")
}
