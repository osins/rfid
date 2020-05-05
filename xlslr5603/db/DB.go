package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// NewDB singleton
func NewDB() (db *gorm.DB, err error) {
	if !envIsExists(".env") {
		if envIsExists("/usr/local/rfid/xlslr5603/.env") {
			err = godotenv.Load("/usr/local/rfid/xlslr5603/.env")
			if err != nil {
				log.Fatal("Error loading .env file: /usr/local/rfid/xlslr5603/.env")
				return nil, err
			}
		} else {
			err = godotenv.Load("E:\\go\\github.com\\wangsying\\rfid\\xlslr5603\\.env")
			if err != nil {
				log.Fatal("Error loading .env file: E:\\go\\github.com\\wangsying\\rfid\\xlslr5603\\.env")
				return nil, err
			}
		}
	} else {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return nil, err
		}
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	tablePrefix := os.Getenv("DB_TABLE_PREFIX")
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"

	log.Println(username + ":****@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local")

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database: " + err.Error())
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.LogMode(true)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(20)

	return
}

func envIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
