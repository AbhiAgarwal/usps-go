package usps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

type RateRequest struct {
	XMLName        xml.Name `xml:Package"`
	Revision       string   `xml:"Revision"`
	Service        string   `xml:"Service"`
	ZipOrigination string   `xml:"ZipOrigination"`
	ZipDestination string   `xml:"ZipDestination"`
	Pounds         string   `xml:"Pounds"`
	Ounces         string   `xml:"Ounces"`
	Container      string   `xml:"Container"`
	Size           string   `xml:"Size"`
	Width          string   `xml:"Width"`
	Length         string   `xml:"Length"`
	Height         string   `xml:"Height"`
	Girth          string   `xml:"Girth"`
}

type RateV4Response struct {
}

func (U *USPS) RateDomestic(rate RateRequest) RateV4Response {
	result := RateV4Response{}
	if U.Username == "" {
		fmt.Println("Username is missing")
		return result
	}

	xmlOut, err := xml.Marshal(rate)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("RateV4&XML=")
	urlToEncode := "<RateV4Request USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</RateV4Request>"
	requestURL.WriteString(URLEncode(urlToEncode))
	fmt.Println(requestURL.String())

	body := U.GetRequest(requestURL.String())
	if body == nil {
		return result
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	fmt.Println(bodyHeaderless)
	err = xml.Unmarshal([]byte(bodyHeaderless), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}

	return result
}
