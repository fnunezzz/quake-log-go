package helpers

import (
	"encoding/json"
	"fmt"
)

func ToJson(obj interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return jsonData, nil
}