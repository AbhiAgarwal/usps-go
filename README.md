# usps-go

[![GoDoc](https://godoc.org/github.com/aureum/usps-go?status.svg)](https://godoc.org/github.com/aureum/usps-go)
[![Circle CI](https://circleci.com/gh/aureum/usps-go/tree/master.svg?style=svg)](https://circleci.com/gh/aureum/usps-go/tree/master)

A wrapper ontop of the USPS API.

Download it using

``go get github.com/aureum/usps-go``

Import it into your project

```go
package main

import "github.com/aureum/usps-go"
```

Authentication onto USPS API is usin the username/password pair that you get in the e-mail from USPS when you sign up for their API. 

```go
var usps USPS
usps.Username = ""
usps.Password = ""
```

### Address Information APIs

The descriptions are from [here](https://www.usps.com/business/web-tools-apis/address-information.htm).

The access level for this API doesn't require a password, but only requires a username to be set.

#### Address Standardization/Verification

Save time and money by reducing shipping errors due to improper address entry. This tool corrects errors in street addresses, including abbreviations and missing information. It also supplies a ZIP+4® Code.

```go
var usps USPS
usps.Username = ""

var address USPS.Address
address.Address2 = "6406 Ivy Lane"
address.City = "Greenbelt"
address.State = "MD"

output := usps.AddressVerification(address)
```

#### Zip Code Lookup

Never get another ZIP Code wrong again with the ZIP Code Lookup tool. Find matching ZIP Codes or ZIP+4 Codes for any given address, city, and state in the U.S.

```go
var usps USPS
usps.Username = ""

var address USPS.Address
address.Address2 = "6406 Ivy Lane"
address.City = "Greenbelt"
address.State = "MD"

output := usps.ZipCodeLookup(address)
```

#### City/State Lookup

Don’t waste time searching for a city or state; use a ZIP Code to get accurate city and state information automatically.

```go
var usps USPS
usps.Username = ""

var address USPS.ZipCode
address.Zip5 = "90210"

output := usps.CityStateLookup(address)
```

#### Hold For Pickup Facility Information

Need pickup times and destination information? This tool provides all USPS® facilities that offer Hold For Pickup service within the destination area.

### Tracking & Delivery Information APIs

The descriptions are from [here](https://www.usps.com/business/web-tools-apis/delivery-information.htm).

The access level for this API doesn't require a password, but only requires a username to be set.

#### Track & Confirm

Track any package shipped via Priority Mail Express™, Global Express Guaranteed®, or Priority Mail Express International™ services. Check the tracking information for packages shipped with USPS Tracking™, Signature Confirmation™, Certified Mail®, or Registered Mail™ services.

```go
var usps USPS
usps.Username = ""

output := usps.TrackPackage("")
```

### Package Pickup

Stop going to the Post Office™ to drop off your packages. With the Package Pickup API, you can submit a pickup request and we will pick up your packages for free.

The access level for this API doesn't require a password, but only requires a username to be set.

#### Package Pickup

Check the availability for Package Pickup at a specific address, Schedule a Pickup, and get first available date for pickup. Package Pickup is not available for all addresses. Available for Priority Mail Express™, Priority Mail®, International, and returns packages. Package Pickup is not available for all addresses.

##### Package Pickup Availability

```go
var usps USPS
usps.Username = ""

var pickup USPS.PickUpRequest
pickup.FirmName = "ABC Corp."
pickup.SuiteOrApt = "Suite 777"
pickup.Address2 = "1390 Market Street"
pickup.Urbanization = ""
pickup.City = "Houston"
pickup.State = "TX"
pickup.ZIP5 = "77058"
pickup.ZIP4 = "1234"

output := usps.PickupAvailability(pickup)
```

##### Package Pickup Change

```go
var usps USPS
usps.Username = ""

var pickup USPS.PickupChangeRequest
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
```

##### Package Pickup Inquiry

```go
var usps USPS
usps.Username = ""

var pickup USPS.PickUpInquiryRequest
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
```

## Coming soon

### Price Calculator

Calculate postage rates quickly and easily online for domestic and international shipping.

The access level for this API doesn't require a password, but only requires a username to be set.

#### Domestic Price Calculator

Calculate how much domestic shipping will cost with Priority Mail Express™, Priority Mail®, First-Class Mail™, Standard Post™, Media Mail®, and Library Mail services.

#### International Price Calculator

Get prices for Global Express Guaranteed®, Priority Mail Express International™, Priority Mail International®, First-Class Package International Service™, and First-Class Mail International® services. Unique mailing restrictions are provided for each country, along with declarations form information and Priority Mail Express International delivery areas.

### Print Shipping Labels

Print a complete shipping label with a tracking barcode to track packages seamlessly with one of our PC Postage® partners.

#### USPS Tracking™ Labels

Generate USPS Tracking barcoded labels for Priority Mail®, First-Class Mail® parcels, and package services parcels, including Standard Post™, Media Mail®, and Library Mail. Optional features include a post-date request and e-mail ship notification to recipient.

#### Signature Confirmation™ Labels

Need proof of delivery? Generate a Signature Confirmation barcoded label for Priority Mail, First-Class Mail parcels, Standard Post, Media Mail, and Library Mail services, and we’ll provide the complete address label, including the Signature Confirmation Service barcode.

#### Priority Mail Open & Distribute® Labels

This tool generates Priority Mail Open & Distribute labels to be placed on a Tag 161 or Tag 190 for Priority Mail Open & Distribute containers. Priority Mail Open & Distribute expedites the transportation of bulk mailings by using Priority Mail service to quickly send the mailings to a destination delivery unit for processing.

#### Priority Mail Express® Labels

Generate a single-ply Priority Mail Express shipping label complete with return and delivery addresses, a barcode, and a mailing record for your use.

#### Electronic Merchandise Return Service Labels

With just one click, online shoppers can print a return label right from your web site. This cuts down on customer service calls—saving time and money. Print return labels for Priority Mail, First-Class Mail, Standard Post, Media Mail, and Library Mail services. Additional information, such as insurance or an authorization number, can be included, too.

#### International Shipping Labels

Send documents and packages globally. USPS® offers reliable, affordable shipping to more than 180 countries. Generate Priority Mail Express International™, Priority Mail International®, First-Class Mail International®, or First-Class Package International Service shipping labels complete with addresses, barcode, customs form, and mailing receipt.

#### Customs Forms

International mail is subject to customs examination in the destination country. Contents and value of an item must be declared on the applicable customs form, PS Form 2976 (Customs Declaration) or PS Form 2976-A (Customs Declaration and Dispatch Note). Generate the customs forms complete with addresses, barcodes, and customs forms.

### Service Standards & Commitments

Get estimates on delivery and receive guaranteed Commitments with certain services.

#### Priority Mail® Service Standards

Request service standards for Priority Mail service between any two 5-digit ZIP Codes™ with an average delivery standard of 2 or 3 days.

#### First-Class Mail® Service Standards

Request service standards for First-Class Mail service between any two 5-digit ZIP Codes™ with an average delivery standard of 2 or 3 days.

#### Priority Mail Express® Service Commitments

Receive our guaranteed commitment between any two 5-digit ZIP Code™ for Priority Mail Express delivery. Choose from up to 200 drop-off locations and drop-off times. Plus, get guaranteed delivery times for upcoming mailings.

#### Service Delivery Calculator

Get estimates on delivery standards between 3-digit ZIP Codes™ for Standard Post™, Library Mail, Media Mail®, and Bound Printed Matter services.

#### Package Services Service Standards

