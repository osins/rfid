package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wangsying/rfid/xlslr5603/event"
)

func main() {
	listenHost := os.Getenv("LISTEN_HOST")
	listenPort := os.Getenv("LISTEN_PORT")

	log.Println("start rfid service " + listenHost + ":" + listenPort + ", request waiting ...")

	event := event.NewHandle()
	http.HandleFunc("/boyang/xlslr5603/active-request", event.ReaderEventHandle) //设置访问的路由
	err := http.ListenAndServe(listenHost+":"+listenPort, nil)                   //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
