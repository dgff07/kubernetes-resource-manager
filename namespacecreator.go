package main

import (
	"context"

	"github.com/dgff07/go-log-lib/logging"
	"github.com/dgff07/kubernetes-resource-manager/model"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type namespaceCreator struct {
	kubeapi *kubernetes.Clientset
}

func (ns *namespaceCreator) apply(jsonData string) error {

	namespace, err := model.ConvertJsonToNamespaceStruct(jsonData)

	if err != nil {
		return err
	}

	logging.Log.Info("Creating namespace " + namespace.Name)

	namespaceSpec := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace.Name}}

	_, err = ns.kubeapi.CoreV1().Namespaces().Create(context.Background(), namespaceSpec, metav1.CreateOptions{})

	if err != nil {
		logging.Log.Error("An error occurred on trying to create the namespace:" + namespace.Name + ". " + err.Error())
		return err
	}

	return nil
}
