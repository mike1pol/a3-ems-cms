package models

type Duty struct {
	Id         int64
	Personal   Personal
	PersonalId int64
	Zone       int `sql:",notnull"`
}
