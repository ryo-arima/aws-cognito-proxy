package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type RoleControllerForPublic interface {
	GetRoles(c *gin.Context)
}

type roleControllerForPublic struct {
	RoleRepository repository.RoleRepository
}

func (roleController roleControllerForPublic) GetRoles(c *gin.Context) {
	var roleRequest request.RoleRequest
	if err := c.Bind(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RoleResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Roles: []response.Role{}})
		return
	}
	res := roleController.RoleRepository.GetRoles()
	c.JSON(http.StatusOK, res)
	return
}


func NewRoleControllerForPublic(roleRepository repository.RoleRepository) RoleControllerForPublic {
	return &roleControllerForPublic{ RoleRepository: roleRepository}
}