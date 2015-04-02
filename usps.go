package usps

const (
	devbase  string = "http://production.shippingapis.com/ShippingAPITest.dll?API="
	prodbase string = "http://production.shippingapis.com/ShippingAPI.dll?API="
)

type USPS struct {
	Username   string
	Password   string
	Production bool `default:"false"`
}
