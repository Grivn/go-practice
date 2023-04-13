package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// convert the property in JSON which type is number or string.

type FloatString float64

func (f *FloatString) UnmarshalJSON(bs []byte) error {
	var v float64
	if bytes.HasPrefix(bs, []byte(`"`)) {
		var s string
		err := json.Unmarshal(bs, &s)
		if err != nil {
			return err
		}
		v, err = strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
	} else {
		err := json.Unmarshal(bs, &v)
		if err != nil {
			return err
		}
	}
	*f = FloatString(v)
	return nil
}

type A struct {
	F FloatString `json:"f"`
}

func main() {
	a := `{"f": "123"}`
	var f A
	_ = json.Unmarshal([]byte(a), &f)
	fmt.Println(f.F)
	fmt.Println(reflect.TypeOf(f.F))
}
