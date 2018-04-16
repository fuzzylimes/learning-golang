package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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
	defer res.Body.Close()

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
