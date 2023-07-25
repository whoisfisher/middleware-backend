package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/namespace"
	"github.com/toolkits/pkg/ginx"
)

type NamespaceController struct {
	Ctx              context.Context
	NamespaceService namespace.NamespaceService
}

func NewNamespaceController() *NamespaceController {
	return &NamespaceController{
		NamespaceService: namespace.NewNamespaceService(),
	}
}

var namespaceController NamespaceController

func init() {
	namespaceController = *NewNamespaceController()
}

// 添加namespace
// @Tags 添加namespace
// @Summary: 添加namespace
// @Description: 添加namespace
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.NamespaceEntity true "request"
// @Success 200 {object} entity.NamespaceEntity
// @Router /api/v1/namespace [post]
func Create(ctx *gin.Context) {
	var namespace entity.NamespaceEntity
	if err := ctx.ShouldBind(&namespace); err != nil {
		logger.Log.Errorf("NamespaceInfo bind failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	err := namespaceController.NamespaceService.CreateNamespace(namespace)
	if err != nil {
		logger.Log.Errorf("Create namespace failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}
