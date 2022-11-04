package model

import (
	"encoding/json"

	"github.com/dgff07/kubernetes-resource-manager/logging"
)

func ConvertJsonToNamespaceStruct(jsonData string) (*Namespace, error) {
	var namespace Namespace
	err := json.Unmarshal([]byte(jsonData), &namespace)

	if err != nil {
		logging.Log.Error("An error occurred on trying to convert the json string to the namespace struct. " + err.Error())
		return nil, err
	}

	return &namespace, nil
}
