package kubesphere

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/kubesphere"
	"github.com/toolkits/pkg/ginx"
)

type WorkspaceController struct {
	Ctx              context.Context
	WorkspaceService kubesphere.WorkspaceService
}

func NewWorkspaceController() *WorkspaceController {
	return &WorkspaceController{
		WorkspaceService: kubesphere.NewWorkspaceService(),
	}
}

var workspaceController WorkspaceController

func init() {
	workspaceController = *NewWorkspaceController()
}

// 根据集群ID获取租户空间列表
// @Tags 根据集群ID获取租户空间列表
// @Summary: 根据集群ID获取租户空间列表
// @Description: 根据集群ID获取租户空间列表
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Success 200 {object}
// @Router /cluster/tenant/workspace/list [get]
func ListWorkspaces(ctx *gin.Context) {
	ClusterId := ctx.Query("cluster_id")
	workspaces, err := workspaceController.WorkspaceService.GetWorkspaceList(ClusterId)
	if err != nil {
		logger.Log.Errorf("List User info failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(workspaces, nil)
}
