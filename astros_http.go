package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	url := "http://api.open-notify.org/astros.json"

	astroClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "astro-tutorial")

	res, getErr := astroClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	} else if res.StatusCode != 200 {
		log.Fatal("Non 200 response returned.\n", res.StatusCode)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	textFormat, _ := prettyPrint(body)
	fmt.Printf("%s\n", textFormat)
	fmt.Println(people1.Number)
	fmt.Println(people1.Message)
	fmt.Println(people1.Person[0].Name)
	fmt.Printf("%T\n", res.Status)

}
