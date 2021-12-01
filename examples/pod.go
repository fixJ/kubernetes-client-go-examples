package examples

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

func CreatePod(namespace, name, image string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	p := coreV1.Pod{
		ObjectMeta: metaV1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app": "nginx",
				"env": "prod",
				"namespace": "dev",
			},

		},
		Spec: coreV1.PodSpec{
			Containers: []coreV1.Container{
				{
					Name: name,
					Image: image,
					Ports: []coreV1.ContainerPort{
						{
							ContainerPort: 80,
							Protocol: coreV1.Protocol("TCP"),
						},
					},
				},
			},
		},
	}
	_, err = client.CoreV1().Pods(namespace).Create(context.TODO(), &p, metaV1.CreateOptions{})
	if err != nil {
		log.Printf("create pod in %s namespace failed, err: %v", namespace, err)
		return err
	}
	log.Printf("create pod in %s namespace successfully", namespace)
	return nil
}
