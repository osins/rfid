package dtype

import (
	"bytes"
	"database/sql/driver"
	"errors"
)

// DBJson database Json type defined
type DBJson []byte

// Value function
func (j DBJson) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

// Scan function
func (j *DBJson) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		//errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// MarshalJSON function
func (j DBJson) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON function
func (j *DBJson) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// IsNull function
func (j DBJson) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

// Equals function
func (j DBJson) Equals(j1 DBJson) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
