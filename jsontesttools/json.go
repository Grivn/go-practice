package jsontesttools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	if _, err := os.Stat("test"); os.IsNotExist(err) {
		os.Mkdir("test", 0777)
		os.Chmod("test", 0777)
	}
}

func GenerateJsonFile(res interface{}, prefix string) {
	json1, errs := json.MarshalIndent(res, "", "\t\t")
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	err := ioutil.WriteFile("test/"+prefix+".json", json1, 0644)
	if err != nil {
		fmt.Println("写文件失败")
	}
}
