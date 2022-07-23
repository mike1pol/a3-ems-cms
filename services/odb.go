package services

import (
	"log"
	"time"

	"github.com/go-pg/pg"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

func createOdb(db *pg.DB, name string, date time.Time) {
	var o = Odb{
		Name:  name,
		Date:  h.FormatISO(date),
		First: date,
		Last:  date,
	}
	err := db.Insert(&o)
	if err != nil {
		log.Println("error insert new ODB: ", err)
	}
}

func updateOdb(db *pg.DB, o Odb, t time.Time) {
	diff := t.Sub(o.Last).Minutes()
	o.Last = t
	if diff < 6 {
		o.Time = o.Time + diff
	}
	err := db.Update(&o)
	if err != nil {
		log.Println("error update ODB: ", err)
	}
}

func syncTime(db *pg.DB, name string, t time.Time) {
	var odb []Odb
	errS := db.Model(&odb).
		Where("name = ? and date = ?", name, h.FormatISO(t)).
		Select()
	if errS != nil {
		log.Println("Error select Online ", errS)
		return
	}

	if len(odb) == 0 {
		createOdb(db, name, t)
	} else {
		updateOdb(db, odb[0], t)
	}
}

func syncEMSDuty(db *pg.DB, servers []int64) {
	log.Println("syncEMSDuty")
	var personal = GetPersonal(db, PersonalFilter{Name: "", Rank: "", Status: "active"})
	var online []Online
	errS := db.Model(&online).
		Where("server_id IN (?)", pg.In(servers)).
		Select()
	if errS != nil {
		log.Println("Error getting online data", errS)
		return
	}
	var onlineEMS = h.OnlineEMS(online, personal)

	var duty []Duty
	errSD := db.Model(&duty).Select()
	if errSD != nil {
		log.Println("Error getting online data", errS)
		return
	}

	var aId []int64
	var a int = 0;
	for _, e := range onlineEMS {
		index := h.InDutyByPersonalId(e.Personal.Id, duty)
		if index == -1 {
			d := Duty{
				PersonalId: e.Personal.Id,
				Zone:       0,
			}
			db.Insert(&d)
			aId = append(aId, d.Id)
			a = a + 1;
		} else {
			aId = append(aId, duty[index].Id)
		}
	}
	var r int = 0
	for _, d := range duty {
		if !h.InInt64(d.Id, aId) {
			db.Delete(&Duty{Id: d.Id})
			r = r + 1;
		}
	}
	log.Printf("added: %d / removed: %d\n", a, r)
	log.Println("end syncEMSDuty")
}
