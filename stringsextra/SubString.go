package stringsextra

import(
	"strings"
	"regexp"
)

func SubStringSearch(content []byte, startKeyword string, endKeyword string, removeWhiteSpace bool) string {

	// Remvoe whitespace
	if(removeWhiteSpace){
		re := regexp.MustCompile(`\s+`)
		content = re.ReplaceAll(content, []byte(""))
	}

	// Find start
	stringStart := strings.Index(string(content), startKeyword) + len(startKeyword)

	// Take substring with ~100 characters
	subString := string(content[stringStart:])

	// Find potition of first (next) comma + get latitide
	endPosition := strings.Index(subString, endKeyword)

	// Whats between them!
	result := string(subString[0:endPosition])

	return result
}
