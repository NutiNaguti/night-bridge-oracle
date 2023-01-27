package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request interface {
	New()
}

func SendTestRequest() {
	requestUrl := "http://localhost:1234"
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Add("version", "v1")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", resBody)
}

func AddNewTransaction() {

}
