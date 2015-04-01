package usps

import "testing"

func TestPickupAPI(t *testing.T) {
	var usps USPS
	usps.Username = ""
	usps.Password = ""

	var address Address
	address.Address2 = "6406 Ivy Lane"
	address.City = "Greenbelt"
	address.State = "MD"

	output := usps.AddressVerification(address)
	if output.Address.Address2 != "6406 IVY LN" {
		t.Error("Address is incorrect")
	}
}

func TestZipCodeLookup(t *testing.T) {
	var usps USPS
	usps.Username = ""
	usps.Password = ""

	var address Address
	address.Address2 = "6406 Ivy Lane"
	address.City = "Greenbelt"
	address.State = "MD"

	output := usps.ZipCodeLookup(address)
	if output.Address.Address2 != "6406 IVY LN" {
		t.Error("Address is incorrect")
	}
}

func TestCityStateLookup(t *testing.T) {
	var usps USPS
	usps.Username = ""
	usps.Password = ""

	var address ZipCode
	address.Zip5 = "90210"

	output := usps.CityStateLookup(address)
	if output.ZipC.Zip5 != "90210" {
		t.Error("Address is incorrect")
	}
}
