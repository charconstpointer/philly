package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(2 * time.Second)
		r, err := http.Get("http://localhost:5010/")
		if err != nil {
			log.Println(err.Error())
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("%s\n", b)
	}
}
