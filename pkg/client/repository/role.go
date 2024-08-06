package repository

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
)

type RoleRepository interface {
	BootstrapRoleForDB(request request.RoleRequest)(response response.RoleResponse) 
	GetRoleForPublic(request request.RoleRequest)(response response.RoleResponse) 
	GetRoleForInternal(request request.RoleRequest)(response response.RoleResponse) 
	GetRoleForPrivate(request request.RoleRequest)(response response.RoleResponse) 
	CreateRoleForPublic(request request.RoleRequest)   (response response.RoleResponse) 
	CreateRoleForInternal(request request.RoleRequest) (response response.RoleResponse)
	CreateRoleForPrivate(request request.RoleRequest)  (response response.RoleResponse)
	UpdateRoleForPublic(request request.RoleRequest)   (response response.RoleResponse)
	UpdateRoleForInternal(request request.RoleRequest) (response response.RoleResponse)
	UpdateRoleForPrivate(request request.RoleRequest)  (response response.RoleResponse)
	DeleteRoleForPublic(request request.RoleRequest)   (response response.RoleResponse)
	DeleteRoleForInternal(request request.RoleRequest) (response response.RoleResponse)
	DeleteRoleForPrivate(request request.RoleRequest)  (response response.RoleResponse)
}

type roleRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (roleRepository roleRepository) BootstrapRoleForDB(request request.RoleRequest) (response response.RoleResponse) {
	fmt.Println("BootstrapRoleForDB")
	return response
}

//GET
func (roleRepository roleRepository) GetRoleForPublic(request request.RoleRequest) (response response.RoleResponse) {
	fmt.Println("GetRoleForPublic")
	return response
}

func (roleRepository roleRepository) GetRoleForInternal(request request.RoleRequest) (response response.RoleResponse ){
	fmt.Println("GetRoleForInternal")
	return response
}

func (roleRepository roleRepository) GetRoleForPrivate(request request.RoleRequest) (response response.RoleResponse) {
	fmt.Println("GetRoleForPrivate")
	return response
}

//CREATE
func (roleRepository roleRepository) CreateRoleForPublic(request request.RoleRequest) (response response.RoleResponse ){
	fmt.Println("CreateRoleForPublic")
	return response
}

func (roleRepository roleRepository) CreateRoleForInternal(request request.RoleRequest) (response response.RoleResponse) {
	fmt.Println("CreateRoleForInternal()")
	return response
}

func (roleRepository roleRepository) CreateRoleForPrivate(request request.RoleRequest) (response response.RoleResponse){
	fmt.Println("CreateRoleForPrivate()")
	return response
}

//UPDATE
func (roleRepository roleRepository) UpdateRoleForPublic(request request.RoleRequest) (response response.RoleResponse){
	fmt.Println("UpdateRoleForPublic()")
	return response
}

func (roleRepository roleRepository) UpdateRoleForInternal(request request.RoleRequest) (response response.RoleResponse){
	fmt.Println("UpdateRoleForInternal")
	return response
}

func (roleRepository roleRepository) UpdateRoleForPrivate(request request.RoleRequest) (response response.RoleResponse){
	fmt.Println("UpdateRoleForPrivate")
	return response
}

//DELETE
func (roleRepository roleRepository) DeleteRoleForPublic(request request.RoleRequest) (response response.RoleResponse){
	fmt.Println("DeleteRoleForPublic")
	return response
}

func (roleRepository roleRepository) DeleteRoleForInternal(request request.RoleRequest) (response response.RoleResponse ){
	fmt.Println("DeleteRoleForInternal")
	return response
}

func (roleRepository roleRepository) DeleteRoleForPrivate(request request.RoleRequest) (response response.RoleResponse){
	fmt.Println("DeleteRoleForPrivate")
	return response
}

func NewRoleRepository(conf config.BaseConfig) RoleRepository {
	return &roleRepository{BaseConfig: conf}
}