package jsonextra

var Seprator = "."

func flatten(inputJSON map[string]interface{}, lkey string, flattened map[string]interface{}) {
	for rkey, value := range inputJSON {
		key := lkey + rkey
		if _, ok := value.(map[string]interface{}); ok {
			flatten(value.(map[string]interface{}), key+Seprator, flattened)
		} else {
			flattened[key] = value
		}
	}
}
