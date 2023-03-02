package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	msg := ErrMsg{
		Err: fmt.Errorf("canoot do it"),
	}

	raw, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(raw))

	msg2 := ErrMsg2{
		Err: fmt.Errorf("canoot do it").Error(),
	}

	raw2, err := json.Marshal(msg2)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(raw2))
}

type ErrMsg struct {
	Err error `json:"err,omitempty"`
}

type ErrMsg2 struct {
	Err string `json:"err,omitempty"`
}
