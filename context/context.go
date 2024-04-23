package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// create http

	req, err := http.NewRequest(http.MethodGet, "https://www.istockphoto.com/photos/kochi-india ", nil)

	if err != nil {
		panic(err)
	}

	//perform http request
	res,err := http.DefaultClient.Do(req)


	if err!= nil{
		panic(err)
	}

	defer res.Body.Close()



	//get image from response 

	imageData,err := ioutil.ReadAll(res.Body)

	fmt.Println("downolad image of size %d/n", len(imageData))


}