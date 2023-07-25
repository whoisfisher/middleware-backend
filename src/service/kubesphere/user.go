package kubesphere

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type UserService interface {
	GetUserList(clusterId string) (*unstructured.UnstructuredList, error)
}

type userService struct {
	clusterRepository repository.ClusterRepository
}

func NewUserService() UserService {
	return &userService{
		clusterRepository: repository.NewClusterRepository(),
	}
}

func (user *userService) GetUserList(clusterId string) (*unstructured.UnstructuredList, error) {
	cluster, err := user.clusterRepository.GetByID(clusterId)
	if err != nil {
		logger.Log.Errorf("Failed to get cluster info:%s", err.Error())
		return nil, err
	}
	config := kubernetes.Config{
		ApiServer:  cluster.ApiServer,
		Token:      cluster.Token,
		KubeConfig: cluster.KubeConfig,
	}
	client, err := kubernetes.NewKubernetesDynamicClient(&config)
	if err != nil {
		logger.Log.Errorf("Failed to create kubernetes dynamic client:%s", err.Error())
		return nil, err
	}
	gvr := kubernetes.GetKubesphereIamUserGvr()
	userList, err := client.Resource(gvr).Namespace("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Log.Errorf("Failed to get kubesphere user:%s", err.Error())
		return nil, err
	}
	return userList, nil
}
