package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	maindata := struct {
		Title string
		H1    string
	}{Title: "joe mama!!!", H1: "HELLO WORLD BUT BIG AND BLUE"}
	indextemp, _ := template.ParseFiles("templates/index.tmpl")
	main := func(w http.ResponseWriter, req *http.Request) {
		indextemp.Execute(w, maindata)
	}
	joedata := struct {
		Title string
		H1    string
		P     string
	}{Title: "joe mama api lol!", H1: "Random Joe Mama Joke API", P: "This is a simple REST API I made to provide random joe mama jokes. This also happens to be my first web app made with Golang, and my first web app in general."}
	joetemp, _ := template.ParseFiles("templates/joekes.tmpl")
	joemama := func(w http.ResponseWriter, req *http.Request) {
		joetemp.Execute(w, joedata)
	}
	apifun := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"content":"hello worlD! (for now...)"}`))
	}
	http.HandleFunc("/", main)
	http.HandleFunc("/joekes", joemama)
	http.HandleFunc("/joekes/api", apifun)
	log.Fatal(http.ListenAndServe(":80", nil))
}
