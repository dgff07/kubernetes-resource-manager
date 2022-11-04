package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgff07/kubernetes-resource-manager/model"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type namespaceCreator struct {
	kubeapi *kubernetes.Clientset
}

func (ns *namespaceCreator) apply(jsonData string) error {

	namespace, err := convertJsonToNamespaceStruct(jsonData)

	if err != nil {
		return err
	}

	fmt.Printf("\nCreating namespace: %s \nannotations: %s\nlabels: %s\n\n", namespace.Name, namespace.Annotations, namespace.Labels)

	namespaceSpec := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace.Name}}

	_, err = ns.kubeapi.CoreV1().Namespaces().Create(context.Background(), namespaceSpec, metav1.CreateOptions{})

	if err != nil {
		fmt.Printf("An error occurred on trying to create the namespace '%s'", namespace.Name)
		return err
	}

	return nil
}

func convertJsonToNamespaceStruct(jsonData string) (*model.Namespace, error) {
	var namespace model.Namespace
	err := json.Unmarshal([]byte(jsonData), &namespace)

	if err != nil {
		fmt.Println("An error occurred on trying to convert the json string to the namespace struct.", err)
		return nil, err
	}

	return &namespace, nil
}
