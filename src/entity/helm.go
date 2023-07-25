package entity

import "github.com/mensylisir/kmpp-middleware/src/model"

type HelmRepositry struct {
	model.HelmRepository
}

type HelmInstance struct {
	RepoId          string `json:"repo_id"`
	ClusterId       string `json:"cluster_id"`
	Namespace       string `json:"namespace"`
	ReleaseName     string `json:"release_name"`
	ChartName       string `json:"chart_name"`
	ValuesYaml      string `json:"values_yaml"`
	CreateNamespace bool   `json:"create_namespace"`
}
