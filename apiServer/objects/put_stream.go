package objects

import (
	"GoObjectStorage/apiServer/heartbeat"
	"GoObjectStorage/lib/objectstream"
	"fmt"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("can not find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
}
