package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		log.Fatal("Could not decode json.")
	}
	return out.Bytes(), err
}

type people struct {
	Number  int      `json:"number"`
	Message string   `json:"message"`
	Person  []person `json:"people"`
}

type person struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}
