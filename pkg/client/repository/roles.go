package repository

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
)

type RolesRepository interface {
	BootstrapRolesForDB(request request.RolesRequest)(response response.RolesResponse) 
	GetRolesForPublic(request request.RolesRequest)(response response.RolesResponse) 
	GetRolesForInternal(request request.RolesRequest)(response response.RolesResponse) 
	GetRolesForPrivate(request request.RolesRequest)(response response.RolesResponse) 
	CreateRolesForPublic(request request.RolesRequest)   (response response.RolesResponse) 
	CreateRolesForInternal(request request.RolesRequest) (response response.RolesResponse)
	CreateRolesForPrivate(request request.RolesRequest)  (response response.RolesResponse)
	UpdateRolesForPublic(request request.RolesRequest)   (response response.RolesResponse)
	UpdateRolesForInternal(request request.RolesRequest) (response response.RolesResponse)
	UpdateRolesForPrivate(request request.RolesRequest)  (response response.RolesResponse)
	DeleteRolesForPublic(request request.RolesRequest)   (response response.RolesResponse)
	DeleteRolesForInternal(request request.RolesRequest) (response response.RolesResponse)
	DeleteRolesForPrivate(request request.RolesRequest)  (response response.RolesResponse)
}

type rolesRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (rolesRepository rolesRepository) BootstrapRolesForDB(request request.RolesRequest) (response response.RolesResponse) {
	fmt.Println("BootstrapRolesForDB")
	return response
}

//GET
func (rolesRepository rolesRepository) GetRolesForPublic(request request.RolesRequest) (response response.RolesResponse) {
	fmt.Println("GetRolesForPublic")
	return response
}

func (rolesRepository rolesRepository) GetRolesForInternal(request request.RolesRequest) (response response.RolesResponse ){
	fmt.Println("GetRolesForInternal")
	return response
}

func (rolesRepository rolesRepository) GetRolesForPrivate(request request.RolesRequest) (response response.RolesResponse) {
	fmt.Println("GetRolesForPrivate")
	return response
}

//CREATE
func (rolesRepository rolesRepository) CreateRolesForPublic(request request.RolesRequest) (response response.RolesResponse ){
	fmt.Println("CreateRolesForPublic")
	return response
}

func (rolesRepository rolesRepository) CreateRolesForInternal(request request.RolesRequest) (response response.RolesResponse) {
	fmt.Println("CreateRolesForInternal()")
	return response
}

func (rolesRepository rolesRepository) CreateRolesForPrivate(request request.RolesRequest) (response response.RolesResponse){
	fmt.Println("CreateRolesForPrivate()")
	return response
}

//UPDATE
func (rolesRepository rolesRepository) UpdateRolesForPublic(request request.RolesRequest) (response response.RolesResponse){
	fmt.Println("UpdateRolesForPublic()")
	return response
}

func (rolesRepository rolesRepository) UpdateRolesForInternal(request request.RolesRequest) (response response.RolesResponse){
	fmt.Println("UpdateRolesForInternal")
	return response
}

func (rolesRepository rolesRepository) UpdateRolesForPrivate(request request.RolesRequest) (response response.RolesResponse){
	fmt.Println("UpdateRolesForPrivate")
	return response
}

//DELETE
func (rolesRepository rolesRepository) DeleteRolesForPublic(request request.RolesRequest) (response response.RolesResponse){
	fmt.Println("DeleteRolesForPublic")
	return response
}

func (rolesRepository rolesRepository) DeleteRolesForInternal(request request.RolesRequest) (response response.RolesResponse ){
	fmt.Println("DeleteRolesForInternal")
	return response
}

func (rolesRepository rolesRepository) DeleteRolesForPrivate(request request.RolesRequest) (response response.RolesResponse){
	fmt.Println("DeleteRolesForPrivate")
	return response
}

func NewRolesRepository(conf config.BaseConfig) RolesRepository {
	return &rolesRepository{BaseConfig: conf}
}