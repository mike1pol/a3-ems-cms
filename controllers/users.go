package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

type UsersPageData struct {
	Header
	List []User
}

var UsersPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	var users []User
	errS := db.Model(&users).OrderExpr("id ASC").Select()
	if errS != nil {
		log.Println("Error getting online data", errS)
	}

	data := UsersPageData{
		List: users,
		Header: Header{
			Page: "admin",
			User: currentUser,
		},
	}
	h.RenderTemplate(w, "users.tpl", data)
})

var ActionUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	action := r.FormValue("action")
	u := User{
		Id: id,
	}
	err := db.Select(&u)
	if err != nil {
		log.Println("Error getting personal for update ", err)
		http.Redirect(w, r, "/user?error=Error getting user info", 303)
		return
	}
	if action == "user" {
		u.IsUser = !u.IsUser
	}
	if action == "admin" {
		u.IsAdmin = !u.IsAdmin
	}
	errU := db.Update(&u)
	if errU != nil {
		log.Println("Error update person  ", errU)
		http.Redirect(w, r, "/user?error=User update error", 303)
		return
	}
	http.Redirect(w, r, "/user", 303)
})
