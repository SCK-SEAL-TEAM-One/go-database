package service

import (
	"encoding/json"
	"fmt"
	"model"
	"repository"
)

func ReadSayhiFormLevelDB() model.Sayhi {

	db := repository.ConnectLevelDB()
	defer db.Close()
	data, _ := db.Get([]byte("sayhi"), nil)

	var sayhi model.Sayhi
	json.Unmarshal([]byte(string(data)), &sayhi)
	fmt.Printf(" Description: %s\n", sayhi.Description)

	return sayhi
}
