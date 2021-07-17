package externals

import "strings"

func SliceAContainsAnyStringInSliceB(sliceA []string, B string) []string {
	result := []string{}
	for _, crypto := range sliceA {
		if strings.Contains(B, crypto) {
			result = append(result, crypto)
		}
	}
	return result
}
