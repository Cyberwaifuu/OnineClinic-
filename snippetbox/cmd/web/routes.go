package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// Add the requireAuthentication middleware to the chain.
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	// Add the requireAuthentication middleware to the chain.
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet))

	mux.Get("/appointment/createAppointment", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createAppointmentForm))
	mux.Post("/appointment/createAppointment", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createAppointment))
	mux.Get("/appointment/:id", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.showAppointment))
	mux.Get("/appointments", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.showAppointments))

	mux.Get("/services", dynamicMiddleware.ThenFunc(app.showServices))

	mux.Get("/category/general", dynamicMiddleware.ThenFunc(app.category))
	mux.Get("/category/international", dynamicMiddleware.ThenFunc(app.category))
	mux.Get("/category/events", dynamicMiddleware.ThenFunc(app.category))
	mux.Get("/category/competition", dynamicMiddleware.ThenFunc(app.category))
	mux.Get("/category/student", dynamicMiddleware.ThenFunc(app.category))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)

}
