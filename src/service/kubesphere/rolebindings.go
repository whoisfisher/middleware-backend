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

type RoleBindingsService interface {
	GetRoleBindingsList(rb *entity.RoleBindings) (*unstructured.UnstructuredList, error)
	GetUserPermisson(rb *entity.RoleBindings) (bool, error)
}

type roleBindingsService struct {
	clusterRepository            repository.ClusterRepository
	globalRoleBindingsService    GlobalRoleBindingsService
	workspaceRoleBindingsService WorkspaceRoleBindingsService
}

func NewRoleBindingsService() RoleBindingsService {
	return &roleBindingsService{
		clusterRepository:            repository.NewClusterRepository(),
		globalRoleBindingsService:    NewGlobalRoleBindingsService(),
		workspaceRoleBindingsService: NewWorkspaceRoleBindingsService(),
	}
}

func (rbs *roleBindingsService) GetRoleBindingsList(rb *entity.RoleBindings) (*unstructured.UnstructuredList, error) {
	cluster, err := rbs.clusterRepository.GetByID(rb.ClusterId)
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
	gvr := kubernetes.GetKubesphereRoleBindingsGvr()
	rlList, err := client.Resource(gvr).Namespace(rb.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Log.Errorf("Failed to get kubesphere role bindings:%s", err.Error())
		return nil, err
	}
	return rlList, nil
}

func (rbs *roleBindingsService) GetUserPermisson(rb *entity.RoleBindings) (bool, error) {
	if rb.Username == "admin" {
		return true, nil
	}
	grb := &entity.GlobalRoleBindings{}
	grb.ClusterId = rb.ClusterId
	grb.Username = rb.Username

	ok, err := rbs.globalRoleBindingsService.GetGlobalUserPermisson(grb)
	if ok {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	wrb := &entity.WorkspaceRoleBindings{}
	wrb.ClusterId = rb.ClusterId
	wrb.WorkspaceName = rb.WorkspaceName
	wrb.Username = rb.Username

	ok1, err1 := rbs.workspaceRoleBindingsService.GetWorkspaceUserPermisson(wrb)
	if ok1 {
		return true, nil
	}
	if err1 != nil {
		return false, err
	}

	rbList, err := rbs.GetRoleBindingsList(rb)
	if err != nil {
		return false, err
	}
	for _, wrl := range rbList.Items {
		if wrb.Username+"-"+constant.PROJECT_ADMIN == wrl.GetName() {
			return true, nil
		}
	}
	return false, nil
}
