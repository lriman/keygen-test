package generator

import (
	"github.com/jinzhu/gorm"
	"github.com/keygen-test/repository"
	"time"
	"log"
)

func KeyGen(db *gorm.DB) {
	for {
		freeKeys, err := repository.FreeKeyCount(db)
		if err == nil {
			if freeKeys < 50 {
				for i := 1; i <= 100; i++ {
					_, err := repository.GenerateNewKey(db)
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