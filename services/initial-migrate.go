package services

import (
	"log"

	"github.com/go-pg/pg"

	. "github.com/mike1pol/rms/models"
)

func InitialMigrate(db *pg.DB) {
	var ranks []Rank
	count, err := db.Model(&ranks).Count()
	if err != nil {
		log.Println("Error getting count ranks: ", err)
		return
	}
	if count == 0 {
		_, err := db.Exec("INSERT INTO ranks (name, next, sort) VALUES ('Интерн', '7', '7'), ('Техник', '14', '6'), ('Специалист', '24', '5'), ('Ст. специалист', '35', '4'), ('Майор', null, '3'), ('Зам. министра', null, '2'), ('Министр', null, '1')")
		if err != nil {
			log.Println("error initial migration data: ", err)
		}
	}

}
