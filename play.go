package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.ReadInConfig()

	go forever()
	select {}
}

func forever() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://forumhouse.zendesk.com/api/v2/search.json?query=status<solved+tags:hr", nil)
	req.SetBasicAuth(viper.GetString("login"), viper.GetString("password"))

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