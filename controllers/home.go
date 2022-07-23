package controllers

import (
	"log"
	"net/http"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

type HomePageData struct {
	Header
	Servers []ServerOnline
}

var HomePage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	var servers []Server
	errSs := db.Model(&servers).Select()
	if errSs != nil {
		log.Println("Error getting servers ", errSs)
	}

	var personal = GetPersonal(db, PersonalFilter{Name: "", Rank: "", Status: "active"})
	var online []Online
	errS := db.Model(&online).Select()
	if errS != nil {
		log.Println("Error getting online data", errS)
	}
	var s []ServerOnline

	for _, so := range servers {
		s = append(s, h.InServerN(so, online, personal))
	}

	data := HomePageData{
		Servers: s,
		Header: Header{
			Page: "home",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "home.tpl", data)
})
