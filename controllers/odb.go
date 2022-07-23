package controllers

import (
	"log"
	"net/http"
	"time"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

type OdbPageData struct {
	Header
	Filter OdbFilter
	List   []OdbMed
}

var OdbPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	var filter = OdbFilter{
		Name:    "",
		SDate:   h.FormatISO(time.Now().Add(time.Hour * 3)),
		EDate:   h.FormatISO(time.Now().Add(time.Hour * 3)),
		UnGroup: true,
	}

	if r.FormValue("s") != "" {
		filter.SDate = r.FormValue("s")
	}
	if r.FormValue("e") != "" {
		filter.EDate = r.FormValue("e")
	}
	if r.FormValue("n") != "" {
		filter.Name = r.FormValue("n")
	}
	filter.UnGroup = r.FormValue("g") == "true"

	var list []Odb

	_, err := db.Query(
		&list,
		`select *
from odbs
where
  lower(name) like lower(?) and
  date >= ? and
  date <= ? and
  time > 1
order by time desc`,
		"%"+filter.Name+"%",
		filter.SDate,
		filter.EDate,
	)

	var personal = GetPersonal(db, PersonalFilter{Name: "", Rank: "", Status: "active"})

	l := h.OdbSortByDate(h.OdbSort(h.InPersonalOdb(list, personal, filter.UnGroup)), filter.UnGroup)
	if err != nil {
		log.Println("Error getting odb ", err)
	}

	data := OdbPageData{
		Filter: filter,
		List:   l,
		Header: Header{
			Page:  "odb",
			Title: "База онлайна",
			User:  currentUser,
		},
	}
	h.RenderTemplate(w, "odb.tpl", data)
})
