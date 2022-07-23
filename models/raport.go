package models

import (
	"html/template"
	"strings"
	"time"
)

// RaportComment Raport Comments model
type RaportComment struct {
	ID       int64
	Date     string
	Raport   Raport
	RaportID int64
	Author   Personal
	AuthorID int64
	Message  string
}

// GetDate get date
func (c RaportComment) GetDate() string {
	t, e := time.Parse("2006-01-02 15:04", c.Date)
	if e != nil {
		return ""
	}
	return t.Format("15:04 02.01.2006")
}

// GetMessage get message body
func (c RaportComment) GetMessage() template.HTML {
	return template.HTML(strings.Replace(c.Message, "\n", "<br>", -1))
}

// Raport Raport model
type Raport struct {
	ID       int64
	From     Personal
	FromID   int64
	To       Personal
	ToID     int64
	Subject  string
	Body     string
	Status   int
	Date     string
	Comments []RaportComment
}

// GetBody Get raport body
func (b Raport) GetBody() template.HTML {
	return template.HTML(strings.Replace(b.Body, "\n", "<br>", -1))
}

// GetCountComments get length comments
func (b Raport) GetCountComments() int {
	return len(b.Comments)
}

// GetDate get raport human date
func (b Raport) GetDate() string {
	t, e := time.Parse("2006-01-02", b.Date)
	if e != nil {
		return ""
	}
	return t.Format("02.01.2006")
}

// GetStatus Get status
func (b Raport) GetStatus() string {
	if b.Status == 2 {
		return "В работе"
	} else if b.Status == 3 {
		return "Закрыт"
	}
	return "Новый"
}
