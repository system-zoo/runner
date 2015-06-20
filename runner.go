package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]
		generateTraffic(url)
	}
}

func generateTraffic(url string) {
	for {
		time.Sleep(2 * time.Second)
		Get(url)
	}
}

func Get(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)
	fmt.Printf("%s", body)
}
