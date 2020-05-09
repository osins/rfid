package event

import (
	"encoding/json"
	"log"
	"strings"

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
	s, err := c.GetRawData() //把  body 内容读入字符串 s
	if err != nil {
		log.Println("gin context error: " + err.Error())
	}

	deviceName := c.Param("device_name")
	readerName, _ := jsonparser.GetString(s, "reader_name")
	eventType, _ := jsonparser.GetString(s, "event_type")
	remoteAddr := c.ClientIP()

	log.Println("device_name: " + deviceName)
	log.Println("reader_name: " + readerName)
	log.Println("event_type: " + eventType)

	dev := &Device{
		DeviceName: deviceName,
		ReaderName: readerName,
	}

	orm := New()
	orm.CreateOrUpdateDevice(dev)

	switch strings.TrimSpace(eventType) {
	case "tag_read":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &TagData{}
			json.Unmarshal(value, &tag)

			taglog := &TagLog{}
			json.Unmarshal(value, &taglog)

			tag.DeviceName = strings.TrimSpace(deviceName)
			tag.ReaderName = strings.TrimSpace(readerName)
			tag.EventType = strings.TrimSpace(eventType)
			tag.RemoteAddr = strings.TrimSpace(remoteAddr)

			taglog.DeviceName = strings.TrimSpace(deviceName)
			taglog.ReaderName = strings.TrimSpace(readerName)
			taglog.EventType = strings.TrimSpace(eventType)
			taglog.RemoteAddr = strings.TrimSpace(remoteAddr)

			ant := &Antenna{
				DeviceName: deviceName,
				Antenna:    tag.Antenna,
				Protocol:   tag.Protocol,
			}

			orm.TagLog(taglog)
			orm.Readed(tag)
			orm.CreateOrUpdateAntenna(ant)
			orm.AntennaReadCountAdd(ant)

			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)
		}, "event_data")
	case "tag_coming":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			tag := &TagData{}
			json.Unmarshal(value, &tag)

			taglog := &TagLog{}
			json.Unmarshal(value, &taglog)

			tag.DeviceName = strings.TrimSpace(deviceName)
			tag.ReaderName = strings.TrimSpace(readerName)
			tag.EventType = strings.TrimSpace(eventType)
			tag.RemoteAddr = strings.TrimSpace(remoteAddr)

			taglog.DeviceName = strings.TrimSpace(deviceName)
			taglog.ReaderName = strings.TrimSpace(readerName)
			taglog.EventType = strings.TrimSpace(eventType)
			taglog.RemoteAddr = strings.TrimSpace(remoteAddr)

			ant := &Antenna{
				DeviceName: deviceName,
				Antenna:    tag.Antenna,
				Protocol:   tag.Protocol,
			}

			orm.TagLog(taglog)
			orm.Readed(tag)
			orm.CreateOrUpdateAntenna(ant)
			orm.AntennaReadCountAdd(ant)

			log.Println("tag_epc: ", tag.Epc)
			log.Println("tag_bank_data: ", tag.BankData)
			log.Println("tag_antenna: ", tag.Antenna)
		}, "event_data")
	case "reader_exception":
		jsonparser.ArrayEach(s, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			ex := &ExceptionData{}
			json.Unmarshal(value, &ex)

			ex.DeviceName = strings.TrimSpace(deviceName)
			ex.ReaderName = strings.TrimSpace(readerName)
			ex.EventType = strings.TrimSpace(eventType)
			ex.RemoteAddr = strings.TrimSpace(remoteAddr)

			log.Println("err_code: ", ex.ErrCode)
			log.Println("err_string: ", ex.ErrString)
			log.Println("timestamp: ", ex.Timestamp)

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
