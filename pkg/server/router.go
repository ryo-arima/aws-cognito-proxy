package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/controller"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/middleware"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

func InitRouter(conf config.BaseConfig) *gin.Engine {
	
	usersRepository := repository.NewUsersRepository(conf)
	usersControllerForPublic := controller.NewUsersControllerForPublic(usersRepository)
	usersControllerForInternal := controller.NewUsersControllerForInternal(usersRepository)
	usersControllerForPrivate := controller.NewUsersControllerForPrivate(usersRepository)
	
	groupsRepository := repository.NewGroupsRepository(conf)
	groupsControllerForPublic := controller.NewGroupsControllerForPublic(groupsRepository)
	groupsControllerForInternal := controller.NewGroupsControllerForInternal(groupsRepository)
	groupsControllerForPrivate := controller.NewGroupsControllerForPrivate(groupsRepository)
	
	rolesRepository := repository.NewRolesRepository(conf)
	rolesControllerForPublic := controller.NewRolesControllerForPublic(rolesRepository)
	rolesControllerForInternal := controller.NewRolesControllerForInternal(rolesRepository)
	rolesControllerForPrivate := controller.NewRolesControllerForPrivate(rolesRepository)
	
	policiesRepository := repository.NewPoliciesRepository(conf)
	policiesControllerForPublic := controller.NewPoliciesControllerForPublic(policiesRepository)
	policiesControllerForInternal := controller.NewPoliciesControllerForInternal(policiesRepository)
	policiesControllerForPrivate := controller.NewPoliciesControllerForPrivate(policiesRepository)
	
	
    router := gin.Default()
	privateAPI := router.Group("api/private")
	internalAPI := router.Group("api/internal")
	publicAPI := router.Group("api/public")

	publicAPI.Use(middleware.ForPublic(conf))
	internalAPI.Use(middleware.ForInternal(conf))
	privateAPI.Use(middleware.ForPrivate(conf))

	
	//users
	publicAPI.GET("/userss", usersControllerForPublic.GetUserss)
	internalAPI.GET("/userss", usersControllerForInternal.GetUserss)
	internalAPI.POST("/users", usersControllerForInternal.CreateUsers)
	internalAPI.PUT("/users", usersControllerForInternal.UpdateUsers)
	internalAPI.DELETE("/users", usersControllerForInternal.DeleteUsers)
	privateAPI.GET("/userss", usersControllerForPrivate.GetUserss)
	privateAPI.POST("/users", usersControllerForPrivate.CreateUsers)
	privateAPI.PUT("/users", usersControllerForPrivate.UpdateUsers)
	privateAPI.DELETE("/users", usersControllerForPrivate.DeleteUsers)
	
	//groups
	publicAPI.GET("/groupss", groupsControllerForPublic.GetGroupss)
	internalAPI.GET("/groupss", groupsControllerForInternal.GetGroupss)
	internalAPI.POST("/groups", groupsControllerForInternal.CreateGroups)
	internalAPI.PUT("/groups", groupsControllerForInternal.UpdateGroups)
	internalAPI.DELETE("/groups", groupsControllerForInternal.DeleteGroups)
	privateAPI.GET("/groupss", groupsControllerForPrivate.GetGroupss)
	privateAPI.POST("/groups", groupsControllerForPrivate.CreateGroups)
	privateAPI.PUT("/groups", groupsControllerForPrivate.UpdateGroups)
	privateAPI.DELETE("/groups", groupsControllerForPrivate.DeleteGroups)
	
	//roles
	publicAPI.GET("/roless", rolesControllerForPublic.GetRoless)
	internalAPI.GET("/roless", rolesControllerForInternal.GetRoless)
	internalAPI.POST("/roles", rolesControllerForInternal.CreateRoles)
	internalAPI.PUT("/roles", rolesControllerForInternal.UpdateRoles)
	internalAPI.DELETE("/roles", rolesControllerForInternal.DeleteRoles)
	privateAPI.GET("/roless", rolesControllerForPrivate.GetRoless)
	privateAPI.POST("/roles", rolesControllerForPrivate.CreateRoles)
	privateAPI.PUT("/roles", rolesControllerForPrivate.UpdateRoles)
	privateAPI.DELETE("/roles", rolesControllerForPrivate.DeleteRoles)
	
	//policies
	publicAPI.GET("/policiess", policiesControllerForPublic.GetPoliciess)
	internalAPI.GET("/policiess", policiesControllerForInternal.GetPoliciess)
	internalAPI.POST("/policies", policiesControllerForInternal.CreatePolicies)
	internalAPI.PUT("/policies", policiesControllerForInternal.UpdatePolicies)
	internalAPI.DELETE("/policies", policiesControllerForInternal.DeletePolicies)
	privateAPI.GET("/policiess", policiesControllerForPrivate.GetPoliciess)
	privateAPI.POST("/policies", policiesControllerForPrivate.CreatePolicies)
	privateAPI.PUT("/policies", policiesControllerForPrivate.UpdatePolicies)
	privateAPI.DELETE("/policies", policiesControllerForPrivate.DeletePolicies)
	

	return router
}