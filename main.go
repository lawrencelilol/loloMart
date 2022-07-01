package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my great site! </h1>")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Lawrence </h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
	<li>Is there a free version? No, you need to pay for this </li>
	</ul>
	`)

}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "page not found", http.StatusNotFound)
	}
}

func main() {

	http.HandleFunc("/", pathHandler)

	log.Print("Listening on :8000...")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
