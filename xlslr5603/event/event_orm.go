package event

import (
	"log"
	"sync"

	"github.com/wangsying/rfid/xlslr5603/db"
)

var once sync.Once

func init() {
	once.Do(autoMigrate)
}

func autoMigrate() {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.AutoMigrate(&TagData{}, &ExceptionData{})

	orm.Model(&TagData{}).AddIndex("idx_epc", "epc")
	orm.Model(&TagData{}).AddIndex("idx_reader_name", "reader_name")
	orm.Model(&TagData{}).AddIndex("idx_event_type", "event_type")
	orm.Model(&TagData{}).AddIndex("idx_antenna", "antenna")
	orm.Model(&TagData{}).AddIndex("idx_reader_event_antenna", "reader_name", "antenna", "event_type")
}

// NewOrm 创建一个事件存储操作对象
func NewOrm() ORM {
	return &orm{}
}

// ORM 读写器主动事件,数据存储操作
type ORM interface {
	GetByID(id uint) TagData
	Readed(tag *TagData)
	Coming(tag *TagData)
	Exception(ex *ExceptionData)
	Heart(h *Heart)
	SyncTime()
}

type orm struct {
}

func (e *orm) GetByID(id uint) TagData {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	var tag TagData
	orm.Model(&TagData{}).Where("id = ?", id).First(&tag)

	return tag
}

func (e *orm) Readed(tag *TagData) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&TagData{}).Create(&tag)
}

func (e *orm) Coming(tag *TagData) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&TagData{}).Create(&tag)
}

func (e *orm) Exception(ex *ExceptionData) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&ExceptionData{}).Create(&ex)
}

func (e *orm) Heart(h *Heart) {

}

func (e *orm) SyncTime() {

}
