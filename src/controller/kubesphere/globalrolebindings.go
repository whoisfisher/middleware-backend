package kubesphere

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/kubesphere"
	"github.com/toolkits/pkg/ginx"
)

type GlobalRoleBindingsController struct {
	Ctx                       context.Context
	GlobalRoleBindingsService kubesphere.GlobalRoleBindingsService
}

func NewGlobalRoleBindingsControllerr() *GlobalRoleBindingsController {
	return &GlobalRoleBindingsController{
		GlobalRoleBindingsService: kubesphere.NewGlobalRoleBindingsService(),
	}
}

var globalRoleBindingsController GlobalRoleBindingsController

func init() {
	globalRoleBindingsController = *NewGlobalRoleBindingsControllerr()
}

// 获取ks平台角色绑定关系
// @Tags 获取ks平台角色绑定关系
// @Summary: 获取ks平台角色绑定关系
// @Description: 获取ks平台角色绑定关系
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Param username query string
// @Success 200 {object}
// @Router /cluster/iam/globalrolebinding/list [get]
func ListGlobalRoleBindings(ctx *gin.Context) {
	grb := entity.GlobalRoleBindings{}
	grb.ClusterId = ctx.Query("cluster_id")
	//grb.Username = ctx.Query("username")
	grbList, err := globalRoleBindingsController.GlobalRoleBindingsService.GetGlobalRoleBindingsList(&grb)
	if err != nil {
		logger.Log.Errorf("List globalrolebinding list info failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(grbList, nil)
}

// 获取ks平台用户权限
// @Tags 获取ks平台用户权限
// @Summary: 获取ks平台用户权限
// @Description: 获取ks平台用户权限
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Param workspace_name query string
// @Param username query string
// @Success 200 {object}
// @Router /cluster/iam/globalrolebinding/permission [get]
func GetGlobalPermisson(ctx *gin.Context) {
	grb := entity.GlobalRoleBindings{}
	grb.ClusterId = ctx.Query("cluster_id")
	grb.Username = ctx.Query("username")
	ok, err := globalRoleBindingsController.GlobalRoleBindingsService.GetGlobalUserPermisson(&grb)
	if err != nil {
		logger.Log.Errorf("Get user permission failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(ok, nil)
}
