package server

import (
	"fmt"
	"html/template"
	"net/http"
)

type Server struct{}

type Message struct {
	Content string
	User    string
	Time    string
}

type HomeTemplateData struct {
	Messages []Message
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Run() {
	http.HandleFunc("/", s.handleHome)
	http.HandleFunc("/message", s.handleMessage)

	// serve static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

func (s Server) handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	messages := []Message{
		{
			Content: "Hello World!",
			User:    "John Doe",
			Time:    "12:00",
		},
		{
			Content: "Hello World!",
			User:    "John Doe",
			Time:    "12:00",
		},
		{
			Content: "Hello World!",
			User:    "John Doe",
			Time:    "12:00",
		},
		{
			Content: "Hello World!",
			User:    "John Doe",
			Time:    "12:00",
		},
		{
			Content: "Hello World!",
			User:    "John Doe",
			Time:    "12:00",
		},
	}

	data := HomeTemplateData{
		Messages: messages,
	}

	t := template.Must(template.ParseFiles("web/templates/home.html"))
	t.Execute(w, data)
}

func (s Server) handleMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
}
