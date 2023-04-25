package main

import (
	"GoObjectStorage/apiServer/global"
	"GoObjectStorage/apiServer/heartbeat"
	"GoObjectStorage/apiServer/locate"
	"GoObjectStorage/apiServer/objects"
	"log"
	"net/http"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(global.LISTEN_ADDRESS, nil))
}
