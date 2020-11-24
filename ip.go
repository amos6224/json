package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type info struct {
	Ip string `json:"ip"`
}

func main() {
	resp, _ := http.Get("https://ipinfo.io")
	defer resp.Body.Close()

	jsonDataFromHttp, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e)
	}
	var ip info
	e = json.Unmarshal([]byte(jsonDataFromHttp), &ip)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("%+v\n", ip)
}
