package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IndexerRequest struct {
	Sender    string
	Receiver  string
	Amount    string
	Timestamp string
}

var baseUri string

func SetBaseURI(uri string) {
	baseUri = uri
}

func SendTestRequest() {
	requestUrl := baseUri
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", resBody)
}

func AddNewTransaction(r *IndexerRequest) {
	requestUrl := fmt.Sprint(baseUri, "/tx/add?sender=", r.Sender, "&receiver=", r.Receiver, "&amount=", r.Amount, "&timestamp=", r.Timestamp)
	req, err := http.NewRequest(http.MethodPut, requestUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("version", "v1")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", resBody)
}
