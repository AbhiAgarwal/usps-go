package usps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

type PickUpRequest struct {
	FirmName     string `xml:"FirmName"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	Urbanization string `xml:"Urbanization"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
}

type Package struct {
	ServiceType string `xml:"ServiceType"`
	Count       string `xml:"Count"`
}

type PickupChangeRequest struct {
	FirstName           string `xml:"FirstName"`
	LastName            string `xml:"LastName"`
	FirmName            string `xml:"FirmName"`
	SuiteOrApt          string `xml:"SuiteOrApt"`
	Address2            string `xml:"Address2"`
	Urbanization        string `xml:"Urbanization"`
	City                string `xml:"City"`
	State               string `xml:"State"`
	ZIP5                string `xml:"ZIP5"`
	ZIP4                string `xml:"ZIP4"`
	Phone               string `xml:"Phone"`
	Extension           string `xml:"Extension"`
	Package             `xml:"Package"`
	EstimatedWeight     string `xml:"EstimatedWeight"`
	PackageLocation     string `xml:"PackageLocation"`
	SpecialInstructions string `xml:"SpecialInstructions"`
	ConfirmationNumber  string `xml:"ConfirmationNumber"`
}

type PickUpInquiryRequest struct {
	FirmName           string `xml:"FirmName"`
	SuiteOrApt         string `xml:"SuiteOrApt"`
	Address2           string `xml:"Address2"`
	Urbanization       string `xml:"Urbanization"`
	City               string `xml:"City"`
	State              string `xml:"State"`
	ZIP5               string `xml:"ZIP5"`
	ZIP4               string `xml:"ZIP4"`
	ConfirmationNumber string `xml:"ConfirmationNumber"`
}

type CarrierPickupAvailabilityResponse struct {
	FirmName     string `xml:"FirmName"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
	DayOfWeek    string `xml:"DayOfWeek"`
	Date         string `xml:"Date"`
	CarrierRoute string `xml:"CarrierRoute"`
	Error        string `xml:"Error"`
}

type CarrierPickupChangeResponse struct {
	FirstName           string `xml:"FirstName"`
	LastName            string `xml:"LastName"`
	FirmName            string `xml:"FirmName"`
	SuiteOrApt          string `xml:"SuiteOrApt"`
	Address2            string `xml:"Address2"`
	Urbanization        string `xml:"Urbanization"`
	City                string `xml:"City"`
	State               string `xml:"State"`
	ZIP5                string `xml:"ZIP5"`
	ZIP4                string `xml:"ZIP4"`
	Phone               string `xml:"Phone"`
	Extension           string `xml:"Extension"`
	Package             `xml:"Package"`
	EstimatedWeight     string `xml:"EstimatedWeight"`
	PackageLocation     string `xml:"PackageLocation"`
	SpecialInstructions string `xml:"SpecialInstructions"`
	ConfirmationNumber  string `xml:"ConfirmationNumber"`
	DayOfWeek           string `xml:"DayOfWeek"`
	Date                string `xml:"Date"`
	Status              string `xml:"Status"`
	Error               string `xml:"Error"`
}

type CarrierPickupInquiryResponse struct {
	FirstName           string    `xml:"FirstName"`
	LastName            string    `xml:"LastName"`
	FirmName            string    `xml:"FirmName"`
	SuiteOrApt          string    `xml:"SuiteOrApt"`
	Address2            string    `xml:"Address2"`
	Urbanization        string    `xml:"Urbanization"`
	City                string    `xml:"City"`
	State               string    `xml:"State"`
	ZIP5                string    `xml:"ZIP5"`
	ZIP4                string    `xml:"ZIP4"`
	Phone               string    `xml:"Phone"`
	Extension           string    `xml:"Extension"`
	Package             []Package `xml:"Package"`
	EstimatedWeight     string    `xml:"EstimatedWeight"`
	PackageLocation     string    `xml:"PackageLocation"`
	SpecialInstructions string    `xml:"SpecialInstructions"`
	ConfirmationNumber  string    `xml:"ConfirmationNumber"`
	DayOfWeek           string    `xml:"DayOfWeek"`
	Date                string    `xml:"Date"`
	Error               string    `xml:"Error"`
}

type Error struct {
	Number      string `xml:"Number"`
	Description string `xml:"Description"`
	Source      string `xml:"Source"`
}

func (U *USPS) PickupAvailability(pickup PickUpRequest) CarrierPickupAvailabilityResponse {
	result := CarrierPickupAvailabilityResponse{}
	if U.Username == "" {
		fmt.Println("Username is missing")
		return result
	}

	xmlOut, err := xml.Marshal(pickup)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("CarrierPickupAvailability&XML=")
	urlToEncode := "<CarrierPickupAvailabilityRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</CarrierPickupAvailabilityRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body := U.GetRequestHTTPS(requestURL.String())
	if body == nil {
		return result
	}

	if len(result.Address2) > 0 {
		bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
		err = xml.Unmarshal([]byte(bodyHeaderless), &result)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		return result
	}
	errorResult := Error{}
	errorBody := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(errorBody), &errorResult)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	result.Error = errorResult.Description
	return result

}

func (U *USPS) PickupChange(pickup PickupChangeRequest) CarrierPickupChangeResponse {
	result := CarrierPickupChangeResponse{}
	if U.Username == "" {
		fmt.Println("Username is missing")
		return result
	}

	xmlOut, err := xml.Marshal(pickup)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("CarrierPickupChange&XML=")
	urlToEncode := "<CarrierPickupChangeRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</CarrierPickupChangeRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body := U.GetRequestHTTPS(requestURL.String())
	if body == nil {
		return result
	}

	if len(result.Address2) > 0 {
		bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
		err = xml.Unmarshal([]byte(bodyHeaderless), &result)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		return result
	}
	errorResult := Error{}
	errorBody := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(errorBody), &errorResult)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	result.Error = errorResult.Description
	return result
}

func (U *USPS) PickupInquiry(pickup PickUpInquiryRequest) CarrierPickupInquiryResponse {
	result := CarrierPickupInquiryResponse{}
	if U.Username == "" {
		fmt.Println("Username is missing")
		return result
	}

	xmlOut, err := xml.Marshal(pickup)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("CarrierPickupInquiry&XML=")
	urlToEncode := "<CarrierPickupInquiryRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</CarrierPickupInquiryRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body := U.GetRequestHTTPS(requestURL.String())
	if body == nil {
		return result
	}

	if len(result.Address2) > 0 {
		bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
		err = xml.Unmarshal([]byte(bodyHeaderless), &result)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		return result
	}
	errorResult := Error{}
	errorBody := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(errorBody), &errorResult)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	result.Error = errorResult.Description
	return result
}
