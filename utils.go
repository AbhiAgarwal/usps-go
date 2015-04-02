package usps

import (
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
