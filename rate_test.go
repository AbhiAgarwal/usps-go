package usps

import (
	"fmt"
	"os"
	"testing"
)

func TestRateDomestic(*testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")
	usps.Production = true

	var rate RateRequest
	rate.Service = "PRIORITY"
	rate.ZipOrigination = "44106"
	rate.ZipDestination = "20770"
	rate.Pounds = "1"
	rate.Ounces = "8"
	rate.Container = "NONRECTANGULAR"
	rate.Size = "LARGE"
	rate.Width = "15"
	rate.Length = "30"
	rate.Height = "15"
	rate.Girth = "55"

	output := usps.RateDomestic(rate)
	fmt.Println(output)

	usps.Production = false
	//if output.Error != "API Authorization failure. User "+usps.Username+" is not authorized to use API CarrierPickupAvailability." {
	//	t.Error("Pickup availability is incorrect.")
	//}
}
