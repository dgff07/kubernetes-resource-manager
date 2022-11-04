package main

import (
	"github.com/dgff07/kubernetes-resource-manager/logging"
)

func main() {
	logging.Init()
	data := createDummyData()

	for _, configData := range data {

		config, err := mapToConfiguration("DELETE_NAMESPACE", configData)

		if err != nil {
			logging.Log.Error(err.Error())
			continue
		}

		err = config.applier.apply(config.jsonData)
		if err != nil {
			logging.Log.Error(err.Error())
		}
	}

}

func createDummyData() []string {
	return []string{"{\"name\": \"test_1\",\"annotations\": [\"\"],\"labels\": [\"\"]}"}
}
