package helpers

import (
	"context"
	"net/http"

	. "github.com/mike1pol/rms/models"
)

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, err := r.Cookie("auth")
		if err == nil && len(auth.Value) == 32 {
			var db = GetDB()
			defer db.Close()
			var us UserSession
			err := db.Model(&us).Relation("User").Limit(1).Where("token = ?", auth.Value).Select()
			if err == nil {
				ctx := context.WithValue(r.Context(), "user", us.User)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func PersonalRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, err := r.Cookie("auth")
		if err == nil && len(auth.Value) == 32 {
			var db = GetDB()
			defer db.Close()
			var us UserSession
			err := db.Model(&us).Relation("User").Limit(1).Where("token = ?", auth.Value).Select()
			if err == nil {
				if us.User.IsAdmin || us.User.IsUser {
					ctx := context.WithValue(r.Context(), "user", us.User)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					RenderTemplate(w, "forbidden.tpl", PageNotFound{Page: "forbidden"})
				}
				return
			}
		}
		RenderTemplate(w, "forbidden.tpl", PageNotFound{Page: "forbidden"})
	})
}

func AdminRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, err := r.Cookie("auth")
		if err == nil && len(auth.Value) == 32 {
			var db = GetDB()
			defer db.Close()
			var us UserSession
			err := db.Model(&us).Relation("User").Limit(1).Where("token = ?", auth.Value).Select()
			if err == nil {
				if us.User.IsAdmin {
					ctx := context.WithValue(r.Context(), "user", us.User)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					RenderTemplate(w, "forbidden.tpl", PageNotFound{Page: "forbidden"})
				}
				return
			}
		}
		RenderTemplate(w, "forbidden.tpl", PageNotFound{Page: "forbidden"})
	})
}

var NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var currentUser User
	if r.Context().Value("user") != nil {
		currentUser = r.Context().Value("user").(User)
	}
	RenderTemplate(w, "not-found.tpl", PageNotFound{Page: "not-found", Header: Header{User: currentUser}})
})
