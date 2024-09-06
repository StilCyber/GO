package sprintThreeTaskThreeSolution

import "strings"

func defangIPaddr(address string) string {
    addressIP := strings.Split(address, "")
	result := []string{}
	for i := 0; i < len(addressIP); i++ {
		if addressIP[i] == "." {
			result = append(result, "[.]")
		} else {
			result = append(result, addressIP[i])
		}
	}

	return strings.Join(result, "")
}