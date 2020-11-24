package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type info struct {
	ip int64 `json:"ip"`
}

func main() {
	var IP info
	U := os.Args[1]
	G, _ := http.Get(U)
        err := json.NewDecoder(G.Body).Decode(&IP)
        if err != nil {
         fmt.Println(err)
        }
	fmt.Printf("%+v\n", IP)
}
