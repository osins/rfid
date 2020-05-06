package event

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestNewOrm(t *testing.T) {
	tag := &TagData{
		Epc: "2343423",
	}

	orm := NewOrm()
	orm.Readed(tag)
	getTag := orm.GetByID(tag.ID)

	// 测试断言
	assert.Equal(t, tag.ID, getTag.ID, "对于保存的TagData数据,进行一个验证,看看是否保存成功了,验证的条件是保存的EPC和通过orm返回的ID获取到的TagData的EPC进行比较,如果相等则通过测试")
}
