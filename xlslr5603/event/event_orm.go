package event

import (
	"log"

	"github.com/wangsying/rfid/xlslr5603/db"
)

func init() {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()
	orm.AutoMigrate(&TagData{}, &ExceptionData{})
}

// NewEventOrm 创建一个事件存储操作对象
func NewEventOrm() EventOrm {
	return &eventOrm{}
}

// EventOrm 读写器主动事件,数据存储操作
type EventOrm interface {
	Readed(tag *TagData)
	Coming(tag *TagData)
	Exception(ex *ExceptionData)
	Heart(h *Heart)
	SyncTime()
}

type eventOrm struct {
}

func (e *eventOrm) Readed(tag *TagData) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&TagData{}).Create(tag)
}

func (e *eventOrm) Coming(tag *TagData) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&TagData{}).Create(tag)
}

func (e *eventOrm) Exception(ex *ExceptionData) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&ExceptionData{}).Create(ex)

}

func (e *eventOrm) Heart(h *Heart) {

}

func (e *eventOrm) SyncTime() {

}
