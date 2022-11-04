package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgff07/kubernetes-resource-manager/kube"
)

const (
	CREATE_NAMESPACE string = "CREATE_NAMESPACE"
	DELETE_NAMESPACE string = "DELETE_NAMESPACE"
)

var configurationTypesMap map[string]configurationApplier = nil

func defineConfigurationTypeMap() {

	kubeapi, err := kube.BuildClient()
	if err != nil {
		fmt.Println("Error:", err)
	}

	configurationTypesMap = map[string]configurationApplier{
		CREATE_NAMESPACE: &namespaceCreator{kubeapi: kubeapi},
		DELETE_NAMESPACE: &namespaceRemoval{kubeapi: kubeapi},
	}
}

type configurationApplier interface {
	apply(jsonData string) error
}

type configuration struct {
	applier  configurationApplier
	jsonData string
}

func mapToConfiguration(configType string, jsonData string) (*configuration, error) {

	if len(strings.TrimSpace(configType)) == 0 {
		return nil, errors.New("you must pass a not empty configuration type")
	}

	if len(strings.TrimSpace(jsonData)) == 0 {
		return nil, errors.New("you must pass a not empty json data")
	}

	if configurationTypesMap == nil {
		defineConfigurationTypeMap()
	}

	applier := configurationTypesMap[configType]

	if applier == nil {
		return nil, errors.New("There is not configuration type with the designation: " + configType)
	}

	return &configuration{applier: applier, jsonData: jsonData}, nil
}
