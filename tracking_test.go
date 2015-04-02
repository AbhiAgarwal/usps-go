package usps

import (
	"os"
	"testing"
)

func TestTrackPackage(t *testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")

	output := usps.TrackPackage("9341989949036022338924")
	if output.TrackInfo.TrackSummary != "The Postal Service could not locate the tracking information for your request. Please verify your tracking number and try again later." {
		t.Error("Tracker is incorrect")
	}
}
