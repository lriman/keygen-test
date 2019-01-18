package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/keygen-test/models"
	"log"
)

type DatabaseConnection struct {
	conn   *gorm.DB
	ready  chan bool
	host   string
	usr    string
	pwd    string
	dbname string
}

func InitDatabaseConnection(host string, usr string, pwd string, dbname string) {
	DB = &DatabaseConnection{
		ready: make(chan bool, 1),
		conn: nil,
		host: host,
		usr: usr,
		pwd: pwd,
		dbname: dbname,
	}

	DB.connect()
	err := DB.conn.AutoMigrate(models.SecretKey{}).Error
	if err != nil{
		log.Panic(err)
	}

	DB.ready <- true
	go DB.reconnect()
	go DB.keyGen()
}

var DB *DatabaseConnection