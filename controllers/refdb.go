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

type RefDBPageData struct {
	Header
	Type RefType
	Name string
	List []RefDB
}

var RefDBPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var name = r.FormValue("name")
	var refType RefType
	if r.FormValue("type") == "2" {
		refType = RefTypeWeapon
	} else if r.FormValue("type") == "1" {
		refType = RefTypeWork
	}
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	var refDB []RefDB
	if refType == 0 {
		errS := db.Model(&refDB).Order("date desc").Where("lower(name) like lower(?)", "%"+name+"%").Select()
		if errS != nil {
			log.Println("Error getting refDB ", errS)
		}
	} else {
		errS := db.Model(&refDB).Order("date desc").Where("type = ? and lower(name) like lower(?)", refType, "%"+name+"%").Select()
		if errS != nil {
			log.Println("Error getting refDB ", errS)
		}
	}

	data := RefDBPageData{
		List: refDB,
		Name: name,
		Type: refType,
		Header: Header{
			Page: "refDB",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "refdb.tpl", data)
})

var NewRefDB = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var refType RefType
	if r.FormValue("type") == "2" {
		refType = RefTypeWeapon
	} else {
		refType = RefTypeWork
	}
	refDB := RefDB{
		Name:         r.FormValue("name"),
		Type:         refType,
		Date:         r.FormValue("date"),
		Conclusion:   r.FormValue("conclusion"),
		Practitioner: r.FormValue("practitioner"),
	}

	err := db.Insert(&refDB)

	if err != nil {
		http.Redirect(w, r, "/refdb?error=Error create new refDB", 303)
		return
	}

	http.Redirect(w, r, "/refdb", 303)
})

var UpdateRefDB = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "refDB", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	var refDB = RefDB{Id: id}
	errS := db.Select(&refDB)
	if errS != nil {
		log.Println("Error get refDB", errS)
		http.Redirect(w, r, fmt.Sprintf("/refdb?error=error select"), 303)
		return
	}
	name := r.FormValue("name")
	date := r.FormValue("date")
	conclusion := r.FormValue("conclusion")
	practitioner := r.FormValue("practitioner")
	refDB.Name = name
	if r.FormValue("type") == "2" {
		refDB.Type = RefTypeWeapon
	} else {
		refDB.Type = RefTypeWork
	}
	refDB.Date = date
	refDB.Conclusion = conclusion
	refDB.Practitioner = practitioner
	errU := db.Update(&refDB)
	if errU != nil {
		log.Println("Error update refDB: ", errU)
		http.Redirect(w, r, "/refdb/?error=Error update data", 303)
		return
	}
	http.Redirect(w, r, "/refdb", 303)
})

var DeleteRefDB = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "refDB", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	refDB := RefDB{
		Id: id,
	}
	errD := db.Delete(&refDB)
	if errD != nil {
		log.Println("Error get refDB: ", errD)
		w.Write([]byte("ERROR"))
		return
	}
	w.Write([]byte("OK"))
})
