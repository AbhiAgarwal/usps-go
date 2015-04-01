package usps

import "net/url"

func URLEncode(urlToEncode string) string {
	return url.QueryEscape(urlToEncode)
}
