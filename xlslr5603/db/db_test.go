package db

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB()
	defer db.Close()

	if err != nil {
		t.Fatal(err.Error())
	}
}
