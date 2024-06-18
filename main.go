package main

import (
	"log"
	"net/http"
)

func setupAPI() {
	manager := newManager()
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/ws", manager.serveWS)
}

func main() {
	setupAPI()
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Hello websocket")
}
