package main

import (
	"fmt"
)

func main() {
	data := createDummyData()

	for _, configData := range data {

		config, err := mapToConfiguration("DELETE_NAMESPACE", configData)

		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		err = config.applier.apply(config.jsonData)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}

func createDummyData() []string {
	return []string{"{\"name\": \"test_1\",\"annotations\": [\"\"],\"labels\": [\"\"]}"}
}
