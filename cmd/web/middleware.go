package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// PrintToConsole just prints a string to the console
func PrintToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serve the next page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRFT to every POST requests
func NoSurf(next http.Handler) http.Handler {
	csrftHandler := nosurf.New(next)
	csrftHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrftHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad (next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}