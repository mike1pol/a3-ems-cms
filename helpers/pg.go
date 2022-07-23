package helpers

import (
	"fmt"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/gobuffalo/envy"

	m "github.com/mike1pol/rms/models"
)

// CreateSchema create models to db
func CreateSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		&m.Server{},
		&m.Online{},
		&m.Personal{},
		&m.PersonalVacation{},
		&m.PersonalRebuke{},
		&m.Rank{},
		&m.PersonalRank{},
		&m.Odb{},
		&m.User{},
		&m.UserSession{},
		&m.Blacklist{},
		&m.RefDB{},
		&m.Psy{},
		&m.Duty{},
		&m.Raport{},
		&m.RaportComment{},
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// GetDB Get database
func GetDB() *pg.DB {
	pgHost, err := envy.MustGet("PG_HOST")
	if err != nil {
		log.Fatal(err)
	}
	pgPort, err := envy.MustGet("PG_PORT")
	if err != nil {
		log.Fatal(err)
	}
	pgUser, err := envy.MustGet("PG_USER")
	if err != nil {
		log.Fatal(err)
	}
	pgPass, err := envy.MustGet("PG_PASS")
	if err != nil {
		log.Fatal(err)
	}
	pgDb, err := envy.MustGet("PG_DB")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig := &pg.Options{
		Addr:     fmt.Sprintf("%s:%s", pgHost, pgPort),
		User:     pgUser,
		Password: pgPass,
		Database: pgDb,
	}
	db := pg.Connect(dbConfig)
	return db
}
