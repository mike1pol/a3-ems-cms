package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("alert personals add release_date...")
		_, err := db.Exec(`alter table personals add column dismissal_date date`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("remove status from personals...")
		_, err := db.Exec(`alter table personals drop column dismissal_date restrict`)
		return err
	})
}
