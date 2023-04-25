package main

import (
	"GoObjectStorage/dataServer/global"
	"GoObjectStorage/dataServer/heartbeat"
	"GoObjectStorage/dataServer/locate"
	"GoObjectStorage/dataServer/objects"
	"log"
	"net/http"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(global.LISTEN_ADDRESS, nil))
}
