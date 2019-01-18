package db

import (
	"github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"log"
	"fmt"
)

func (db *DatabaseConnection) connect() {
	conn, err := gorm.Open("postgres", fmt.Sprintf(`
			host=%s
			user=%s
			password=%s
			dbname=%s
			sslmode=disable
		`, db.host, db.usr, db.pwd, db.dbname))

	if err != nil {
		log.Println("Could not connect to Database:", err)
	}

	log.Println("Sucess database connection")
	db.conn = conn
}

func (db *DatabaseConnection) reconnect() {

	select {
	case <-db.ready:
	default:
		return
	}

	defer func() { db.ready <- true }()

	if db.conn == nil {
		db.connect()
	}

	connErr := db.conn.DB().Ping()
	for connErr != nil {
		log.Println("Connection to Database was lost. Waiting for 5s...")
		db.conn.Close()

		time.Sleep(5 * time.Second)
		log.Println("Reconnecting to Database...")

		db.connect()
		connErr = db.conn.DB().Ping()
	}
}

func (db *DatabaseConnection) checkConnection(err error) {
	pqerr, ok := err.(*pq.Error)
	if ok && pqerr.Code[0:2] == "08" {
		go db.reconnect()
	}
	// PostgreSQL "Connection Exceptions" are class "08"
	// http://www.postgresql.org/docs/9.4/static/errcodes-appendix.html#ERRCODES-TABLE
}

func (db *DatabaseConnection) keyGen() {
	for {
		freeKeys, err := db.FreeKeyCount()
		if err == nil {
			if freeKeys < 50 {
				for i := 1; i <= 100; i++ {
					_, err := db.GenerateNewKey()
					if err != nil {
						log.Println("Generation error:", err)
					}
				}
			}
		} else {
			log.Println("Count error:", err)
		}
		time.Sleep(10 * time.Second)
	}
}