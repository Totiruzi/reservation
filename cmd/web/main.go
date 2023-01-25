package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Totiruzi/reservation/pkg/config"
	"github.com/Totiruzi/reservation/pkg/handlers"
	"github.com/Totiruzi/reservation/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCatche()
	if err != nil {
		log.Fatal("could not create template cache")
	}
	app.TamplateCatch = tc
	app.UseCatche = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("server is starting on port", port)

	srv := http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
