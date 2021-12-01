package examples

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

func GetNodes(name string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	get, err := client.CoreV1().Nodes().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		log.Printf("get node: %s failed, %v", name, err)
		return err
	}
	log.Printf("get node: %s successfully, node: %s", name, get)
	return nil
}
