package main

import (
	"fmt"
)

func main() {
	data := createDummyData()

	for _, configData := range data {

		config, err := mapToConfiguration("CREATE_NAMESPACE", configData)

		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		config.applier.apply(config.jsonData)
	}

}

func createDummyData() []string {
	return []string{"{\"name\": \"test_1\",\"annotations\": [\"\"],\"labels\": [\"\"]}"}
}
