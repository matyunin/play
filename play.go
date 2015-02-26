package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/kr/pretty"
)

func main() {
	go forever()
	select {}
}

func forever() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://forumhouse.zendesk.com/api/v2/search.json?query=status<solved+tags:hr", nil)
	req.SetBasicAuth("projects@forumhouse.ru/token", "py9in1GRCopEw5QgRlHg23qx2rlJ3jZXJxXlEK2L")

	for {	
		resp, _ := client.Do(req)

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var jsonBlob = []byte(string(body))
		var arbitaryJson interface{}

		_ = json.Unmarshal(jsonBlob, &arbitaryJson)

		fmt.Printf("%# v\n", pretty.Formatter(arbitaryJson.(map[string]interface {})["results"].([]interface {})[0].(map[string]interface {})["url"]))

		time.Sleep(time.Second * 3)
	}
}