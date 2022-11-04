package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type namespaceRemoval struct {
	kubeapi *kubernetes.Clientset
}

func (ns *namespaceRemoval) apply(jsonData string) error {
	namespace, err := convertJsonToNamespaceStruct(jsonData)

	if err != nil {
		return err
	}

	fmt.Printf("\nDeleting namespace: %s\n\n", namespace.Name)

	err = ns.kubeapi.CoreV1().Namespaces().Delete(context.Background(), namespace.Name, metav1.DeleteOptions{})

	if err != nil {
		fmt.Printf("An error occurred on trying to delete the namespace '%s'\n", namespace.Name)
		return err
	}

	return nil
}
