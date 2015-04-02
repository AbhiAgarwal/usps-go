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

Save time and money by reducing shipping errors due to improper address entry. This tool corrects errors in street addresses, including abbreviations and missing information. It also supplies a ZIP+4 Code.

```go
var usps USPS
usps.Username = ""

var address Address
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

var address Address
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

var address ZipCode
address.Zip5 = "90210"

output := usps.CityStateLookup(address)
```

### Tracking & Delivery Information APIs

The descriptions are from [here](https://www.usps.com/business/web-tools-apis/delivery-information.htm).

The access level for this API doesn't require a password, but only requires a username to be set.

#### Track & Confirm

Track any package shipped via Priority Mail Express, Global Express Guaranteed, or Priority Mail Express International services. Check the tracking information for packages shipped with USPS Tracking, Signature Confirmation, Certified Mail, or Registered Mail™ services.

```go
var usps USPS
usps.Username = ""

output := usps.TrackPackage("")
```
