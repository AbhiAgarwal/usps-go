package usps

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func URLEncode(urlToEncode string) string {
	return url.QueryEscape(urlToEncode)
}

func (U *USPS) GetRequest(requestURL string) []byte {
	currentURL := ""
	if U.Production {
		currentURL += prodbase
	} else {
		currentURL += devbase
	}
	currentURL += requestURL

	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (U *USPS) GetRequestHTTPS(requestURL string) []byte {
	currentURL := ""
	if U.Production {
		currentURL += prodhttpsbase
	} else {
		currentURL += devhttpsbase
	}
	currentURL += requestURL

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(currentURL)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
