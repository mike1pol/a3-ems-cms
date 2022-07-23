package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gobuffalo/envy"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

func getConfig() (clientSecret string, redirectUrl string) {
	cSecret, err := envy.MustGet("CLIENT_SECRET")
	if err != nil {
		log.Fatal(err)
	}
	clientSecret = cSecret
	rUrl, err := envy.MustGet("REDIRECT_URL")
	if err != nil {
		log.Fatal(err)
	}
	redirectUrl = rUrl
	return
}

var LoginPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	secret, redirectUrl := getConfig()
	opId := h.NewOpenId(r, redirectUrl)
	switch opId.Mode() {
	case "":
		http.Redirect(w, r, opId.AuthUrl(), 303)
	case "cancel":
		w.Write([]byte("Authorization cancelled"))
	default:
		steamId, err := opId.ValidateAndGetId()
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/?error=Login Error", 303)
			return
		}
		p, e := GetPlayerSummaries(steamId, secret)
		if e != nil {
			log.Println(e)
			http.Redirect(w, r, "/?error=Login Error", 303)
			return
		}
		db := h.GetDB()
		var users []User
		err = db.Model(&users).Where("steam_id = ?", steamId).Select()
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/?error=Login Error", 303)
			return
		}
		var u User
		var us UserSession
		if len(users) == 0 {
			log.Println("new user")
			u = User{
				Name:       p.PersonaName,
				SteamId:    p.SteamId,
				ProfileUrl: p.ProfileUrl,
				Avatar:     p.Avatar,
				IsAdmin:    false,
			}
			err := db.Insert(&u)
			if err != nil {
				http.Redirect(w, r, "/?error=Login Error", 303)
				return
			}
		} else {
			log.Println("already registred")
			u = users[0]
		}
		us = UserSession{
			UserId: u.Id,
			Token:  h.RandString(32),
		}
		errUs := db.Insert(&us)
		if errUs != nil {
			http.Redirect(w, r, "/?error=Login Error", 303)
			return
		}
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "auth", Value: us.Token, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 303)
	}
})

var LogoutPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	auth, err := r.Cookie("auth")
	if err == nil && len(auth.Value) == 32 {
		var db = h.GetDB()
		defer db.Close()
		_, err := db.Exec(`delete from user_sessions where token = ?`, auth.Value)
		if err != nil {
			log.Println(err)
		}
	}

	expiration := time.Now()
	cookie := http.Cookie{Name: "auth", Value: "delete", Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", 303)
})
