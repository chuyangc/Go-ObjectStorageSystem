package locate

import (
	"GoObjectStorage/dataServer/global"
	"GoObjectStorage/lib/rabbitmq"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(global.RABBITMQ_SERVER)
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		if Locate("./tmp" + "/objects/" + object) {
			q.Send(msg.ReplyTo, global.LISTEN_ADDRESS)
		}
	}
}
