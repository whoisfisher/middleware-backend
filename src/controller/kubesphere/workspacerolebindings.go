package kubesphere

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/kubesphere"
	"github.com/toolkits/pkg/ginx"
)

type WorkspaceRoleBindingsController struct {
	Ctx                          context.Context
	WorkspaceRoleBindingsService kubesphere.WorkspaceRoleBindingsService
}

func NewWorkspaceRoleBindingsControllerr() *WorkspaceRoleBindingsController {
	return &WorkspaceRoleBindingsController{
		WorkspaceRoleBindingsService: kubesphere.NewWorkspaceRoleBindingsService(),
	}
}

var workspaceRoleBindingsController WorkspaceRoleBindingsController

func init() {
	workspaceRoleBindingsController = *NewWorkspaceRoleBindingsControllerr()
}

// 获取ks平台租户空间角色绑定关系
// @Tags 获取ks平台租户空间角色绑定关系
// @Summary: 获取ks平台租户空间角色绑定关系
// @Description: 获取ks平台租户空间角色绑定关系
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Param workspace_name query string
// @Param username query string
// @Success 200 {object}
// @Router /cluster/iam/workspacerolebinding/list [get]
func ListWorkspaceRoleBindings(ctx *gin.Context) {
	wrb := entity.WorkspaceRoleBindings{}
	wrb.ClusterId = ctx.Query("cluster_id")
	wrb.WorkspaceName = ctx.Query("workspace_name")
	//wrb.Username = ctx.Query("username")
	wrbList, err := workspaceRoleBindingsController.WorkspaceRoleBindingsService.GetWorkspaceRoleBindingsList(&wrb)
	if err != nil {
		logger.Log.Errorf("List workspacerolebinding list info failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(wrbList, nil)
}

// 获取ks平台租户空间用户权限
// @Tags 获取ks平台租户空间用户权限
// @Summary: 获取ks平台租户空间用户权限
// @Description: 获取ks平台租户空间用户权限
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Param workspace_name query string
// @Param username query string
// @Success 200 {object}
// @Router /cluster/iam/workspacerolebinding/permission [get]
func GetWorkspacePermisson(ctx *gin.Context) {
	wrb := entity.WorkspaceRoleBindings{}
	wrb.ClusterId = ctx.Query("cluster_id")
	wrb.WorkspaceName = ctx.Query("workspace_name")
	wrb.Username = ctx.Query("username")
	ok, err := workspaceRoleBindingsController.WorkspaceRoleBindingsService.GetWorkspaceUserPermisson(&wrb)
	if err != nil {
		logger.Log.Errorf("Get user permission failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(ok, nil)
}
