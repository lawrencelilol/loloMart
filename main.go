package main

import (
	"log"
	"loloMart/controllers"
	"loloMart/models"
	"loloMart/templates"
	"loloMart/views"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(
			templates.FS,
			"home.html", "tailwind.html",
		))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(
			templates.FS,
			"contact.html", "tailwind.html",
		))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(
			templates.FS,
			"faq.html", "tailwind.html",
		))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userService := models.UserService{
		DB: db,
	}
	usersC := controllers.Users{
		UserService: &userService, //TODO: implement UserService
	}

	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.html", "tailwind.html",
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.html", "tailwind.html",
	))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	log.Print("Listening on :4000...")

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),

		//TODO:Fix this before deployment
		csrf.Secure(false),
	)

	err = http.ListenAndServe(":4000", csrfMw(r))

	if err != nil {
		log.Fatal(err)
	}
}
