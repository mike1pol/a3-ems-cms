package models

import (
	"time"
	"strings"
	"html/template"
	"fmt"
)

type RefType int

const (
	RefTypeWork   RefType = 1
	RefTypeWeapon RefType = 2
)

type RefDB struct {
	Id           int64
	Name         string
	Type         RefType
	Date         string
	Conclusion   string
	Practitioner string
}

func (r RefDB) GetDate() string {
	t, e := time.Parse("2006-01-02", r.Date)
	if e != nil {
		return ""
	}
	return t.Format("02.01.2006")
}

func (r RefDB) GetType() string {
	if r.Type == 1 {
		return "На работу"
	} else {
		return "На оружие"
	}
}

func (r RefDB) GetConclusion() template.HTML {
	return template.HTML(strings.Replace(r.Conclusion, "\n", "<br>", -1))
}

func (r RefDB) GetFormMsg() string {
	return fmt.Sprintf("              Справка №%d\n\n%s\n\n\n%s               Пациент - %s\n\n          Дата - %s", r.Id, r.Conclusion, r.Practitioner, r.Name, r.GetDate())
}
