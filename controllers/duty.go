package controllers

import (
  "log"
  "net/http"
  "strconv"
  "fmt"

  h "github.com/mike1pol/rms/helpers"
  . "github.com/mike1pol/rms/models"
)

type DutyPageData struct {
  Header
  Z0  []Duty
  Z0C int
  Z1  []Duty
  Z1C int
  Z2  []Duty
  Z2C int
  Z3  []Duty
  Z3C int
  Z4  []Duty
  Z4C int
  Z5  []Duty
  Z5C int
}

var DutyPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  db := h.GetDB()
  defer db.Close()
  var currentUser User
  if r.Context().Value("user") != nil {
    currentUser = r.Context().Value("user").(User)
  }

  var duty []Duty
  errS := db.Model(&duty).
    Relation("Personal").
    Relation("Personal.Ranks").
    Relation("Personal.Ranks.Rank").
    Select()
  if errS != nil {
    log.Println("Errro getting duty: ", errS)
  }

  var z0 []Duty
  var z1 []Duty
  var z2 []Duty
  var z3 []Duty
  var z4 []Duty
  var z5 []Duty

  for _, d := range duty {
    if d.Zone == 0 {
      z0 = append(z0, d)
    } else if d.Zone == 1 {
      z1 = append(z1, d)
    } else if d.Zone == 2 {
      z2 = append(z2, d)
    } else if d.Zone == 3 {
      z3 = append(z3, d)
    } else if d.Zone == 4 {
      z4 = append(z4, d)
    } else if d.Zone == 5 {
      z5 = append(z5, d)
    }
  }
  data := DutyPageData{
    Z0:  h.DutySortByRank(z0),
    Z0C: len(z0),
    Z1:  z1,
    Z1C: len(z1),
    Z2:  z2,
    Z2C: len(z2),
    Z3:  z3,
    Z3C: len(z3),
    Z4:  z4,
    Z4C: len(z4),
    Z5:  z5,
    Z5C: len(z5),
    Header: Header{
      Page: "duty",
      User: currentUser,
    },
  }
  h.RenderTemplate(w, "duty.tpl", data)
})

var ActionDuty = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  var currentUser User
  if r.Context().Value("user") != nil {
    currentUser = r.Context().Value("user").(User)
  }
  id, e := strconv.ParseInt(r.FormValue("id"), 10, 64)
  val, ev := strconv.ParseInt(r.FormValue("to"), 10, 64)
  if e != nil || ev != nil || id == 0 {
    h.RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "duty", Header: Header{User: currentUser}})
    return
  }
  db := h.GetDB()
  defer db.Close()
  var duty = Duty{Id: id}
  errS := db.Select(&duty)
  if errS != nil {
    log.Println("Error get duty", errS)
    http.Redirect(w, r, fmt.Sprintf("/duty?error=error select"), 303)
    return
  }
  duty.Zone = int(val)
  errU := db.Update(&duty)
  if errU != nil {
    log.Println("Error update duty: ", errU)
    http.Redirect(w, r, "/duty?error=Error update data", 303)
    return
  }
  http.Redirect(w, r, "/duty", 303)
})
