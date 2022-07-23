package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"

	a "github.com/mike1pol/rms/api"
	c "github.com/mike1pol/rms/controllers"
	h "github.com/mike1pol/rms/helpers"
	s "github.com/mike1pol/rms/services"
)

func main() {
	r := mux.NewRouter()
	var db = h.GetDB()
	defer db.Close()
	errSchema := h.CreateSchema(db)
	if errSchema != nil {
		log.Panic("Error shema creating: ", errSchema)
	}

	s.InitialMigrate(db)
	go func() {
		s.ServersUpdate()
	}()

	cr := cron.New()
	cr.AddFunc("@every 5m", func() {
		ts := time.Now()
		log.Printf("cronjob start %s\n", ts)
		s.ServersUpdate()
		log.Printf("cronjob end %s\n work time: %s", time.Now(), time.Now().Sub(ts))
	})
	cr.Start()

	h.LoadConfiguration()
	h.LoadTemplates()

	// API
	r.Handle("/api/v1/online", a.OnlineAPIGet).Methods("GET")
	r.Handle("/api/v1/odb/search", a.OdbSearchAPIGet).Methods("GET")

	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	r.Handle("/", c.HomePage).Methods("GET")
	// ODB
	r.Handle("/odb", h.PersonalRequiredMiddleware(c.OdbPage)).Methods("GET")
	// Raport
	r.Handle("/raport", h.PersonalRequiredMiddleware(c.RaportCreatePage)).Methods("GET")
	r.Handle("/raport", h.PersonalRequiredMiddleware(c.RaportCreateHandler)).Methods("POST")
	r.Handle("/raport/list", h.AdminRequiredMiddleware(c.RaportListPage)).Methods("GET")
	r.Handle("/raport/{id}", h.AdminRequiredMiddleware(c.RaportViewPage)).Methods("GET")
	r.Handle("/raport/{id}/status", h.AdminRequiredMiddleware(c.RaportStatusUpdateHandler)).Methods("GET")
	r.Handle("/raport/{id}/comment", h.AdminRequiredMiddleware(c.RaportCommentHandler)).Methods("POST")
	// Blacklist
	r.Handle("/blacklist", c.BlacklistPage).Methods("GET")
	r.Handle("/blacklist", h.AdminRequiredMiddleware(c.NewBlacklist)).Methods("POST")
	r.Handle("/blacklist/{id}", h.AdminRequiredMiddleware(c.UpdateBlacklist)).Methods("POST")
	r.Handle("/blacklist/{id}", h.AdminRequiredMiddleware(c.DeleteBlacklist)).Methods("DELETE")
	// RefDB
	r.Handle("/refdb", c.RefDBPage).Methods("GET")
	r.Handle("/refdb", h.PersonalRequiredMiddleware(c.NewRefDB)).Methods("POST")
	r.Handle("/refdb/{id}", h.PersonalRequiredMiddleware(c.UpdateRefDB)).Methods("POST")
	r.Handle("/refdb/{id}", h.AdminRequiredMiddleware(c.DeleteRefDB)).Methods("DELETE")
	// Psy
	r.Handle("/psy", c.PsyPage).Methods("GET")
	r.Handle("/psy", h.PersonalRequiredMiddleware(c.NewPsy)).Methods("POST")
	r.Handle("/psy/{id}", h.PersonalRequiredMiddleware(c.UpdatePsy)).Methods("POST")
	r.Handle("/psy/{id}", h.AdminRequiredMiddleware(c.DeletePsy)).Methods("DELETE")
	// Duty
	r.Handle("/duty", c.DutyPage).Methods("GET")
	r.Handle("/duty/action", h.PersonalRequiredMiddleware(c.ActionDuty)).Methods("GET")
	// Personal
	r.Handle("/personal", h.PersonalRequiredMiddleware(c.ViewPersons)).Methods("GET")
	r.Handle("/personal/", h.PersonalRequiredMiddleware(c.ViewPersons)).Methods("GET")
	r.Handle("/personal", h.AdminRequiredMiddleware(c.NewPerson)).Methods("POST")
	r.Handle("/personal/{id}", h.PersonalRequiredMiddleware(c.ViewPerson)).Methods("GET")
	r.Handle("/personal/{id}", h.AdminRequiredMiddleware(c.EditPerson)).Methods("POST")
	// Personal rank
	r.Handle("/personal/{id}/rank", h.AdminRequiredMiddleware(c.NewPersonRank)).Methods("POST")
	r.Handle("/personal/{personId}/rank/{id}", h.AdminRequiredMiddleware(c.ChangePersonRank)).Methods("POST")
	r.Handle("/personal/{personId}/rank/{id}", h.AdminRequiredMiddleware(c.DeletePersonRank)).Methods("DELETE")
	// Personal vacation
	r.Handle("/personal/{id}/vacation", h.AdminRequiredMiddleware(c.NewVacation)).Methods("POST")
	r.Handle("/personal/{personId}/vacation/{id}", h.AdminRequiredMiddleware(c.UpdateVacation)).Methods("POST")
	r.Handle("/personal/{personId}/vacation/{id}", h.AdminRequiredMiddleware(c.DeleteVacation)).Methods("DELETE")
	// Personal rebuke
	r.Handle("/personal/{id}/rebuke", h.AdminRequiredMiddleware(c.NewRebuke)).Methods("POST")
	r.Handle("/personal/{personId}/rebuke/{id}", h.AdminRequiredMiddleware(c.UpdateRebuke)).Methods("POST")
	r.Handle("/personal/{personId}/rebuke/{id}", h.AdminRequiredMiddleware(c.DeleteRebuke)).Methods("DELETE")
	// Personal report
	r.Handle("/report", h.AdminRequiredMiddleware(c.ViewReport)).Methods("GET")
	r.Handle("/report/list", h.AdminRequiredMiddleware(c.ViewReport)).Methods("GET")
	// Rank
	r.Handle("/rank", h.AdminRequiredMiddleware(c.ViewRanks)).Methods("GET")
	r.Handle("/rank", h.AdminRequiredMiddleware(c.NewRank)).Methods("POST")
	r.Handle("/rank/{id}", h.AdminRequiredMiddleware(c.EditRank)).Methods("POST")
	r.Handle("/rank/{id}", h.AdminRequiredMiddleware(c.DeleteRank)).Methods("DELETE")
	// Users
	r.Handle("/user", h.AdminRequiredMiddleware(c.UsersPage)).Methods("GET")
	r.Handle("/user/{id}/action", h.AdminRequiredMiddleware(c.ActionUser)).Methods("GET")
	r.Handle("/login", c.LoginPage).Methods("GET", "POST")
	r.Handle("/logout", c.LogoutPage).Methods("GET")
	// Server
	r.Handle("/server", h.AdminRequiredMiddleware(c.ServersPage)).Methods("GET")
	r.Handle("/server/{id}/action", h.AdminRequiredMiddleware(c.ActionServer)).Methods("GET")
	r.Handle("/server/refresh", h.AdminRequiredMiddleware(c.ServersRefreshPage)).Methods("GET")
	r.Handle("/server", h.AdminRequiredMiddleware(c.NewServer)).Methods("POST")
	r.NotFoundHandler = h.NotFound
	r.Use(h.UserMiddleware)

	log.Println("server starting http://0.0.0.0:8080")
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))

}
