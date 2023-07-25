package namespace

import (
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/cluster"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
)

type NamespaceService interface {
	CreateNamespace(namespaceEntity entity.NamespaceEntity) error
}

type namespaceService struct {
	clusterService cluster.ClusterService
}

func NewNamespaceService() NamespaceService {
	return &namespaceService{
		clusterService: cluster.NewClusterService(),
	}
}

func (ns *namespaceService) CreateNamespace(namespaceEntity entity.NamespaceEntity) error {
	cluster, err := ns.clusterService.GetByID(namespaceEntity.ClusterId)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by clusterid[%s]: %s", namespaceEntity.ClusterId, err.Error())
		return err
	}
	namespaceEntity.Cluster = cluster
	return kubernetes.CreateNamespace(&namespaceEntity)
}
