package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "How can I sell my things here?",
			Answer:   "You need to log in",
		},
		{
			Question: "Can i use this website for free?",
			Answer:   "Yes you can!",
		},
		{
			Question: "How to contact me?",
			Answer:   `Email us <a>jackchacne4@gmail.com</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
