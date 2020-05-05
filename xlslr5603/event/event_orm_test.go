package event

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestNewEventOrm(t *testing.T) {
	tag := &TagData{
		Epc: "2343423",
	}

	ex := &ExceptionData{
		ErrCode:   1,
		ErrString: "this is a test.",
	}

	orm := NewEventOrm()
	orm.Readed(*tag)
	orm.Exception(*ex)
}
