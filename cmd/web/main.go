package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PiotrSochaczewski/GoBooking/pkg/config"
	"github.com/PiotrSochaczewski/GoBooking/pkg/handlers"
	"github.com/PiotrSochaczewski/GoBooking/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	//change this to true when in production
	app.InProduction = false

	//set up  the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	//method Cookie.Persist if "true" even if you close web session will not expire,  
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// Starting web server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
