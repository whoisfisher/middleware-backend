package kubernetes

import (
	"fmt"
	mclient "github.com/minio/operator/pkg/client/clientset/versioned"
	helm "github.com/mittwald/go-helm-client"
	"github.com/pkg/errors"
	pclient "github.com/zalando/postgres-operator/pkg/generated/clientset/versioned"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type Config struct {
	ApiServer  string
	Token      string
	KubeConfig string
}

func GetKubernetesRestConfig(c *Config) *rest.Config {
	kubeConf := &rest.Config{
		Host:        c.ApiServer,
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	return kubeConf
}

func GetKubernetesRestConfigFromKubeConfig(c *Config) (*rest.Config, error) {
	kubeConf, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.KubeConfig))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("new kubernetes client with config failed: %v", err))
	}
	return kubeConf, nil
}

func NewKubernetesClient(c *Config) (*kubernetes.Clientset, error) {
	kubeConf, err := GetKubernetesRestConfigFromKubeConfig(c)
	if err != nil {
		kubeConf = GetKubernetesRestConfig(c)
	}
	client, err := kubernetes.NewForConfig(kubeConf)
	if err != nil {
		return client, errors.Wrap(err, fmt.Sprintf("new kubernetes client with config failed: %v", err))
	}
	return client, nil
}

func NewKubernetesDynamicClient(c *Config) (dynamic.Interface, error) {
	kubeConf, err := GetKubernetesRestConfigFromKubeConfig(c)
	if err != nil {
		kubeConf = GetKubernetesRestConfig(c)
	}
	dynamicClient, err1 := dynamic.NewForConfig(kubeConf)
	if err1 != nil {
		return nil, errors.Wrap(err1, fmt.Sprintf("new kubernetes dynamic client with config failed: %v", err1))
	}
	return dynamicClient, nil
}

func NewKubernetesDiscoveryClient(c *Config) (*discovery.DiscoveryClient, error) {
	client, err := NewKubernetesClient(c)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("new kubernetes dynamic client with config failed: %v", err))
	}
	discoveryClient := discovery.NewDiscoveryClient(client.RESTClient())
	return discoveryClient, nil
}

func NewHelmClient(c *Config) (helm.Client, error) {
	options := &helm.KubeConfClientOptions{
		Options: &helm.Options{
			Namespace:        "default", // Change this to the namespace you wish to install the chart in.
			RepositoryCache:  "/tmp/.helmcache",
			RepositoryConfig: "/tmp/.helmrepo",
			Debug:            true,
			Linting:          true, // Change this to false if you don't want linting.
			DebugLog: func(format string, v ...interface{}) {
				log.Printf(format, v...)
			},
		},
		KubeContext: "",
		KubeConfig:  []byte(c.KubeConfig),
	}
	client, err := helm.NewClientFromKubeConf(options, helm.Burst(100), helm.Timeout(10e9))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewPostgresqlClient(c *Config) (*pclient.Clientset, error) {
	kubeConf, err := GetKubernetesRestConfigFromKubeConfig(c)
	if err != nil {
		kubeConf = GetKubernetesRestConfig(c)
	}
	client, err := pclient.NewForConfig(kubeConf)
	if err != nil {
		return client, errors.Wrap(err, fmt.Sprintf("new postgresql client with config failed: %v", err))
	}
	return client, nil
}

func NewMinioClient(c *Config) (*mclient.Clientset, error) {
	kubeConf, err := GetKubernetesRestConfigFromKubeConfig(c)
	if err != nil {
		kubeConf = GetKubernetesRestConfig(c)
	}
	client, err := mclient.NewForConfig(kubeConf)
	if err != nil {
		return client, errors.Wrap(err, fmt.Sprintf("new postgresql client with config failed: %v", err))
	}
	return client, nil
}
