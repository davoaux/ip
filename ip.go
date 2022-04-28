package ip

import (
	"strconv"
	"strings"
)

func IsValid(address string) bool {
	octets := strings.Split(address, ".")
	if len(octets) != 4 {
		return false
	}
	for _, octet := range octets {
		if octet[0] == '0' && len(octet) > 1 {
			return false
		}
		n, err := strconv.Atoi(octet)
		if err != nil || n < 0 || n > 255 {
			return false
		}
	}
	return true
}
