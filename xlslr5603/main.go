package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wangsying/rfid/xlslr5603/event"
)

func main() {
	// 数据迁移(初始化数据结构)
	event.New().AutoMigrate()

	// 载入服务启动后的host和port配置
	listenHost := os.Getenv("LISTEN_HOST")
	listenPort := os.Getenv("LISTEN_PORT")

	log.Println("start rfid service " + listenHost + ":" + listenPort + ", request waiting ...")

	event := event.NewHandle()

	r := gin.Default()
	r.GET("/boyang/xlslr5603/active-request/:device_name", event.ReaderEventHandle)
	r.Run(listenHost + ":" + listenPort) // listen and serve on 0.0.0.0:8080
}
