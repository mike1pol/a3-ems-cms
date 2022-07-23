package models

import "time"

type Blacklist struct {
	Id       int64
	Name     string
	Reason   string
	Date     string
	IsActive bool `sql:",notnull"`
}

func (b Blacklist) GetDate() string {
	t, e := time.Parse("2006-01-02", b.Date)
	if e != nil {
		return ""
	}
	return t.Format("02.01.2006")
}
