package main

import (
  "flag"
  "fmt"
  "log"
  "os"

  "github.com/go-pg/migrations"
  "github.com/go-pg/pg"
  "github.com/gobuffalo/envy"
)

const usageText = `This program runs command on the db. Supported commands are:
  - up - runs all available migrations.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func check(e error) {
  if e != nil {
    log.Fatal(e)
  }
}

func main() {
  flag.Usage = usage
  flag.Parse()

  pgHost, err := envy.MustGet("PG_HOST")
  check(err)
  pgPort, err := envy.MustGet("PG_PORT")
  check(err)
  pgUser, err := envy.MustGet("PG_USER")
  check(err)
  pgPass, err := envy.MustGet("PG_PASS")
  check(err)
  pgDb, err := envy.MustGet("PG_DB")
  check(err)
  dbConfig := &pg.Options{
    Addr:     fmt.Sprintf("%s:%s", pgHost, pgPort),
    User:     pgUser,
    Password: pgPass,
    Database: pgDb,
  }

  db := pg.Connect(dbConfig)

  oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
  if err != nil {
    exitF(err.Error())
  }
  if newVersion != oldVersion {
    fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
  } else {
    fmt.Printf("version is %d\n", oldVersion)
  }
}

func usage() {
  fmt.Printf(usageText)
  flag.PrintDefaults()
  os.Exit(2)
}

func errorF(s string, args ...interface{}) {
  fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitF(s string, args ...interface{}) {
  errorF(s, args...)
  os.Exit(1)
}
