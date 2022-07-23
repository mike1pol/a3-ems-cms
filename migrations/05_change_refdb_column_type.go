package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("alert ref_dbs change column date type to date...")
		_, err := db.Exec(`alter table ref_dbs alter column date type date using to_date(date, 'YYYY-MM-DD');`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("alert ref_dbs change column date type to text...")
		_, err := db.Exec(`alter table ref_dbs alter column date type TEXT;`)
		return err
	})
}
