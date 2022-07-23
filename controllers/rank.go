package controllers

import (
	"net/http"
	"log"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
	"strconv"
	"github.com/gorilla/mux"
)

var EditRank = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "admin", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	rank := Rank{
		Id: id,
	}
	errS := db.Select(&rank)
	if errS != nil || rank.Name == "" {
		http.Redirect(w, r, "/rank/%d?error=Not found", 303)
		return
	}
	name := r.FormValue("name")
	sort, es := strconv.ParseInt(r.FormValue("sort"), 10, 64)
	next, en := strconv.ParseInt(r.FormValue("next"), 10, 64)
	if es != nil || en != nil || name == "" {
		http.Redirect(w, r, "/rank/%d?error=Error params", 303)
		return
	}
	rank.Name = name
	rank.Sort = int(sort)
	rank.Next = int(next)
	errU := db.Update(&rank)
	if errU != nil {
		log.Println("Error update rank: ", errU)
		http.Redirect(w, r, "/rank?error=Error update rank", 303)
		return
	}
	http.Redirect(w, r, "/rank", 303)
})

var NewRank = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	name := r.FormValue("name")
	sort, es := strconv.ParseInt(r.FormValue("sort"), 10, 64)
	next, en := strconv.ParseInt(r.FormValue("next"), 10, 64)
	if es != nil || en != nil || name == "" {
		http.Redirect(w, r, "/rank/%d?error=Error params", 303)
		return
	}
	ins := Rank{
		Name: name,
		Sort: int(sort),
		Next: int(next),
	}
	errI := db.Insert(&ins)
	if errI != nil {
		log.Println("Error insert new rank: ", errI)
		http.Redirect(w, r, "/rank?error=Error create new rank", 303)
		return
	}
	http.Redirect(w, r, "/rank", 303)
})

type RanksPageData struct {
	Header
	List []Rank
}

var ViewRanks = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	var ranks []Rank
	errSr := db.Model(&ranks).Select()
	if errSr != nil {
		log.Println("Error getting ranks ", errSr)
	}

	ranks = h.RankSort(ranks, true)

	data := RanksPageData{
		List: ranks,
		Header: Header{
			Title: "Ранги",
			Page:  "admin",
			User:  currentUser,
		},
	}
	h.RenderTemplate(w, "ranks.tpl", data)
})

var DeleteRank = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "admin", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	rank := Rank{
		Id: id,
	}
	errD := db.Delete(&rank)
	if errD != nil {
		log.Println("Error get rank: ", errD)
		w.Write([]byte("ERROR"))
		return
	}
	w.Write([]byte("OK"))
})
