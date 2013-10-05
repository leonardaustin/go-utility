package stringsextra

import(
	"strings"
)

func ContainsAny(haystack string, needles []string, caseSensitive bool) bool {

	if(!caseSensitive){
		haystack = strings.ToLower(haystack)
	}

	for _, needle := range needles {

		if(strings.Contains(haystack, string(needle))){
			return true
		}

	}

	return false
}