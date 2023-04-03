package main

import (
	"encoding/json"
	"fmt"
	"github.com/henrylee2cn/ameda"
)

var j = `{
	"a": "b"
}`

func main() {
	fmt.Println(j)
	var item interface{}

	if err := json.Unmarshal([]byte(j), &item); err != nil {
		panic(err)
	}
	fmt.Println(item)
	itemMap := item.(map[string]interface{})
	fmt.Println(itemMap["a"])

	fromMsg := IntMsg{ID: int64(1)}

	toMsg := StrMsg{}

	m := make(map[string]interface{})

	raw, err := json.Marshal(fromMsg)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(raw, &m); err != nil {
		panic(err)
	}

	m, err = numberToString(m, "id")
	if err != nil {
		panic(err)
	}

	raw, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(raw, &toMsg); err != nil {
		panic(err)
	}
	fmt.Println(toMsg)
}

type StrMsg struct {
	ID string `json:"id,omitempty"`
}

type IntMsg struct {
	ID int64 `json:"id,omitempty"`
}

func numberToString(jsonMap map[string]interface{}, key string) (map[string]interface{}, error) {
	rawID, ok := jsonMap[key]
	if !ok {
		// cannot find 'id' in panel's json map.
		return jsonMap, nil
	}

	float64ID, ok := rawID.(float64)
	if ok {
		int64ID, err := ameda.Float64ToInt64(float64ID)
		if err != nil {
			return jsonMap, err
		}
		jsonMap[key] = ameda.Int64ToString(int64ID)
		return jsonMap, nil
	}

	return jsonMap, nil
}
