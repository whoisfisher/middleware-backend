package kubesphere

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type GlobalRoleBindingsService interface {
	GetGlobalRoleBindingsList(grb *entity.GlobalRoleBindings) (*unstructured.UnstructuredList, error)
	GetGlobalUserPermisson(grb *entity.GlobalRoleBindings) (bool, error)
}

type globalRoleBindingsService struct {
	clusterRepository repository.ClusterRepository
}

func NewGlobalRoleBindingsService() GlobalRoleBindingsService {
	return &globalRoleBindingsService{
		clusterRepository: repository.NewClusterRepository(),
	}
}

func (gb *globalRoleBindingsService) GetGlobalRoleBindingsList(grb *entity.GlobalRoleBindings) (*unstructured.UnstructuredList, error) {
	cluster, err := gb.clusterRepository.GetByID(grb.ClusterId)
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
	gvr := kubernetes.GetKubesphereGlobalRoleBindingsGvr()
	grlList, err := client.Resource(gvr).Namespace("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Log.Errorf("Failed to get kubesphere global role bindings:%s", err.Error())
		return nil, err
	}
	return grlList, nil
}

func (gb *globalRoleBindingsService) GetGlobalUserPermisson(grb *entity.GlobalRoleBindings) (bool, error) {
	if grb.Username == "admin" {
		return true, nil
	}
	grlList, err := gb.GetGlobalRoleBindingsList(grb)
	if err != nil {
		return false, err
	}
	for _, grl := range grlList.Items {
		if grb.Username+"-"+constant.PLATFORM_ADMIN == grl.GetName() {
			return true, nil
		}
	}
	return false, nil
}
