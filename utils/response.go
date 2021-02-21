package utils

import "encoding/json"

func getJsonResponse() ([]byte, error) {
	persona := make(map[string]string)
	persona["name"] = "Luis Fernando"
	persona["surname"] = "Raga Renteria"
	persona["age"] = "31"
	persona["stack"] = "Go, JS, HTML, CSS"

	return json.MarshalIndent(persona, "", "  ")
}
