package usps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

type TrackResponse struct {
	TrackInfo struct {
		TrackSummary string `xml:"TrackSummary"`
	} `xml:"TrackInfo"`
}

func (U *USPS) TrackPackage(trackingID string) TrackResponse {
	result := TrackResponse{}

	var requestURL bytes.Buffer
	requestURL.WriteString("TrackV2&XML=")
	urlToEncode := "<TrackRequest USERID=\"" + U.Username + "\">"
	urlToEncode += "<TrackID ID=\"" + trackingID + "\"></TrackID>"
	urlToEncode += "</TrackRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body := GetRequest(requestURL.String())
	if body == nil {
		return result
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err := xml.Unmarshal([]byte(bodyHeaderless), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}

	return result
}
