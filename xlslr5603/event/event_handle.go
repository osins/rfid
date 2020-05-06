package event

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
)

// NewHandle 创建Reader主动事件(来自Handle接口)
func NewHandle() Handle {
	return &handle{}
}

// Handle Reader 主动事件接口
type Handle interface {
	ReaderEventHandle(w http.ResponseWriter, r *http.Request)
}

type handle struct {
}

func (e *handle) ReaderEventHandle(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s

	readerName, _ := jsonparser.GetString(s, "reader_name")
	eventType, _ := jsonparser.GetString(s, "event_type")

	log.Println("reader_name: " + readerName)
	log.Println("event_type: " + eventType)

	switch eventType {
	case "tag_read":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &TagData{}
			json.Unmarshal(value, &tag)
			tag.ReaderName = readerName
			tag.EventType = eventType
			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)

			orm := NewOrm()
			orm.Readed(tag)
		}, "event_data")
	case "tag_coming":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &TagData{}
			json.Unmarshal(value, &tag)
			tag.ReaderName = readerName
			tag.EventType = eventType
			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)

			orm := NewOrm()
			orm.Readed(tag)
		}, "event_data")
	case "reader_exception":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			ex := &ExceptionData{}
			json.Unmarshal(value, &ex)
			ex.ReaderName = readerName
			ex.EventType = eventType
			log.Println("err_code: ", ex.ErrCode)
			log.Println("err_string: ", ex.ErrString)
			log.Println("timestamp: ", ex.Timestamp)

			orm := NewOrm()
			orm.Exception(ex)
		}, "event_data")
	case "heart_beat":
		h := &Heart{}
		json.Unmarshal(s, &h)
		log.Println("event_data: ", h.EventData)
	case "sync_time_req":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &TagData{}
			json.Unmarshal(value, &tag)
			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)
		}, "event_data")
	}

	log.Println("-----------------------------------------------")
}
