package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const milliBetweenRequests = 10

func main() {

	if len(os.Args) > 1 {
		url, err := url.Parse(os.Args[1])
		if err != nil {
			fmt.Printf("Invalid url %s\n", os.Args[1])
			fmt.Println(err)
			os.Exit(1)
		}
		generateTraffic(url)
	}
}

func generateTraffic(target *url.URL) {
	for {
		randomInt := RandInt(1, 1000)
		time.Sleep(milliBetweenRequests * time.Millisecond)
		// Append a random number on on the path section of the url.
		target.Path = fmt.Sprintf("%v", randomInt)
		Get(target)
	}
}

func Get(target *url.URL) {
	log.Println(target)
	// Make the GET request to the target.
	response, err := http.Get(target.String())
	if err != nil {
		fmt.Println(err)
	}
	// Close the body once this variable gets out of scope.
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	if !strings.Contains(string(body), target.Path) {
		log.Println("Response body did not contain " + target.Path)
	}
	log.Printf("%s %d\n", response.Status, response.StatusCode)
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
