package repository

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func ConnectLevelDB() *leveldb.DB {
	db, err := leveldb.OpenFile("/Users/captainamerica/Documents/ldb/sckseal", nil)
	if err != nil {
		log.Fatal("connect error!")
	}

	//	data, err := db.Get([]byte("sayhi"), nil)

	//	var sayhi model.Sayhi
	//	json.Unmarshal([]byte(string(data)), &sayhi)
	//	fmt.Printf(" Description: %s\n", sayhi.Description)
	return db
}
