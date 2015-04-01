package usps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AddressValidateResponse struct {
	Address struct {
		Address2 string `xml:"Address2"`
		City     string `xml:"City"`
		State    string `xml:"State"`
		Zip5     string `xml:"Zip5"`
		Zip4     string `xml:"Zip4"`
	} `xml:"Address"`
}

type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Zip5     string
	Zip4     string
}

func (U *USPS) AddressVerification(address Address) AddressValidateResponse {
	xmlOut, err := xml.Marshal(address)
	result := AddressValidateResponse{}
	if err != nil {
		fmt.Println(err)
		return result
	}
	var requestURL bytes.Buffer
	requestURL.WriteString("/ShippingAPITest.dll?API=Verify&XML=")
	urlToEncode := "<AddressValidateRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</AddressValidateRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	currentURL := base + requestURL.String()
	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}

	return result
}
