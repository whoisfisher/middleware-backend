package entity

type NamespaceEntity struct {
	ClusterId     string  `json:"cluster_id"`
	Namespace     string  `json:"namespace"`
	Cluster       Cluster `json:"cluster"`
	Username      string  `json:"username"`
	WorkspaceName string  `json:"workspace_name"`
}

type GlobalRoleBindings struct {
	ClusterId string `json:"cluster_id"`
	Username  string `json:"username"`
}

type WorkspaceRoleBindings struct {
	ClusterId     string `json:"cluster_id"`
	WorkspaceName string `json:"workspace_name"`
	Username      string `json:"username"`
}

type RoleBindings struct {
	ClusterId     string `json:"cluster_id"`
	WorkspaceName string `json:"workspace_name"`
	Namespace     string `json:"namespace"`
	Username      string `json:"username"`
}
