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

func GetRequest(requestURL string) []byte {
	currentURL := base + requestURL
	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
