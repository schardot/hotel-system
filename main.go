package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func bookPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func submitBooking(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	email := r.FormValue("email")
	tickets := r.FormValue("tickets")
	fmt.Printf("name: %s, email: %s, tickets %s", name, email, tickets)
	fmt.Fprintf(w, "Thank you, %s! Your booking has been received.", name)
}

func main() {
	http.HandleFunc("/book", bookPage)
	http.HandleFunc("/submit", submitBooking)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
