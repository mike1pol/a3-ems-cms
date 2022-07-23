package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

// Rebuke
var NewRebuke = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	date := r.FormValue("date")
	reason := r.FormValue("reason")
	description := r.FormValue("description")
	if date == "" || reason == "" {
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Date error", id), 303)
		return
	}
	db := h.GetDB()
	defer db.Close()
	ins := PersonalRebuke{
		PersonalId:  id,
		Date:        date,
		Reason:      reason,
		Description: description,
	}
	errI := db.Insert(&ins)
	if errI != nil {
		log.Println("Error insert person rebuke: ", errI)
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Error insert data", id), 303)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/personal/%d", id), 303)
})

var UpdateRebuke = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	personId, e := strconv.ParseInt(vars["personId"], 10, 64)
	if e != nil || personId == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	var rebuke = PersonalRebuke{Id: id, PersonalId: personId}
	errS := db.Select(&rebuke)
	if errS != nil {
		log.Println("Error get vacation", errS)
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Date error", id), 303)
		return
	}
	date := r.FormValue("date")
	reason := r.FormValue("reason")
	description := r.FormValue("description")
	if date == "" || reason == "" {
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Date error", id), 303)
		return
	}
	rebuke.Date = date
	rebuke.Reason = reason
	rebuke.Description = description
	errU := db.Update(&rebuke)
	if errU != nil {
		log.Println("Error insert rank person: ", errU)
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Error update data", id), 303)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/personal/%d", personId), 303)
})

var DeleteRebuke = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	personId, e := strconv.ParseInt(vars["personId"], 10, 64)
	if e != nil || personId == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	vac := PersonalRebuke{
		Id:         id,
		PersonalId: personId,
	}
	errD := db.Delete(&vac)
	if errD != nil {
		log.Println("Error get vacation: ", errD)
		w.Write([]byte("ERROR"))
		return
	}
	w.Write([]byte("OK"))
})

// Vacation
var NewVacation = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	start := r.FormValue("Start")
	end := r.FormValue("End")
	if start == "" || end == "" {
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Date error", id), 303)
		return
	}
	db := h.GetDB()
	defer db.Close()
	ins := PersonalVacation{
		PersonalId: id,
		Start:      start,
		End:        end,
	}
	errI := db.Insert(&ins)
	if errI != nil {
		log.Println("Error insert rank person: ", errI)
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Error insert data", id), 303)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/personal/%d", id), 303)
})

var UpdateVacation = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	personId, e := strconv.ParseInt(vars["personId"], 10, 64)
	if e != nil || personId == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	var vac = PersonalVacation{Id: id, PersonalId: personId}
	errS := db.Select(&vac)
	if errS != nil {
		log.Println("Error get vacation", errS)
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Date error", id), 303)
		return
	}
	start := r.FormValue("Start")
	end := r.FormValue("End")
	if start == "" || end == "" {
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Date error", id), 303)
		return
	}
	vac.Start = start
	vac.End = end
	errU := db.Update(&vac)
	if errU != nil {
		log.Println("Error insert rank person: ", errU)
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Error update data", id), 303)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/personal/%d", personId), 303)
})

var DeleteVacation = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	personId, e := strconv.ParseInt(vars["personId"], 10, 64)
	if e != nil || personId == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	vac := PersonalVacation{
		Id:         id,
		PersonalId: personId,
	}
	errD := db.Delete(&vac)
	if errD != nil {
		log.Println("Error get vac: ", errD)
		w.Write([]byte("ERROR"))
		return
	}
	w.Write([]byte("OK"))
})

// Person
var EditPerson = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	err := editPerson(db, id, r)
	if err != nil {
		if r.FormValue("plain") == "" {
			http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Error create new person", id), 303)
		} else {
			w.Write([]byte("ERROR"))
		}
		return
	}
	if r.FormValue("plain") == "" {
		http.Redirect(w, r, fmt.Sprintf("/personal/%d", id), 303)
	} else {
		w.Write([]byte("OK"))
	}
})

var NewPerson = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	err := createPerson(db, r)
	if err != nil {
		http.Redirect(w, r, "/personal?error=Error create new person", 303)
		return
	}
	http.Redirect(w, r, "/personal", 303)
})

type PersonalViewData struct {
	Header
	Person PersonalJson
	Ranks  []Rank
}

var ViewPerson = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	var p = GetPersonalById(db, id)
	if p.Name == "" {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	p.Odbs = p.GetReverseOdbs()
	var ranks []Rank
	errSr := db.Model(&ranks).Select()
	if errSr != nil {
		log.Println("Error getting ranks ", errSr)
	}

	ranks = h.RankSort(ranks, false)
	data := PersonalViewData{
		Header: Header{
			Title: p.Name,
			Page:  "personal",
			User:  currentUser,
		},
		Person: p,
		Ranks:  ranks,
	}
	h.RenderTemplate(w, "person.tpl", data)
})

type PersonalPageData struct {
	Header
	List   []PersonalList
	Count  int
	Ranks  []Rank
	RanksR []Rank
	Filter PersonalFilter
}

var ViewPersons = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	filter := PersonalFilter{
		Rank:   "",
		Name:   "",
		Status: "active",
	}

	if r.FormValue("r") != "" {
		filter.Rank = r.FormValue("r")
	}

	if r.FormValue("n") != "" {
		filter.Name = r.FormValue("n")
	}
	if r.FormValue("s") != "" {
		filter.Status = r.FormValue("s")
	}

	list := h.RankDateSortByRank(GetPersonal(db, filter))

	var ranks []Rank
	errSr := db.Model(&ranks).Select()
	if errSr != nil {
		log.Println("Error getting ranks ", errSr)
	}

	ranks = h.RankSort(ranks, false)

	var rr = h.RankSort(ranks, true)

	data := PersonalPageData{
		List:   list,
		Count:  len(list),
		Ranks:  ranks,
		RanksR: rr,
		Filter: filter,
		Header: Header{
			Title: "Личный состав",
			Page:  "personal",
			User:  currentUser,
		},
	}
	h.RenderTemplate(w, "personal.tpl", data)
})

func editPerson(db *pg.DB, id int64, r *http.Request) error {
	name := r.FormValue("Name")
	status := r.FormValue("Status")
	dismissalDate := r.FormValue("DismissalDate")
	steamID := r.FormValue("steamId")
	dm := Personal{
		Id: id,
	}
	err := db.Select(&dm)
	if err != nil {
		log.Println("Error getting personal for update ", err)
		return err
	}
	if len(name) > 0 {
		dm.Name = name
	}
	if len(steamID) > 0 {
		dm.SteamId = steamID
	}
	if len(status) > 0 {
		dm.Status = status
	}
	if len(dismissalDate) > 0 {
		dm.DismissalDate = dismissalDate
	}
	errU := db.Update(&dm)
	if errU != nil {
		log.Println("Error update person  ", errU)
		return errU
	}
	return nil
}

func createPerson(db *pg.DB, r *http.Request) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Error start transaction")
		return err
	}
	defer tx.Rollback()
	rnk, e := strconv.ParseInt(r.FormValue("Rank"), 10, 64)
	if e != nil {
		log.Println("Error parse rank ", e)
		return e
	}

	person := Personal{
		Name:    r.FormValue("Name"),
		SteamId: r.FormValue("steamId"),
		Status:  "active",
	}

	errIP := tx.Insert(&person)
	if errIP != nil {
		log.Println("Error insert person: ", errIP)
		return errIP
	}

	rank := PersonalRank{
		PersonalId: person.Id,
		RankId:     rnk,
		Date:       r.FormValue("Date"),
	}
	errIr := tx.Insert(&rank)
	if errIr != nil {
		log.Println("Error insert person: ", errIr)
		return errIr
	}
	tx.Commit()
	return nil
}

// Person rank

var NewPersonRank = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	rnk, _ := strconv.ParseInt(r.FormValue("Rank"), 10, 64)
	date := r.FormValue("Date")
	db := h.GetDB()
	defer db.Close()
	rank := PersonalRank{
		PersonalId: id,
		RankId:     rnk,
		Date:       date,
	}
	errIr := db.Insert(&rank)
	if errIr != nil {
		log.Println("Error insert rank person: ", errIr)
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Error insert data", id), 303)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/personal/%d", id), 303)
})

var DeletePersonRank = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	personId, e := strconv.ParseInt(vars["personId"], 10, 64)
	if e != nil || personId == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	db := h.GetDB()
	defer db.Close()
	rank := PersonalRank{
		Id:         id,
		PersonalId: personId,
	}
	errD := db.Delete(&rank)
	if errD != nil {
		log.Println("Error get rank: ", errD)
		w.Write([]byte("ERROR"))
		return
	}
	w.Write([]byte("OK"))
})

var ChangePersonRank = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	vars := mux.Vars(r)
	personId, e := strconv.ParseInt(vars["personId"], 10, 64)
	if e != nil || personId == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil || id == 0 {
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	rnk, _ := strconv.ParseInt(r.FormValue("Rank"), 10, 64)
	date := r.FormValue("Date")
	db := h.GetDB()
	defer db.Close()
	rank := PersonalRank{
		Id:         id,
		PersonalId: personId,
	}
	errS := db.Select(&rank)
	if errS != nil {
		log.Println("Error get rank: ", errS)
		h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "personal", Header: Header{User: currentUser}})
		return
	}
	rank.RankId = rnk
	rank.Date = date
	errU := db.Update(&rank)
	if errU != nil {
		log.Println("Error update rank person: ", errU)
		http.Redirect(w, r, fmt.Sprintf("/personal/%d?error=Error update data", personId), 303)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/personal/%d", personId), 303)
})
