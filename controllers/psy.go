package controllers

import (
	"log"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

type PsyPageData struct {
	Header
	Name string
	List []Psy
}

var PsyPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var name = r.FormValue("name")
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	var psy []Psy
	errS := db.Model(&psy).Where("lower(name) like lower(?)", "%"+name+"%").Select()
	if errS != nil {
		log.Println("Error getting psy ", errS)
	}

	data := PsyPageData{
		List: psy,
		Name: name,
		Header: Header{
			Page: "psy",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "psy.tpl", data)
})

var NewPsy = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	psy := Psy{
		Name:         r.FormValue("name"),
		Date:         r.FormValue("date"),
		Conclusion:   r.FormValue("conclusion"),
		Practitioner: r.FormValue("practitioner"),
	}

	err := db.Insert(&psy)

	if err != nil {
		http.Redirect(w, r, "/psy?error=Error create new psy", 303)
		return
	}

	http.Redirect(w, r, "/psy", 303)
})

var UpdatePsy = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "psy", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	var psy = Psy{Id: id}
	errS := db.Select(&psy)
	if errS != nil {
		log.Println("Error get psy", errS)
		http.Redirect(w, r, fmt.Sprintf("/refdb?error=error select"), 303)
		return
	}
	psy.Name = r.FormValue("name")
	psy.Date = r.FormValue("date")
	psy.Conclusion = r.FormValue("conclusion")
	psy.Practitioner = r.FormValue("practitioner")
	errU := db.Update(&psy)
	if errU != nil {
		log.Println("Error update psy: ", errU)
		http.Redirect(w, r, "/psy?error=Error update data", 303)
		return
	}
	http.Redirect(w, r, "/psy", 303)
})

var DeletePsy = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "psy", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	psy := Psy{
		Id: id,
	}
	errD := db.Delete(&psy)
	if errD != nil {
		log.Println("Error get psy: ", errD)
		w.Write([]byte("ERROR"))
		return
	}
	w.Write([]byte("OK"))
})
