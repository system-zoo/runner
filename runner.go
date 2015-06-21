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
		// The test-service can also respond to urls with a number for the path.
		// Append the random number to the path section of the url.
		target.Path = fmt.Sprintf("%v", randomInt)
		body, err := Get(target)
		if err != nil {
			log.Println(err)
		} else {
			if !strings.Contains(string(body), target.Path) {
				log.Println("The response did not contain " + target.Path)
			}
		}
	}
}

func Get(target *url.URL) ([]byte, error) {
	log.Println(target)
	// Make the GET request to the target.
	response, err := http.Get(target.String())
	if err != nil {
		return nil, err
	}
	// Close the body once this variable gets out of scope.
	defer response.Body.Close()
	log.Println(response.Status)
	// Read all the bytes from the response.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
