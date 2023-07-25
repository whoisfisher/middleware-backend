package helm

import (
	"context"
	"encoding/json"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
	helm "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
)

type HelmService interface {
	AddHelmRepository(helmRepo entity.HelmRepositry) (*entity.HelmRepositry, error)
	DeleteHelmRepository(name string) error
	UpdateHelmRepository(helmRepo entity.HelmRepositry) (*entity.HelmRepositry, error)
	AddOrUpdateChartRepo(instance entity.HelmInstance) error
	InstallOrUpgradeChart(instance entity.HelmInstance) (*release.Release, error)
	ListDeployedReleases(instance entity.HelmInstance) ([]*release.Release, error)
	UninstallRelease(instance entity.HelmInstance) error
}

type helmService struct {
	helmRepository    repository.HelmRepository
	clusterRepository repository.ClusterRepository
}

func NewHelmService() HelmService {
	return &helmService{
		helmRepository:    repository.NewHelmRepository(),
		clusterRepository: repository.NewClusterRepository(),
	}
}

func (hs *helmService) AddHelmRepository(helmRepo entity.HelmRepositry) (*entity.HelmRepositry, error) {
	err := hs.helmRepository.Save(&helmRepo.HelmRepository)
	if err != nil {
		logger.Log.Errorf("Faile to save helm repository info: %s", err.Error())
		return nil, err
	}

	return &helmRepo, nil
}

func (hs *helmService) DeleteHelmRepository(name string) error {
	err := hs.helmRepository.Delete(name)
	if err != nil {
		logger.Log.Errorf("Faile to delete helm repository info: %s", err.Error())
		return err
	}
	return nil
}

func (hs *helmService) UpdateHelmRepository(helmRepo entity.HelmRepositry) (*entity.HelmRepositry, error) {
	mapRepository := make(map[string]interface{})
	jsonRepository, err := json.Marshal(helmRepo)
	if err != nil {
		logger.Log.Errorf("Faile to marshal helm repository info: %s", err.Error())
		return nil, err
	}
	err = json.Unmarshal(jsonRepository, &mapRepository)
	if err != nil {
		logger.Log.Errorf("Faile to unmarshal helm repository info: %s", err.Error())
		return nil, err
	}
	err = hs.helmRepository.Update(helmRepo.ID, mapRepository)
	if err != nil {
		logger.Log.Errorf("Faile to update helm repository info: %s", err.Error())
		return nil, err
	}
	var helmRepositoryEntity *entity.HelmRepositry
	helmRepositoryEntity.HelmRepository, err = hs.helmRepository.GetByID(helmRepo.ID)
	if err != nil {
		logger.Log.Errorf("Faile to get helm repository info by id[%s]: %s", helmRepo.ID, err.Error())
		return nil, err
	}
	return helmRepositoryEntity, nil
}

func (hs *helmService) AddOrUpdateChartRepo(instance entity.HelmInstance) error {
	helmRepo, err := hs.helmRepository.GetByID(instance.RepoId)
	if err != nil {
		logger.Log.Errorf("Faile to get helm repository info by id[%s]: %s", instance.RepoId, err.Error())
		return err
	}
	cluster, err := hs.clusterRepository.GetByID(instance.ClusterId)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", instance.ClusterId, err.Error())
		return err
	}
	helmClient, err := kubernetes.GetHelmClient(cluster.KubeConfig)
	if err != nil {
		logger.Log.Errorf("Faile to helm client: %s", err.Error())
		return err
	}
	var repoEntry repo.Entry
	repoEntry.Name = helmRepo.Name
	repoEntry.URL = helmRepo.Url
	repoEntry.Username = helmRepo.Username
	repoEntry.Password = helmRepo.Password
	repoEntry.CertFile = helmRepo.CertFile
	repoEntry.CAFile = helmRepo.CAFile
	repoEntry.KeyFile = helmRepo.KeyFile
	repoEntry.InsecureSkipTLSverify = helmRepo.InsecureSkipTlsVerify
	if err := helmClient.AddOrUpdateChartRepo(repoEntry); err != nil {
		logger.Log.Errorf("Faile to add or update helm repository to cluster: %s", err.Error())
		return err
	}
	return nil
}

func (hs *helmService) InstallOrUpgradeChart(instance entity.HelmInstance) (*release.Release, error) {
	cluster, err := hs.clusterRepository.GetByID(instance.ClusterId)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", instance.ClusterId, err.Error())
		return nil, err
	}
	helmClient, err := kubernetes.GetHelmClient(cluster.KubeConfig)
	if err != nil {
		logger.Log.Errorf("Faile to helm client: %s", err.Error())
		return nil, err
	}
	chartSpec := helm.ChartSpec{
		ReleaseName:     instance.ReleaseName,
		ChartName:       instance.ChartName,
		Namespace:       instance.Namespace,
		ValuesYaml:      instance.ValuesYaml,
		CreateNamespace: instance.CreateNamespace,
	}
	release1, err := helmClient.InstallOrUpgradeChart(context.TODO(), &chartSpec, &helm.GenericHelmOptions{})
	if err != nil {
		logger.Log.Errorf("Faile to install or update chart release: %s", err.Error())
		return nil, err
	}
	return release1, nil
}

func (hs *helmService) ListDeployedReleases(instance entity.HelmInstance) ([]*release.Release, error) {
	cluster, err := hs.clusterRepository.GetByID(instance.ClusterId)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", instance.ClusterId, err.Error())
		return nil, err
	}
	helmClient, err := kubernetes.GetHelmClient(cluster.KubeConfig)
	if err != nil {
		logger.Log.Errorf("Faile to helm client: %s", err.Error())
		return nil, err
	}
	releases, err := helmClient.ListDeployedReleases()
	if err != nil {
		logger.Log.Errorf("Faile to list deployed helm chart release: %s", err.Error())
		return nil, err
	}
	return releases, nil
}

func (hs *helmService) UninstallRelease(instance entity.HelmInstance) error {
	cluster, err := hs.clusterRepository.GetByID(instance.ClusterId)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", instance.ClusterId, err.Error())
		return err
	}
	helmClient, err := kubernetes.GetHelmClient(cluster.KubeConfig)
	if err != nil {
		logger.Log.Errorf("Faile to helm client: %s", err.Error())
		return err
	}
	chartSpec := helm.ChartSpec{
		ReleaseName:     instance.ReleaseName,
		ChartName:       instance.ChartName,
		Namespace:       instance.Namespace,
		ValuesYaml:      instance.ValuesYaml,
		CreateNamespace: instance.CreateNamespace,
	}
	err = helmClient.UninstallRelease(&chartSpec)
	if err != nil {
		logger.Log.Errorf("Faile to uninstall deployed helm chart release[%s]: %s", instance.ReleaseName, err.Error())
		return err
	}
	return nil
}
