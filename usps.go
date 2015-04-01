package usps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	base string = "http://production.shippingapis.com"
)

type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Zip5     string
	Zip4     string
}

type USPS struct {
	Username string
	Password string
}

func (U *USPS) AddressVerification(address Address) {
	xmlOut, err := xml.Marshal(address)
	if err != nil {
		fmt.Println(err)
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("/ShippingAPITest.dll?API=Verify&XML=")
	requestURL.WriteString("<AddressValidateRequest USERID=\"" + U.Username + "\">")
	requestURL.WriteString(string(xmlOut))
	requestURL.WriteString("</AddressValidateRequest>")

	currentURL := base + requestURL.String()
	requestURLString, err := url.Parse(currentURL)
	fmt.Println(err)

	resp, err := http.Get(requestURLString.String())
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
