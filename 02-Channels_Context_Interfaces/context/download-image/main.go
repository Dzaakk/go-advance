package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*700)
	defer cancel()

	// create HTTP request without ctx
	//req, err := http.NewRequest(http.MethodGet, "http://placehold.it/2000x2000", nil)

	// create HTTP request with ctx
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	// perform HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// get data from HTTP response
	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("downloaded image of size %d\n", len(imageData))
}
