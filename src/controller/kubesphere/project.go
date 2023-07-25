package kubesphere

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/kubesphere"
	"github.com/toolkits/pkg/ginx"
)

type ProjectController struct {
	Ctx            context.Context
	ProjectService kubesphere.ProjectService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		ProjectService: kubesphere.NewProjectService(),
	}
}

var projectController ProjectController

func init() {
	projectController = *NewProjectController()
}

// 获取ks项目列表
// @Tags 获取ks项目列表
// @Summary: 获取ks项目列表
// @Description: 获取ks项目列表
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Param workspace_name query string
// @Param username query string
// @Success 200 {object}
// @Router /cluster/workspace/project/list [get]
func ListProjects(ctx *gin.Context) {
	nsEntity := entity.NamespaceEntity{}
	nsEntity.ClusterId = ctx.Query("cluster_id")
	nsEntity.WorkspaceName = ctx.Query("workspace_name")
	nsEntity.Username = ctx.Query("username")
	projects, err := projectController.ProjectService.GetProjectList(&nsEntity)
	if err != nil {
		logger.Log.Errorf("List User info failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(projects, nil)
}
