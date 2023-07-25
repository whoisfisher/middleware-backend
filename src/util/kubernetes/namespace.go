package kubernetes

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func CreateNamespace(namespaceEntity *entity.NamespaceEntity) error {
	config := &Config{
		ApiServer:  namespaceEntity.Cluster.ApiServer,
		Token:      namespaceEntity.Cluster.Token,
		KubeConfig: namespaceEntity.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return err
	}
	namespace := corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Namespace",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceEntity.Namespace,
		},
	}

	_, err = client.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func GetNamespaces(namespaceEntity *entity.NamespaceEntity) ([]corev1.Namespace, error) {
	config := &Config{
		ApiServer:  namespaceEntity.Cluster.ApiServer,
		Token:      namespaceEntity.Cluster.Token,
		KubeConfig: namespaceEntity.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return nil, err
	}
	options := metav1.ListOptions{}
	if namespaceEntity.WorkspaceName != "" {
		var labelSelector labels.Selector
		selector := &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"kubesphere.io/workspace": namespaceEntity.WorkspaceName,
			},
		}
		labelSelector, err = metav1.LabelSelectorAsSelector(selector)
		if err != nil {
			return nil, err
		}
		options.LabelSelector = labelSelector.String()
	}
	namespaceList, err := client.CoreV1().Namespaces().List(context.TODO(), options)
	if err != nil {
		return nil, err
	}
	namespaces := []corev1.Namespace{}
	if namespaceEntity.Username != "" {
		for _, namespace := range namespaceList.Items {
			keyAnno := "kubesphere.io/creator"
			for key, value := range namespace.Annotations {
				if key == keyAnno && value == namespaceEntity.Username {
					namespaces = append(namespaces, namespace)
					break
				}
			}
		}
	} else {
		namespaces = namespaceList.Items
	}
	return namespaces, nil
}
