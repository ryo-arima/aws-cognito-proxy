package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/controller"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/middleware"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

func InitRouter(conf config.BaseConfig) *gin.Engine {
	
	userRepository := repository.NewUserRepository(conf)
	userControllerForPublic := controller.NewUserControllerForPublic(userRepository)
	userControllerForInternal := controller.NewUserControllerForInternal(userRepository)
	userControllerForPrivate := controller.NewUserControllerForPrivate(userRepository)
	
	groupRepository := repository.NewGroupRepository(conf)
	groupControllerForPublic := controller.NewGroupControllerForPublic(groupRepository)
	groupControllerForInternal := controller.NewGroupControllerForInternal(groupRepository)
	groupControllerForPrivate := controller.NewGroupControllerForPrivate(groupRepository)
	
	roleRepository := repository.NewRoleRepository(conf)
	roleControllerForPublic := controller.NewRoleControllerForPublic(roleRepository)
	roleControllerForInternal := controller.NewRoleControllerForInternal(roleRepository)
	roleControllerForPrivate := controller.NewRoleControllerForPrivate(roleRepository)
	
	policyRepository := repository.NewPolicyRepository(conf)
	policyControllerForPublic := controller.NewPolicyControllerForPublic(policyRepository)
	policyControllerForInternal := controller.NewPolicyControllerForInternal(policyRepository)
	policyControllerForPrivate := controller.NewPolicyControllerForPrivate(policyRepository)
	
	
    router := gin.Default()
	privateAPI := router.Group("api/private")
	internalAPI := router.Group("api/internal")
	publicAPI := router.Group("api/public")

	publicAPI.Use(middleware.ForPublic(conf))
	internalAPI.Use(middleware.ForInternal(conf))
	privateAPI.Use(middleware.ForPrivate(conf))

	
	//user
	publicAPI.GET("/users", userControllerForPublic.GetUsers)
	internalAPI.GET("/users", userControllerForInternal.GetUsers)
	internalAPI.POST("/user", userControllerForInternal.CreateUser)
	internalAPI.PUT("/user", userControllerForInternal.UpdateUser)
	internalAPI.DELETE("/user", userControllerForInternal.DeleteUser)
	privateAPI.GET("/users", userControllerForPrivate.GetUsers)
	privateAPI.POST("/user", userControllerForPrivate.CreateUser)
	privateAPI.PUT("/user", userControllerForPrivate.UpdateUser)
	privateAPI.DELETE("/user", userControllerForPrivate.DeleteUser)
	
	//group
	publicAPI.GET("/groups", groupControllerForPublic.GetGroups)
	internalAPI.GET("/groups", groupControllerForInternal.GetGroups)
	internalAPI.POST("/group", groupControllerForInternal.CreateGroup)
	internalAPI.PUT("/group", groupControllerForInternal.UpdateGroup)
	internalAPI.DELETE("/group", groupControllerForInternal.DeleteGroup)
	privateAPI.GET("/groups", groupControllerForPrivate.GetGroups)
	privateAPI.POST("/group", groupControllerForPrivate.CreateGroup)
	privateAPI.PUT("/group", groupControllerForPrivate.UpdateGroup)
	privateAPI.DELETE("/group", groupControllerForPrivate.DeleteGroup)
	
	//role
	publicAPI.GET("/roles", roleControllerForPublic.GetRoles)
	internalAPI.GET("/roles", roleControllerForInternal.GetRoles)
	internalAPI.POST("/role", roleControllerForInternal.CreateRole)
	internalAPI.PUT("/role", roleControllerForInternal.UpdateRole)
	internalAPI.DELETE("/role", roleControllerForInternal.DeleteRole)
	privateAPI.GET("/roles", roleControllerForPrivate.GetRoles)
	privateAPI.POST("/role", roleControllerForPrivate.CreateRole)
	privateAPI.PUT("/role", roleControllerForPrivate.UpdateRole)
	privateAPI.DELETE("/role", roleControllerForPrivate.DeleteRole)
	
	//policy
	publicAPI.GET("/policys", policyControllerForPublic.GetPolicys)
	internalAPI.GET("/policys", policyControllerForInternal.GetPolicys)
	internalAPI.POST("/policy", policyControllerForInternal.CreatePolicy)
	internalAPI.PUT("/policy", policyControllerForInternal.UpdatePolicy)
	internalAPI.DELETE("/policy", policyControllerForInternal.DeletePolicy)
	privateAPI.GET("/policys", policyControllerForPrivate.GetPolicys)
	privateAPI.POST("/policy", policyControllerForPrivate.CreatePolicy)
	privateAPI.PUT("/policy", policyControllerForPrivate.UpdatePolicy)
	privateAPI.DELETE("/policy", policyControllerForPrivate.DeletePolicy)
	

	return router
}