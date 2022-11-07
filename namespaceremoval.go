package main

import (
	"context"

	"github.com/dgff07/go-log-lib/logging"
	"github.com/dgff07/kubernetes-resource-manager/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type namespaceRemoval struct {
	kubeapi *kubernetes.Clientset
}

func (ns *namespaceRemoval) apply(jsonData string) error {
	namespace, err := model.ConvertJsonToNamespaceStruct(jsonData)

	if err != nil {
		return err
	}

	logging.Log.Info("Deleting namespace " + namespace.Name)

	err = ns.kubeapi.CoreV1().Namespaces().Delete(context.Background(), namespace.Name, metav1.DeleteOptions{})

	if err != nil {
		logging.Log.Error("An error occurred on trying to delete the namespace:" + namespace.Name + ". " + err.Error())
		return err
	}

	return nil
}
