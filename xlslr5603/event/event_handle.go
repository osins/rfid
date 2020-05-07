package event

import (
	"encoding/json"
	"log"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
)

// NewHandle 创建Reader主动事件(来自Handle接口)
func NewHandle() Handle {
	return &handle{}
}

// Handle Reader 主动事件接口
type Handle interface {
	ReaderEventHandle(c *gin.Context)
}

type handle struct {
}

func (e *handle) ReaderEventHandle(c *gin.Context) {
	s, _ := c.GetRawData() //把  body 内容读入字符串 s

	deviceName := c.Param("device_name")
	readerName, _ := jsonparser.GetString(s, "reader_name")
	eventType, _ := jsonparser.GetString(s, "event_type")
	remoteAddr := c.ClientIP()

	log.Println("device_name: " + deviceName)
	log.Println("reader_name: " + readerName)
	log.Println("event_type: " + eventType)

	switch eventType {
	case "tag_read":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &TagData{}

			json.Unmarshal(value, &tag)

			tag.ReadEvent.DeviceName = deviceName
			tag.ReadEvent.ReaderName = readerName
			tag.ReadEvent.EventType = eventType
			tag.ReadEvent.RemoteAddr = remoteAddr

			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)

			orm := New()
			orm.Readed(tag)
		}, "event_data")
	case "tag_coming":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &TagData{}

			json.Unmarshal(value, &tag)

			tag.ReadEvent.DeviceName = deviceName
			tag.ReadEvent.ReaderName = readerName
			tag.ReadEvent.EventType = eventType
			tag.ReadEvent.RemoteAddr = remoteAddr

			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)

			orm := New()
			orm.Readed(tag)
		}, "event_data")
	case "reader_exception":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			ex := &ExceptionData{}
			json.Unmarshal(value, &ex)

			ex.ReadEvent.DeviceName = deviceName
			ex.ReadEvent.ReaderName = readerName
			ex.ReadEvent.EventType = eventType
			ex.ReadEvent.RemoteAddr = remoteAddr

			log.Println("err_code: ", ex.ErrCode)
			log.Println("err_string: ", ex.ErrString)
			log.Println("timestamp: ", ex.Timestamp)

			orm := New()
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
