package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("alert personals add status...")
		_, err := db.Exec(`alter table personals add column status text`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("remove status from personals...")
		_, err := db.Exec(`alter table personals drop column status restrict`)
		return err
	})
}
