package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	h "github.com/mike1pol/rms/helpers"
	m "github.com/mike1pol/rms/models"
)

type raportCreatePageData struct {
	Personal []m.PersonalJson
	Header   m.Header
	IsActive bool
}

// RaportCreatePage Raport create page
var RaportCreatePage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser m.User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(m.User)
	}
	filter := m.PersonalFilter{
		Rank:   "",
		Name:   "",
		Status: "active",
	}

	personal := m.GetPersonal(db, filter)

	data := raportCreatePageData{
		Personal: personal,
		Header: m.Header{
			Page: "raport",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "raport_create.tpl", data)
})

type raportListFilter struct {
	Status int64
}

type raportListPageData struct {
	Header   m.Header
	IsActive bool
	Filter   raportListFilter
	List     []m.Raport
}

// RaportListPage Raport list
var RaportListPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser m.User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(m.User)
	}
	var list []m.Raport

	filter := raportListFilter{
		Status: 0,
	}

	if r.FormValue("s") != "" {
		status, _ := strconv.ParseInt(r.FormValue("s"), 10, 64)
		filter.Status = status
	}

	var sel = db.Model(&list).
		Relation("From").
		Relation("To").
		Relation("Comments")

	if filter.Status > 0 {
		sel = sel.Where("raport.status = ?", filter.Status)
	}

	err := sel.Select()

	if err != nil {
		log.Println("Error getting Raport list", err)
	}

	data := raportListPageData{
		List:   list,
		Filter: filter,
		Header: m.Header{
			Page: "raport_list",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "raport_list.tpl", data)
})

// RaportCreateHandler Create raport handler
var RaportCreateHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()

	var currentUser m.User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(m.User)
	}

	sID := currentUser.SteamId

	var personal []m.Personal
	errS := db.Model(&personal).Where("steam_id = ? and status = ?", sID, "active").Select()
	if errS != nil {
		log.Println("Error getting personal ", errS)
		http.Redirect(w, r, "/raport?error=Error getting personal", 303)
		return
	}

	if len(personal) == 0 {
		log.Println("Personal not found ")
		http.Redirect(w, r, "/raport?error=Personal not found", 303)
		return
	}

	toID, _ := strconv.ParseInt(r.FormValue("toID"), 10, 64)

	raport := m.Raport{
		FromID:  personal[0].Id,
		ToID:    toID,
		Subject: r.FormValue("subject"),
		Body:    r.FormValue("body"),
		Status:  1,
		Date:    r.FormValue("date"),
	}

	err := db.Insert(&raport)

	if err != nil {
		http.Redirect(w, r, "/raport?error=Error create new raport", 303)
		return
	}

	http.Redirect(w, r, "/raport", 303)
})

type raportViewPageData struct {
	Raport m.Raport
	Header m.Header
}

// RaportViewPage Raport view page
var RaportViewPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser m.User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(m.User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", m.PageNotFound{Page: "raport_list", Header: m.Header{User: currentUser}})
		return
	}
	var raport = m.Raport{}
	err := db.Model(&raport).
		Where("raport.id = ?", id).
		Relation("To").
		Relation("From").
		Relation("Comments").
		Relation("Comments.Author").
		First()
	if err != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", m.PageNotFound{Page: "raport_list", Header: m.Header{User: currentUser}})
		return
	}

	data := raportViewPageData{
		Raport: raport,
		Header: m.Header{
			Page: "raport_list",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "raport_view.tpl", data)
})

// RaportStatusUpdateHandler Update raport handler
var RaportStatusUpdateHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	status, e1 := strconv.ParseInt(r.FormValue("status"), 10, 64)
	if e != nil || e1 != nil || id == 0 {
		http.Redirect(w, r, "/raport?error=Error getting id", 303)
		return
	}

	raport := m.Raport{ID: id}

	err := db.Select(&raport)

	if err != nil || raport.ID == 0 {
		http.Redirect(w, r, "/raport?error=Error create new raport", 303)
		return
	}

	raport.Status = int(status)

	errU := db.Update(&raport)

	if errU != nil {
		log.Println("Error update status ", errU)
		http.Redirect(w, r, fmt.Sprintf("/raport/%d?error=Error update status", id), 303)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/raport/%d", id), 303)
})

// RaportCommentHandler Add comment handler
var RaportCommentHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()

	var currentUser m.User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(m.User)
	}

	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		http.Redirect(w, r, "/raport?error=Error getting id", 303)
		return
	}

	sID := currentUser.SteamId

	var personal []m.Personal
	errS := db.Model(&personal).Where("steam_id = ? and status = ?", sID, "active").Select()
	if errS != nil {
		log.Println("Error getting personal ", errS)
		http.Redirect(w, r, "/raport?error=Error getting personal", 303)
		return
	}

	if len(personal) == 0 {
		log.Println("Error getting personal ", errS)
		http.Redirect(w, r, "/raport?error=Error getting personal", 303)
		return
	}

	time := time.Now()
	date := time.Format("2006-01-02 15:04")

	comment := m.RaportComment{
		AuthorID: personal[0].Id,
		Date:     date,
		RaportID: id,
		Message:  r.FormValue("message"),
	}

	err := db.Insert(&comment)

	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/raport/%d?error=Error create comment", id), 303)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/raport/%d", id), 303)
})
