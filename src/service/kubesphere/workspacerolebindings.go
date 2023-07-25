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

type WorkspaceRoleBindingsService interface {
	GetWorkspaceRoleBindingsList(wrb *entity.WorkspaceRoleBindings) (*unstructured.UnstructuredList, error)
	GetWorkspaceUserPermisson(wrb *entity.WorkspaceRoleBindings) (bool, error)
}

type workspaceRoleBindingsService struct {
	clusterRepository         repository.ClusterRepository
	globalRoleBindingsService GlobalRoleBindingsService
}

func NewWorkspaceRoleBindingsService() WorkspaceRoleBindingsService {
	return &workspaceRoleBindingsService{
		clusterRepository:         repository.NewClusterRepository(),
		globalRoleBindingsService: NewGlobalRoleBindingsService(),
	}
}

func (wb *workspaceRoleBindingsService) GetWorkspaceRoleBindingsList(wrb *entity.WorkspaceRoleBindings) (*unstructured.UnstructuredList, error) {
	cluster, err := wb.clusterRepository.GetByID(wrb.ClusterId)
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
	gvr := kubernetes.GetKubesphereWorkspaceRoleBindingsGvr()
	wrlList, err := client.Resource(gvr).Namespace("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Log.Errorf("Failed to get kubesphere Workspace role bindings:%s", err.Error())
		return nil, err
	}
	return wrlList, nil
}

func (wb *workspaceRoleBindingsService) GetWorkspaceUserPermisson(wrb *entity.WorkspaceRoleBindings) (bool, error) {
	if wrb.Username == "admin" {
		return true, nil
	}

	grb := &entity.GlobalRoleBindings{}
	grb.ClusterId = wrb.ClusterId
	grb.Username = wrb.Username

	ok, err := wb.globalRoleBindingsService.GetGlobalUserPermisson(grb)
	if ok {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	wrlList, err := wb.GetWorkspaceRoleBindingsList(wrb)
	if err != nil {
		return false, err
	}
	for _, wrl := range wrlList.Items {
		if wrb.Username+"-"+wrb.WorkspaceName+"-"+constant.WORKSPACE_ADMIN == wrl.GetName() {
			return true, nil
		}
	}
	return false, nil
}
