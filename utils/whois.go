package utils

import (
	"github.com/likexian/whois"

	"strings"
)

//
// GetCountry (ipaddr string) (country string)
//

func GetCountry(ipaddr string) string {
	result, err := whois.Whois(ipaddr)
	if err != nil {
		return "XX"
	}

	lines := strings.Split(result, "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), "country") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	// XX would be unknown
	return "XX"
}
