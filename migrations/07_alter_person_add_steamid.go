package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("alert personals add column steamId text...")
		_, err := db.Exec(`alter table personals add column steam_id text`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("alert odbs drop column first...")
		_, err := db.Exec(`alter table personals drop column steam_id restrict`)
		return err
	})
}
