package api

import (
	"net/http"
	"log"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

type OnlineAPIData struct {
	Status  string
	Message string
	Count   int
	Servers []ServerOnline
}

var OnlineAPIGet = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()

	var servers []Server
	errSs := db.Model(&servers).Select()
	if errSs != nil {
		log.Println("Error getting servers ", errSs)
		h.SendJson(w, OnlineAPIData{Status: "error", Message: "Error get servers"}, http.StatusInternalServerError)
		return
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

	data := OnlineAPIData{
		Status:  "success",
		Count:   len(s),
		Servers: s,
	}
	h.SendJson(w, data, http.StatusOK)
})
