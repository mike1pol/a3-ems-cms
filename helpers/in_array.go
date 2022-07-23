package helpers

import (
  "time"

  . "github.com/mike1pol/rms/models"
)

func InInt64(id int64, dt []int64) bool {
  for i := range dt {
    if dt[i] == id {
      return true
    }
  }
  return false
}

func InDutyByPersonalId(id int64, dt []Duty) int {
  for i := range dt {
    if dt[i].PersonalId == id {
      return i
    }
  }
  return -1
}

func MedsInServer(s []OnlineMed) (c int) {
  for i := range s {
    if s[i].Personal.Id != 0 {
      c = c + 1
    }
  }
  return
}

func OnlineEMS(array []Online, personal []PersonalJson) (s []OnlineMed) {
  for i := range array {
    p := InPersonal(array[i].Name, personal)
    v := p.GetVacForDate(time.Now())
    o := OnlineMed{
      Online: array[i],
    }
    if p.Id > 0 && v.Id == 0 {
      o.Personal = InPersonal(array[i].Name, personal)
      s = append(s, o)
    }
  }
  return
}

func InServerN(server Server, array []Online, personal []PersonalJson) (s ServerOnline) {
  s = ServerOnline{
    Id:         server.Id,
    Name:       server.Name,
    Status:     server.Status,
    Online:     server.Online,
    OnlineMed:  0,
    MaxPlayers: server.MaxPlayers,
    LastUpdate: server.LastUpdate,
  }
  for i := range array {
    if array[i].ServerId == server.Id {
      p := InPersonal(array[i].Name, personal)
      v := p.GetVacForDate(time.Now())
      o := OnlineMed{
        Online: array[i],
      }
      if v.Id == 0 {
        o.Personal = InPersonal(array[i].Name, personal)
      }
      s.Players = append(s.Players, o)
    }
  }
  s.Players = OnlineSort(s.Players)
  s.OnlineMed = MedsInServer(s.Players)
  return
}

func InPersonal(val string, array []PersonalJson) (ok PersonalJson) {
  for i := range array {
    if array[i].Name == val {
      ok = array[i]
      return
    }
  }
  return
}

func InPersonalOdb(array []Odb, personal []PersonalJson, unGroup bool) (list []OdbMed) {
  for i := range array {
    om := inOdbMed(array[i], list)
    if om >= 0 && !unGroup {
      list[om].Odb.Time = list[om].Odb.Time + array[i].Time
    } else {
      o := OdbMed{
        Odb:      array[i],
        Personal: InPersonal(array[i].Name, personal),
      }
      list = append(list, o)
    }
  }
  return
}

func inOdbMed(o Odb, list []OdbMed) int {
  for i := range list {
    if list[i].Odb.Name == o.Name {
      return i
    }
  }
  return -1
}
