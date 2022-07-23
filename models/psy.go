package models

import (
	"time"
	"html/template"
	"strings"
	"fmt"
)

type Psy struct {
	Id           int64
	Name         string
	Date         string
	Conclusion   string
	Practitioner string
}

func (r Psy) GetDate() string {
	t, e := time.Parse("2006-01-02", r.Date)
	if e != nil {
		return ""
	}
	return t.Format("02.01.2006")
}

func (r Psy) GetConclusion() template.HTML {
	return template.HTML(strings.Replace(r.Conclusion, "\n", "<br>", -1))
}

func (r Psy) GetFormMsg() string {
	return fmt.Sprintf("         Заключение психиотрического отдела №%d\n\n%s\n\n\n%s               Пациент - %s\n\n          Дата - %s", r.Id, r.Conclusion, r.Practitioner, r.Name, r.GetDate())
}
