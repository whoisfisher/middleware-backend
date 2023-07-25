package kubesphere

import (
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
	corev1 "k8s.io/api/core/v1"
)

type ProjectService interface {
	GetProjectList(namespaceEntity *entity.NamespaceEntity) ([]corev1.Namespace, error)
}

type projectService struct {
	clusterRepository repository.ClusterRepository
}

func NewProjectService() ProjectService {
	return &projectService{
		clusterRepository: repository.NewClusterRepository(),
	}
}

func (ps *projectService) GetProjectList(namespaceEntity *entity.NamespaceEntity) ([]corev1.Namespace, error) {
	cluster, err := ps.clusterRepository.GetByID(namespaceEntity.ClusterId)
	if err != nil {
		logger.Log.Errorf("Failed to get cluster info:%s", err.Error())
		return nil, err
	}
	namespaceEntity.Cluster.Cluster = cluster
	projectList, err := kubernetes.GetNamespaces(namespaceEntity)
	if err != nil {
		logger.Log.Errorf("Failed to get kubesphere project:%s", err.Error())
		return nil, err
	}
	return projectList, nil
}
