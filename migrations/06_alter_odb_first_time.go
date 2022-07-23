package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("alert odbs add column first timestamp...")
		_, err := db.Exec(`alter table odbs add column first timestamp`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("alert odbs drop column first...")
		_, err := db.Exec(`alter table odbs drop column first restrict`)
		return err
	})
}
