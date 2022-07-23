package api

import (
  "net/http"
  "log"

  h "github.com/mike1pol/rms/helpers"
)

type OdbSearchAPIData struct {
  Status  string
  Message string
  Data    []string
}

var OdbSearchAPIGet = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  db := h.GetDB()
  defer db.Close()

  name := r.FormValue("name")

  var data []string
  _, err := db.Query(
    &data,
    "select distinct name from odbs where lower(name) like lower(?) order by name asc",
    "%"+name+"%",
  )
  if err != nil {
    log.Println("Error getting odbs ", err)
    h.SendJson(w, OdbSearchAPIData{Status: "error", Message: "Error get list"}, http.StatusInternalServerError)
    return
  }

  h.SendJson(w, OdbSearchAPIData{Status: "success", Data: data}, http.StatusOK)
})
