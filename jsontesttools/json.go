package main

import (
	"encoding/json"
	"fmt"
)

type TestSchema struct {
	Desc     string     `json:"desc"`
	DataList []TestData `json:"data_list"`
}

type TestData struct {
	ID int64 `json:"id"`
}

func main() {
	s := TestSchema{Desc: "test", DataList: []TestData{{ID: 1}, {ID: 2}}}
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	processor(data)
}

func processor(data []byte) {
	itemMap := make(map[string]interface{})
	if err := json.Unmarshal(data, &itemMap); err != nil {
		panic(err)
	}
	for key, item := range itemMap {
		switch v := item.(type) {
		case []interface{}:

		case []TestData:
			fmt.Println(key, v)
		case string:
			fmt.Println(key, v)
		default:
			panic("unknown")
		}
	}
}
