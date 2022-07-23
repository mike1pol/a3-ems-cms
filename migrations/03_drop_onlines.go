package main

import (
  "fmt"

  "github.com/go-pg/migrations"
)

func init() {
  migrations.Register(func(db migrations.DB) error {
    fmt.Println("remove onlines")
    _, err := db.Exec(`drop table if exists online`)
    return err
  }, func(db migrations.DB) error {
    fmt.Println("remove onlines")
    _, err := db.Exec(`drop table if exists online`)
    return err
  })
}
