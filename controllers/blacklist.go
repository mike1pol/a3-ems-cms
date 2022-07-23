package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

type blackList struct {
	Id       int64
	Name     string
	Reason   string
	Date     string
	IsActive bool `sql:",notnull"`
	IsOnline bool
}

func (b blackList) GetDate() string {
	t, e := time.Parse("2006-01-02", b.Date)
	if e != nil {
		return ""
	}
	return t.Format("02.01.2006")
}

type blacklistPageData struct {
	Header
	IsActive bool
	List     []blackList
}

func searchInOnline(name string, list []Online) bool {
	for _, o := range list {
		if o.Name == name {
			return true
		}
	}
	return false
}

var BlacklistPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var IsActive = !(r.FormValue("close") == "on")
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	var bList []Blacklist
	errS := db.Model(&bList).Where("is_active = ?", IsActive).Select()
	if errS != nil {
		log.Println("Error getting servers ", errS)
	}

	var sO []Server
	var sOI []int64
	errSOS := db.Model(&sO).Select()
	if errSOS != nil {
		log.Println("Error getting servers ", errS)
	}

	for _, s := range sO {
		sOI = append(sOI, s.Id)
	}

	var online []Online
	errSO := db.Model(&online).WhereIn("server_id = ?", sOI).Select()
	if errSO != nil {
		log.Println("Error getting servers online ", errSO)
	}

	var list []blackList

	for _, l := range bList {
		list = append(list, blackList{
			Id:       l.Id,
			Name:     l.Name,
			Reason:   l.Reason,
			Date:     l.Date,
			IsActive: l.IsActive,
			IsOnline: searchInOnline(l.Name, online),
		})
	}

	data := blacklistPageData{
		List:     list,
		IsActive: IsActive,
		Header: Header{
			Page: "blacklist",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "blacklist.tpl", data)
})

var NewBlacklist = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	blacklist := Blacklist{
		Name:     r.FormValue("name"),
		Date:     r.FormValue("date"),
		Reason:   r.FormValue("reason"),
		IsActive: r.FormValue("isActive") == "on",
	}

	err := db.Insert(&blacklist)

	if err != nil {
		http.Redirect(w, r, "/blacklist?error=Error create new blacklist", 303)
		return
	}

	http.Redirect(w, r, "/blacklist", 303)
})

var UpdateBlacklist = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "blacklist", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	var blacklist = Blacklist{Id: id}
	errS := db.Select(&blacklist)
	if errS != nil {
		log.Println("Error get blacklist", errS)
		http.Redirect(w, r, fmt.Sprintf("/blacklist?error=error select"), 303)
		return
	}
	name := r.FormValue("name")
	isActive := r.FormValue("isActive") == "on"
	date := r.FormValue("date")
	reason := r.FormValue("reason")
	if date == "" || reason == "" {
		http.Redirect(w, r, "/blacklist?error=Date error", 303)
		return
	}
	blacklist.Date = date
	blacklist.Reason = reason
	blacklist.Name = name
	blacklist.IsActive = isActive
	errU := db.Update(&blacklist)
	if errU != nil {
		log.Println("Error update blacklist: ", errU)
		http.Redirect(w, r, "/blacklist/?error=Error update data", 303)
		return
	}
	http.Redirect(w, r, "/blacklist", 303)
})

var DeleteBlacklist = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "blacklist", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	blacklist := Blacklist{
		Id: id,
	}
	errD := db.Select(&blacklist)
	if errD != nil {
		log.Println("Error get blacklist: ", errD)
		w.Write([]byte("ERROR"))
		return
	}
	blacklist.IsActive = false
	errU := db.Update(&blacklist)
	if errU != nil {
		log.Println("Error update blacklist: ", errU)
		w.Write([]byte("ERROR"))
		return
	}
	w.Write([]byte("OK"))
})
