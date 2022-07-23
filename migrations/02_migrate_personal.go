package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("migrate personals add default status...")
		_, err := db.Exec(`update personals set status='active'`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("remove status from personals...")
		return nil
	})
}
