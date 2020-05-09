package event

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/wangsying/rfid/xlslr5603/db"
)

// New 创建一个事件存储操作对象
func New() ORM {
	return &orm{}
}

// ORM 读写器主动事件,数据存储操作
type ORM interface {
	AutoMigrate()
	CreateOrUpdateDevice(dev *Device)
	CreateOrUpdateAntenna(ant *Antenna)
	AntennaReadCountAdd(ant *Antenna)
	GetByID(id uint) TagData
	TagLog(tag *TagLog)
	Readed(tag *TagData)
	Coming(tag *TagLog)
	Exception(ex *ExceptionData)
	Heart(h *Heart)
	SyncTime()
}

type orm struct {
}

func (e *orm) AutoMigrate() {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.AutoMigrate(&Device{}, &Antenna{}, &TagData{}, &TagLog{}, &ExceptionData{})

	orm.Model(&TagData{}).AddIndex("idx_epc", "epc")
	orm.Model(&TagData{}).AddIndex("idx_reader_name", "reader_name")
	orm.Model(&TagData{}).AddIndex("idx_event_type", "event_type")
	orm.Model(&TagData{}).AddIndex("idx_antenna", "antenna")
	orm.Model(&TagData{}).AddIndex("idx_reader_event_antenna", "reader_name", "antenna", "event_type")
}

func (e *orm) CreateOrUpdateDevice(dev *Device) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	var count int
	orm.Model(&Device{}).Where("device_name = ?", dev.DeviceName).Count(&count)
	if count == 0 {
		orm.Model(&Device{}).Create(&dev)
		return
	}

	orm.Model(&Device{}).Where("device_name = ?", dev.DeviceName).Update(&dev)
}

func (e *orm) CreateOrUpdateAntenna(ant *Antenna) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	var count int
	orm.Model(&Antenna{}).Where("device_name = ? and antenna = ?", ant.DeviceName, ant.Antenna).Count(&count)
	if count == 0 {
		orm.Model(&Antenna{}).Create(&ant)
		return
	}

	orm.Model(&Antenna{}).Where("device_name = ? and antenna = ?", ant.DeviceName, ant.Antenna).Update(&ant)
}

func (e *orm) AntennaReadCountAdd(ant *Antenna) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&Antenna{}).Where("device_name = ? and antenna = ?", ant.DeviceName, ant.Antenna).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
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

func (e *orm) TagLog(tag *TagLog) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&TagLog{}).Create(&tag)
}

func (e *orm) Readed(tag *TagData) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	var count int
	orm.Model(&TagData{}).Where("event_type = ? and epc = ?", tag.EventType, tag.Epc).Count(&count)
	if count == 0 {
		orm.Model(&TagData{}).Create(&tag)
		return
	}

	orm.Model(&TagData{}).Where("event_type = ? and epc = ?", tag.EventType, tag.Epc).Update(&tag)
}

func (e *orm) Coming(tag *TagLog) {
	orm, err := db.NewDB()
	if err != nil {
		log.Println(err)
	}

	defer orm.Close()

	orm.Model(&TagLog{}).Create(&tag)
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
