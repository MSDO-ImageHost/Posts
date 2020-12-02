package utils

import (
	"encoding/json"
	"log"
)

func PrettyFormatMap(d interface{}) string {
	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
