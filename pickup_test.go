package usps

import (
	"os"
	"testing"
)

func TestPickupAvailability(t *testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")

	var pickup PickUpRequest
	pickup.FirmName = "ABC Corp."
	pickup.SuiteOrApt = "Suite 777"
	pickup.Address2 = "1390 Market Street"
	pickup.Urbanization = ""
	pickup.City = "Houston"
	pickup.State = "TX"
	pickup.ZIP5 = "77058"
	pickup.ZIP4 = "1234"

	output := usps.PickupAvailability(pickup)
	if output.Error != "API Authorization failure. User "+usps.Username+" is not authorized to use API CarrierPickupAvailability." {
		t.Error("Pickup availability is incorrect.")
	}
}

func TestPickupChange(t *testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")

	var pickup PickupChangeRequest
	pickup.FirstName = "John"
	pickup.LastName = "Doe"
	pickup.FirmName = ""
	pickup.SuiteOrApt = ""
	pickup.Address2 = "1390 Market Street"
	pickup.Urbanization = ""
	pickup.City = "Houston"
	pickup.State = "HX"
	pickup.ZIP5 = ""
	pickup.ZIP4 = ""
	pickup.Phone = "(555) 555-1234"
	pickup.Extension = ""
	pickup.Package.ServiceType = "PriorityMail"
	pickup.Package.Count = "1"
	pickup.EstimatedWeight = "14"
	pickup.PackageLocation = "Front Door"
	pickup.SpecialInstructions = ""
	pickup.ConfirmationNumber = "WTC123456789"

	output := usps.PickupChange(pickup)
	if output.Error != "API Authorization failure. User "+usps.Username+" is not authorized to use API CarrierPickupChange." {
		t.Error("Pickup change is incorrect.")
	}
}

func TestPickupInquiry(t *testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")

	var pickup PickUpInquiryRequest
	pickup.FirmName = ""
	pickup.SuiteOrApt = ""
	pickup.Address2 = "1390 Market Street"
	pickup.Urbanization = ""
	pickup.City = ""
	pickup.State = ""
	pickup.ZIP5 = "77058"
	pickup.ZIP4 = ""
	pickup.ConfirmationNumber = "WTC123456789"

	output := usps.PickupInquiry(pickup)
	if output.Error != "API Authorization failure. User "+usps.Username+" is not authorized to use API CarrierPickupInquiry." {
		t.Error("Pickup change is incorrect.")
	}
}
