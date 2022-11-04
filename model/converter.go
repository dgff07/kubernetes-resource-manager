package model

import (
	"encoding/json"
	"fmt"
)

func ConvertJsonToNamespaceStruct(jsonData string) (*Namespace, error) {
	var namespace Namespace
	err := json.Unmarshal([]byte(jsonData), &namespace)

	if err != nil {
		fmt.Println("An error occurred on trying to convert the json string to the namespace struct.", err)
		return nil, err
	}

	return &namespace, nil
}
