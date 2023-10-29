package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mauFade/bookings/pkg/config"
	"github.com/mauFade/bookings/pkg/handler"
	"github.com/mauFade/bookings/pkg/render"
)

const PORT = ":8081"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Application running on port", PORT)

	serve := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()

	log.Fatal(err)
}
