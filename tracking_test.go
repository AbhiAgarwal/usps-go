package usps

import "testing"

func TestTrackPackage(t *testing.T) {
	var usps USPS
	usps.Username = ""

	output := usps.TrackPackage("")
	if output.TrackInfo.TrackSummary != "The Postal Service could not locate the tracking information for your request. Please verify your tracking number and try again later." {
		t.Error("Tracker is incorrect")
	}
}
