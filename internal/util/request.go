package util

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Request(method string, url string, body []byte) []byte {
	fmt.Println("POST URL:", url)

	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth("mifos", "password")
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("fineract-platform-tenantid", "default")

	// skip ssl certs
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response Status:", response.Status)
	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response Body:", string(responseBody))
	return responseBody
}