package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {


     // if the http method is over these time , the request will cancel
	timeOutContext,cancel := context.WithTimeout(context.Background(), time.Millisecond *500)

	defer cancel()
	// create http

	req, err := http.NewRequestWithContext(timeOutContext,http.MethodGet, "https://www.istockphoto.com/photos/kochi-india ", nil)

	if err != nil {
		panic(err)
	}

	//perform http request
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	//get image from response

	imageData, err := ioutil.ReadAll(res.Body)

	fmt.Println("downolad image of size %d/n", len(imageData))

}
