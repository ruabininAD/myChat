package main

import (
	myHandler "chater/Handlers"
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("template/autorisation.html"))

func main() {

	fmt.Println("runServer")

	mux := http.NewServeMux()

	mux.HandleFunc("/", myHandler.Authorization)
	mux.HandleFunc("/chat", myHandler.Chat)

	fs := http.FileServer(http.Dir("template/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(":3000", mux)

}
