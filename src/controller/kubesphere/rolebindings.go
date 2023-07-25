package kubesphere

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/kubesphere"
	"github.com/toolkits/pkg/ginx"
)

type RoleBindingsController struct {
	Ctx                 context.Context
	RoleBindingsService kubesphere.RoleBindingsService
}

func NewRoleBindingsControllerr() *RoleBindingsController {
	return &RoleBindingsController{
		RoleBindingsService: kubesphere.NewRoleBindingsService(),
	}
}

var roleBindingsController RoleBindingsController

func init() {
	roleBindingsController = *NewRoleBindingsControllerr()
}

// 获取ks平台租户空间中项目角色绑定关系
// @Tags 获取ks平台租户空间中项目角色绑定关系
// @Summary: 获取ks平台租户空间中项目角色绑定关系
// @Description: 获取ks平台租户空间中项目角色绑定关系
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Param workspace_name query string
// @Param namespace query string
// @Param username query string
// @Success 200 {object}
// @Router /cluster/iam/rolebinding/list [get]
func ListRoleBindings(ctx *gin.Context) {
	rb := entity.RoleBindings{}
	rb.ClusterId = ctx.Query("cluster_id")
	rb.WorkspaceName = ctx.Query("workspace_name")
	rb.Namespace = ctx.Query("namespace")
	//rb.Username = ctx.Query("username")
	rbList, err := roleBindingsController.RoleBindingsService.GetRoleBindingsList(&rb)
	if err != nil {
		logger.Log.Errorf("List rolebinding list info failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(rbList, nil)
}

// 获取ks平台租户空间用户权限
// @Tags 获取ks平台租户空间用户权限
// @Summary: 获取ks平台租户空间用户权限
// @Description: 获取ks平台租户空间用户权限
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Param workspace_name query string
// @Param namespace query string
// @Param namespace query string
// @Param username query string
// @Success 200 {object}
// @Router /cluster/iam/rolebinding/permission [get]
func GetPermisson(ctx *gin.Context) {
	rb := entity.RoleBindings{}
	rb.ClusterId = ctx.Query("cluster_id")
	rb.WorkspaceName = ctx.Query("workspace_name")
	rb.Namespace = ctx.Query("namespace")
	rb.Username = ctx.Query("username")
	ok, err := roleBindingsController.RoleBindingsService.GetUserPermisson(&rb)
	if err != nil {
		logger.Log.Errorf("Get user permission failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(ok, nil)
}
