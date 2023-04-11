package main

import (
	"GoObjectStorage/objects"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
