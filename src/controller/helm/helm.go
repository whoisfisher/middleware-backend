package helm

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/helm"
	"github.com/toolkits/pkg/ginx"
)

type HelmController struct {
	Ctx         context.Context
	HelmService helm.HelmService
}

func NewHelmController() *HelmController {
	return &HelmController{
		HelmService: helm.NewHelmService(),
	}
}

var helmController HelmController

func init() {
	helmController = *NewHelmController()
}

// 添加仓库
// @Tags 添加仓库
// @Summary: 添加仓库
// @Description: 添加仓库
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.HelmRepositry true "request"
// @Success 200 {object} entity.HelmRepositry
// @Router /api/v1/helm/repository [post]
func Add(ctx *gin.Context) {
	var helmRepo entity.HelmRepositry
	if err := ctx.ShouldBind(&helmRepo); err != nil {
		logger.Log.Errorf("HelmRepository bind failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	helmRepo1, err := helmController.HelmService.AddHelmRepository(helmRepo)
	if err != nil {
		logger.Log.Errorf("HelmRepository add failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(helmRepo1, nil)
}

// 更新仓库
// @Tags 更新仓库
// @Summary: 更新仓库
// @Description: 更新仓库
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.HelmRepositry true "request"
// @Success 200 {object} entity.HelmRepositry
// @Router /api/v1/helm/repository [patch]
func Update(ctx *gin.Context) {
	var helmRepo entity.HelmRepositry
	if err := ctx.ShouldBind(&helmRepo); err != nil {
		logger.Log.Errorf("HelmRepository bind failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	helmRepo1, err := helmController.HelmService.UpdateHelmRepository(helmRepo)
	if err != nil {
		logger.Log.Errorf("HelmRepository update failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(helmRepo1, nil)
}

// 删除仓库
// @Tags 删除仓库
// @Summary: 删除仓库
// @Description: 删除仓库
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   name     query    string     true        "用户名"
// @Success 200
// @Router /api/v1/helm/repository [delete]
func Delete(ctx *gin.Context) {
	name := ctx.Query("name")
	err := helmController.HelmService.DeleteHelmRepository(name)
	if err != nil {
		logger.Log.Errorf("HelmRepository delete failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

// 添加仓库到集群
// @Tags 添加仓库到集群
// @Summary: 添加仓库到集群
// @Description: 添加仓库到集群
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.HelmInstance true "request"
// @Success 200 {object} entity.HelmInstance
// @Router /api/v1/cluster/helm/repository [post]
func AddToCluster(ctx *gin.Context) {
	var helmInstance entity.HelmInstance
	if err := ctx.ShouldBind(&helmInstance); err != nil {
		logger.Log.Errorf("HelmInstance bind failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	err := helmController.HelmService.AddOrUpdateChartRepo(helmInstance)
	if err != nil {
		logger.Log.Errorf("Failed to add helm repository to cluster: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

// 安装应用到集群
// @Tags 安装应用到集群
// @Summary: 安装应用到集群
// @Description: 安装应用到集群
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.HelmInstance true "request"
// @Success 200 {object} entity.HelmInstance
// @Router /api/v1/cluster/helm/application [post]
func InstallChart(ctx *gin.Context) {
	var helmInstance entity.HelmInstance
	if err := ctx.ShouldBind(&helmInstance); err != nil {
		logger.Log.Errorf("HelmInstance bind failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	release, err := helmController.HelmService.InstallOrUpgradeChart(helmInstance)
	if err != nil {
		logger.Log.Errorf("HelmInstance %s install failed: %s", helmInstance.ReleaseName, err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(release, nil)
}

// 列出集群已经安装的helm应用
// @Tags 列出集群已经安装的helm应用
// @Summary: 列出集群已经安装的helm应用
// @Description: 列出集群已经安装的helm应用
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.HelmInstance true "request"
// @Success 200 {object} entity.HelmInstance
// @Router /api/v1/cluster/helm/application [post]
func ListDeployChart(ctx *gin.Context) {
	var helmInstance entity.HelmInstance
	if err := ctx.ShouldBind(&helmInstance); err != nil {
		logger.Log.Errorf("HelmInstance bind failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	releases, err := helmController.HelmService.ListDeployedReleases(helmInstance)
	if err != nil {
		logger.Log.Errorf("List deploy helm release failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(releases, nil)
}

// 列出集群已经安装的helm应用
// @Tags 列出集群已经安装的helm应用
// @Summary: 列出集群已经安装的helm应用
// @Description: 列出集群已经安装的helm应用
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.HelmInstance true "request"
// @Success 200 {object} entity.HelmInstance
// @Router /api/v1/cluster/helm/application [post]
func UninstallChart(ctx *gin.Context) {
	var helmInstance entity.HelmInstance
	if err := ctx.ShouldBind(&helmInstance); err != nil {
		logger.Log.Errorf("HelmInstance bind failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	err := helmController.HelmService.UninstallRelease(helmInstance)
	if err != nil {
		logger.Log.Errorf("HelmInstance %s uninstall failed: %s", helmInstance.ReleaseName, err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}
