package event

import "github.com/jinzhu/gorm"

// Event Read读写器主动事件
type Event struct {
	ReaderName string `json:"reader_name"`
	EventType  string `json:"event_type"`
}

// TagData Read读写器事件中的标签数据
type TagData struct {
	gorm.Model
	Event
	Epc                string `json:"epc"`
	BankData           string `json:"bank_data"`
	Antenna            int    `json:"antenna"`
	ReadCount          int    `json:"read_count"`
	Protocol           int    `json:"protocol"`
	Rssi               int    `json:"rssi"`
	FirstseenTimestamp int64  `json:"firstseen_timestamp"`
	LastseenTimestamp  int64  `json:"lastseen_timestamp"`
}

// ExceptionData Read读写器事件中的异常数据
type ExceptionData struct {
	gorm.Model
	Event
	ErrCode   int    `json:"err_code"`
	ErrString string `json:"err_string"`
	Timestamp int64  `json:"timestamp"`
}

// Heart 心跳事件数据对象
type Heart struct {
	Event
	EventData int `json:"event_data"`
}

// SyncTime 同步时间请求事件数据对象
type SyncTime struct {
	Event
	EventData struct {
	} `json:"event_data"`
}

// SyncTimeResponse 同步时间返回给读写器的数据对象
type SyncTimeResponse struct {
	CommandType string `json:"command_type"`
	CommandData int64  `json:"command_data"`
}
