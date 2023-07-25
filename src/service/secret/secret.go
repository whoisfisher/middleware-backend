package secret

import (
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/cluster"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
)

type SecretService interface {
	GetSecrets(instance entity.Instance) ([]entity.SecretInfo, error)
}

type secretService struct {
	clusterService cluster.ClusterService
}

func NewSecretService() SecretService {
	return &secretService{
		clusterService: cluster.NewClusterService(),
	}
}

func (s secretService) GetSecrets(instance entity.Instance) ([]entity.SecretInfo, error) {
	clusterObj, err := s.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by clusterid[%s]: %s", instance.ClusterID, err.Error())
		return nil, err
	}
	instance.Cluster = clusterObj.Cluster
	secretInfos, err := kubernetes.GetSecretInfo(&instance)
	if err != nil {
		logger.Log.Errorf("Faile to get secret info from cluster: %s", err.Error())
		return nil, err
	}
	return secretInfos, nil
}
