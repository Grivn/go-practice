package main

import (
	"encoding/json"
	"fmt"
	//"strconv"
	//"log"
)

func main() {

	jsonStr := `[
      {
        "dept": "IT",
        "condition": {
          "employee": [
            "emp1"
          ]
        }
      },
      {
        "dept": "HR",
        "condition": {
          "employee": [
            "emp2",
            "emp3"
          ]
        }
      }
    ]`

	empMap := `{"emp1": "14325", "emp3": "49184", "emp2": "21518"}`

	type GetEmployee []struct {
		Dept      string `json:"dept"`
		Condition struct {
			Employee []string `json:"employee"`
		} `json:"condition"`
	}

	var empResponse GetEmployee
	var ids map[string]string
	unmarshallingError := json.Unmarshal([]byte(string(jsonStr)), &empResponse)
	if unmarshallingError != nil {
		fmt.Println(unmarshallingError.Error())
	}
	json.Unmarshal([]byte(empMap), &ids)
	fmt.Println(empResponse)
	fmt.Println(ids)

	for i, e := range empResponse {
		fmt.Println(e)
		for j, val := range empResponse[i].Condition.Employee {
			if _, ok := ids[val]; ok {
				empResponse[i].Condition.Employee[j] = ids[val]
			}
		}
	}

	fmt.Println(empResponse)

}
