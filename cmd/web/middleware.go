package main

import (
	"net/http"

	"github.com/bertoxic/bert/helpers"
	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("we are using custom middlewares")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}

func SessionsLoad(next http.Handler) http.Handler {
	return sessions.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r){
			sessions.Put(r.Context(),"error","Please Login first")
			http.Redirect(w, r, "/user-login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w,r)
	})
}