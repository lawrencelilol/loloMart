package main

import (
	"log"
	"loloMart/controllers"
	"loloMart/models"
	"loloMart/templates"
	"loloMart/views"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})

	log.Print("Listening on :3000...")

	err = http.ListenAndServe(":3000", r)

	if err != nil {
		log.Fatal(err)
	}
}
