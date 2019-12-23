package util

import ( 
	"strings"
) 

//Upper returns string uppercased
func Upper(input string) string {
	return strings.ToUpper(input)
}

//Lower returns string lowercased
func Lower(input string) string {
	return strings.ToLower(input)
}
