package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
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
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err: no .env file found")
	}

	SID := os.Getenv("TWILIO_ACCOUNT_SID")
	TOKEN := os.Getenv("TWILIO_AUTH_TOKEN")
	m_phone := os.Getenv("M_PHONE")
	fmt.Printf(SID,TOKEN,m_phone)

	client := twilio.NewRestClient()

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+12016376817")
	params.SetFrom(m_phone)
	params.SetBody("yo")

	resp,err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	
	/*
	params := twilio.rest.api.v2010.Api.CreateMessageParams{}
	params.SetTo("+12016376817")
	params.SetFrom(m_phone)
	params.SetBody("yo")

	resp,err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})
	*/
	
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
