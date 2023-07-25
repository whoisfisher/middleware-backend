package kubernetes

import (
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GetGVR(gvr entity.GroupVersionKind) schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    gvr.Group,
		Version:  gvr.Version,
		Resource: gvr.Resource,
	}
}

func GetKubesphereIamUserGvr() schema.GroupVersionResource {
	gvr := entity.GroupVersionKind{
		Group:    "iam.kubesphere.io",
		Version:  "v1alpha2",
		Resource: "users",
	}
	return GetGVR(gvr)
}

func GetKubesphereTenantWorkspaceGvr() schema.GroupVersionResource {
	gvr := entity.GroupVersionKind{
		Group:    "tenant.kubesphere.io",
		Version:  "v1alpha1",
		Resource: "workspaces",
	}
	return GetGVR(gvr)
}

func GetKubesphereGlobalRoleBindingsGvr() schema.GroupVersionResource {
	gvr := entity.GroupVersionKind{
		Group:    "iam.kubesphere.io",
		Version:  "v1alpha2",
		Resource: "globalrolebindings",
	}
	return GetGVR(gvr)
}

func GetKubesphereWorkspaceRoleBindingsGvr() schema.GroupVersionResource {
	gvr := entity.GroupVersionKind{
		Group:    "iam.kubesphere.io",
		Version:  "v1alpha2",
		Resource: "workspacerolebindings",
	}
	return GetGVR(gvr)
}

func GetKubesphereRoleBindingsGvr() schema.GroupVersionResource {
	gvr := entity.GroupVersionKind{
		Group:    "rbac.authorization.k8s.io",
		Version:  "v1",
		Resource: "rolebindings",
	}
	return GetGVR(gvr)
}
