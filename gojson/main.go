package main

import (
	"github.com/Grivn/gojson"
	"os"
)

func main() {
	input, err := os.ReadFile("testdata/camel_snake/" + file)

	gojson.JSONSchemaCamel2Snake(input)
}
