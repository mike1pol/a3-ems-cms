package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
	s "github.com/mike1pol/rms/services"
)

type ServersPageData struct {
	Header
	List []Server
}

var ServersPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	var servers []Server
	errS := db.Model(&servers).Select()
	if errS != nil {
		log.Println("Error getting servers ", errS)
	}

	data := ServersPageData{
		List: servers,
		Header: Header{
			Page: "admin",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "servers.tpl", data)
})

var NewServer = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()

	port, e := strconv.ParseInt(r.FormValue("Port"), 10, 64)

	if e != nil {
		http.Redirect(w, r, "/server?error=Error create new server", 303)
		return
	}

	server := Server{
		Name: r.FormValue("Name"),
		Ip:   r.FormValue("Ip"),
		Port: port,
	}

	err := db.Insert(&server)

	if err != nil {
		http.Redirect(w, r, "/server?error=Error create new server", 303)
		return
	}

	errU := s.ServerUpdate(server)
	if errU != nil {
		http.Redirect(w, r, "/server?error=Error update server", 303)
		return
	}

	http.Redirect(w, r, "/server", 303)
})

var ActionServer = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	action := r.FormValue("type")
	server := Server{
		Id: id,
	}
	err := db.Select(&server)
	if err != nil {
		log.Println("Error getting server for action ", err)
		http.Redirect(w, r, "/server?error=Error getting server", 303)
		return
	}
	log.Printf("action: %s", action)
	if action == "delete" {
		err := db.Delete(&server)
		if err != nil {
			http.Redirect(w, r, "/server?error=Error delete server", 303)
			return
		}
	} else if action == "refresh" {
		err := s.ServerUpdate(server)
		if err != nil {
			http.Redirect(w, r, "/server?error=Error update server", 303)
			return
		}
	}
	http.Redirect(w, r, "/server", 303)
})

var ServersRefreshPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()

	s.ServersUpdate()

	http.Redirect(w, r, "/server", 303)
})
