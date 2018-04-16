package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type people struct {
	Number  int      `json:"number"`
	Message string   `json:"message"`
	Person  []person `json:"people"`
}

type person struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		log.Fatal("Could not decode json.")
	}
	return out.Bytes(), err
}

func main() {
	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
	textBytes := []byte(text)

	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println(err)
		return
	}

	textFormat, _ := prettyPrint(textBytes)
	fmt.Printf("%s\n", textFormat)
	fmt.Println(people1)
	fmt.Println(people1.Number)
	fmt.Println(people1.Message)
	fmt.Println(people1.Person[0].Name)

}
