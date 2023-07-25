package kubernetes

import (
	helm "github.com/mittwald/go-helm-client"
)

func GetHelmClient(kubeconfig string) (helm.Client, error) {
	config := &Config{
		KubeConfig: kubeconfig,
	}
	return NewHelmClient(config)
}
