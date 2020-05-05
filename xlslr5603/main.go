package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wangsying/rfid/xlslr5603/event"
)

func eventHandle(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s

	readerName, _ := jsonparser.GetString(s, "reader_name")
	eventType, _ := jsonparser.GetString(s, "event_type")

	log.Println("reader_name: " + readerName)
	log.Println("event_type: " + eventType)

	switch eventType {
	case "tag_read":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &event.TagData{}
			json.Unmarshal(value, &tag)
			tag.ReaderName = readerName
			tag.EventType = eventType
			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)

			orm := event.NewEventOrm()
			orm.Readed(tag)
		}, "event_data")
	case "tag_coming":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &event.TagData{}
			json.Unmarshal(value, &tag)
			tag.ReaderName = readerName
			tag.EventType = eventType
			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)

			orm := event.NewEventOrm()
			orm.Readed(tag)
		}, "event_data")
	case "reader_exception":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			ex := &event.ExceptionData{}
			json.Unmarshal(value, &ex)
			ex.ReaderName = readerName
			ex.EventType = eventType
			log.Println("err_code: ", ex.ErrCode)
			log.Println("err_string: ", ex.ErrString)
			log.Println("timestamp: ", ex.Timestamp)

			orm := event.NewEventOrm()
			orm.Exception(ex)
		}, "event_data")
	case "heart_beat":
		h := &event.Heart{}
		json.Unmarshal(s, &h)
		log.Println("event_data: ", h.EventData)
	case "sync_time_req":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &event.TagData{}
			json.Unmarshal(value, &tag)
			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)
		}, "event_data")
	}

	log.Println("-----------------------------------------------")
}

func main() {
	listenHost := os.Getenv("LISTEN_HOST")
	listenPort := os.Getenv("LISTEN_PORT")

	log.Println("start rfid service " + listenHost + ":" + listenPort + ", request waiting ...")

	http.HandleFunc("/boyang/xlslr5603/active-request", eventHandle) //设置访问的路由
	err := http.ListenAndServe(listenHost+":"+listenPort, nil)       //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func envIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
