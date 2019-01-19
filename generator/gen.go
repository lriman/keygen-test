package generator

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"github.com/keygen-test/repository"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getKey(ind []int) string{
	s := ""
	for _, i := range ind{
		s += string(letterRunes[i])
	}
	return s
}

func KeyGen(db *gorm.DB) {
	var err error
	var query string
	vals := []interface{}{}

	lastKey := "0000"
	k, err := repository.GetLastKey(db)
	if err == nil {
		lastKey =  k.Key
	}


	lk := []int{0,0,0,0}
	for lk[0] = 0;  lk[0] < len(letterRunes); lk[0]++ {
		for lk[1] = 0;  lk[1] < len(letterRunes); lk[1]++ {
			for lk[2] = 0;  lk[2] < len(letterRunes); lk[2]++ {
				for lk[3] = 0;  lk[3] < len(letterRunes); lk[3]++ {
					nextKey := getKey(lk)

					if nextKey > lastKey{
						if len(vals) == 0{
							query = "INSERT INTO secret_keys(key) VALUES "
							vals = vals[:0]
						}

						query += "(?),"
						vals = append(vals, nextKey)

						if len(vals) == 200{
							query = query[0:len(query)-1]

							tx := db.Begin()
							err = tx.Exec(query, vals...).Error
							for err != nil{
								log.Println("Can't bath next pack of keys:", err)
								tx.Rollback()

								time.Sleep(5*time.Second)
								err = db.Exec(query, vals...).Error
							}
							tx.Commit()
							vals = vals[:0]

							/*
							Можно контролировать рост таблицы и не генерировать все ключи сразу.
							А, например, поддерживать 100К / 200К / 300К свободных ключей
							в зависимости от скорости их выдачи
							 */
							time.Sleep(5*time.Second)
						}
					}
				}
			}
		}
	}
}
