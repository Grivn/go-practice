package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(time.Millisecond)
		harm()
	}
}

func harm() {
	body := &harmBody{companyName: "wow"}
	raw, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, "https://api.byted.icu/api/life/harm", bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	httpClient := http.Client{
		Timeout: time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	msg, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(msg))
}

type harmBody struct {
	companyName string
}
