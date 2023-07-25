package kubesphere

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/kubesphere"
	"github.com/toolkits/pkg/ginx"
)

type UserController struct {
	Ctx         context.Context
	UserService kubesphere.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: kubesphere.NewUserService(),
	}
}

var userController UserController

func init() {
	userController = *NewUserController()
}

// 根据集群ID获取用户列表
// @Tags 根据集群ID获取用户列表
// @Summary: 根据集群ID获取用户列表
// @Description: 根据集群ID获取用户列表
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param cluster_id query string
// @Success 200 {object}
// @Router /cluster/iam/user/list [get]
func ListUsers(ctx *gin.Context) {
	ClusterId := ctx.Query("cluster_id")
	users, err := userController.UserService.GetUserList(ClusterId)
	if err != nil {
		logger.Log.Errorf("List User info failed: %s", err.Error())
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(users, nil)
}
