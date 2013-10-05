package stringsextra

import(
	"strings"
)

func IndexAny(haystack string, needles []string, caseSensitive bool) int {

	if(!caseSensitive){
		haystack = strings.ToLower(haystack)
	}

	index := 0
	for _, needle := range needles {

		if(strings.Contains(haystack, string(needle))){
			return index
		}

		index++
	}

	return -1
}