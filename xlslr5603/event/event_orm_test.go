package event

import (
	"encoding/json"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	data := []byte(`{
	"reader_name": "silion_reader/192.168.1.100",
	"event_type": "tag_read",
	"epc": "E20041453116009820603EFF",
	"bank_data": "",
	"antenna": 7,
	"read_count": 30,
	"protocol": 5,
	"rssi": -63,
	"firstseen_timestamp": 1550734256000,
	"lastseen_timestamp": 1550734272000
	}`)

	tag := &TagData{}
	json.Unmarshal(data, &tag)

	taglog := &TagLog{}
	json.Unmarshal(data, &taglog)

	orm := New()
	orm.AutoMigrate()
	orm.TagLog(taglog)
	orm.Readed(tag)
	getTag := orm.GetByID(tag.ID)

	// 测试断言
	assert.Equal(t, tag.ID, getTag.ID, "对于保存的TagData数据,进行一个验证,看看是否保存成功了,验证的条件是保存的EPC和通过orm返回的ID获取到的TagData的EPC进行比较,如果相等则通过测试")
}
