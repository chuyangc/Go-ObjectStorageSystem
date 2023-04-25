package heartbeat

import (
	"GoObjectStorage/dataServer/global"
	"GoObjectStorage/lib/rabbitmq"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(global.RABBITMQ_SERVER)
	defer q.Close()
	for {
		q.Publish("apiServer", global.LISTEN_ADDRESS)
		time.Sleep(5 * time.Second)
	}
}
