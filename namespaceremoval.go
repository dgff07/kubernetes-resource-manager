package main

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
)

type namespaceRemoval struct {
	kubeapi *kubernetes.Clientset
}

func (*namespaceRemoval) apply(jsonData string) error {
	fmt.Println("TODO, remove namespace. data: ", jsonData)
	return nil
}
