package controllers

import (
	"log"
	"net/http"
	"time"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

type ReportPageData struct {
	Header
	Ch     int
	List   []PersonalJson
	Days   []string
	Ranks  []Rank
	Filter PersonalFilter
}

var ViewReport = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db := h.GetDB()
	defer db.Close()
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}

	startD, _ := time.Parse("2006-01-02", time.Now().Add(-time.Hour * 24 * 7).Format("2006-01-02"))
	endD, _ := time.Parse("2006-01-02", time.Now().Add(time.Hour * 3).Format("2006-01-02"))

	filter := PersonalFilter{
		Name:   "",
		Rank:   "",
		Status: "active",
		Start:  startD,
		End:    endD,
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
	if r.FormValue("start") != "" {
		start, err := time.Parse("2006-01-02", r.FormValue("start"))
		if err == nil {
			filter.Start = start
		}
	}
	if r.FormValue("end") != "" {
		end, err := time.Parse("2006-01-02", r.FormValue("end"))
		if err == nil {
			filter.End = end
		}
	}

	diff := int(filter.End.Sub(filter.Start).Hours() / 24)
	var dlist []string
	for i := 0; i < diff+1; i = i + 1 {
		if i == 0 {
			dlist = append(dlist, filter.Start.Format("02.01"))
		} else {
			dlist = append(dlist, filter.Start.Add(time.Hour * time.Duration(24*i)).Format("02.01"))
		}
	}
	var list []PersonalJson
	l := h.RankDateSortByRank(GetPersonal(db, filter))

	for i := range l {
		for c := range l[i].List {
			list = append(list, l[i].List[c])
		}
	}

	var ranks []Rank
	errSr := db.Model(&ranks).Select()
	if errSr != nil {
		log.Println("Error getting ranks ", errSr)
	}

	data := ReportPageData{
		List:   list,
		Days:   dlist,
		Filter: filter,
		Ch:     2 + len(dlist),
		Ranks:  h.RankSort(ranks, true),
		Header: Header{
			Title: "Отчеты",
			Page:  "report",
			User:  currentUser,
		},
	}
	h.RenderTemplate(w, "report.tpl", data)
})
