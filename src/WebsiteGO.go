package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const PORT = 8080

type Sample_View struct {
	Href  string
	Title string
}

// handler: basic handler for the web server
func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("HackWeb.html")
	sampView := []Sample_View{{Href: "#", Title: "Doctor Portal"}}
	t.Execute(w, sampView)
}

func api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "no")
	case "POST":
		fmt.Fprintf(w, "yes")
	default:
		fmt.Fprintf(w, "No support yet!")
	}
}

func main() {
	// convert port to string
	port := strconv.Itoa(PORT)

	// running a file
	fmt.Printf("Server running: http://localhost:%s\n", port)

	// web server starts at this directory
	fs := http.FileServer(http.Dir("./"))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
