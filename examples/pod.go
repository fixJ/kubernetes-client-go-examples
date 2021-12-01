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

func DeletePod(name, namespace string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	err = client.CoreV1().Pods("dev").Delete(context.TODO(), name, metaV1.DeleteOptions{})
	if err != nil {
		log.Printf("delete %s pod in %s namespace failed, err: %v", name, namespace, err)
		return err
	}
	log.Printf("delete %s pod in %s namespace successfully", name, namespace)
	return nil
}

func GetPod(name, namespace string) (*coreV1.Pod, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	get, err := client.CoreV1().Pods(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		log.Printf("get %s pod in %s namespace failed, err: %v", name, namespace, err)
		return nil, err
	}
	log.Printf("get %s pod in %s namespace successfully, pod: %v", name, namespace, get.String())
	return get, nil
}

func UpdatePod(namespace string, p *coreV1.Pod) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	_, err = client.CoreV1().Pods(namespace).Update(context.TODO(), p, metaV1.UpdateOptions{})
	if err != nil {
		log.Printf("update pod in %s namespace failed, err: %v", namespace, err)
		return err
	}
	log.Printf("update pod in %s namespace successfully", namespace)
	return nil
}