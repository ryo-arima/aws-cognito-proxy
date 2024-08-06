package repository

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
)

type UsersRepository interface {
	BootstrapUsersForDB(request request.UsersRequest)(response response.UsersResponse) 
	GetUsersForPublic(request request.UsersRequest)(response response.UsersResponse) 
	GetUsersForInternal(request request.UsersRequest)(response response.UsersResponse) 
	GetUsersForPrivate(request request.UsersRequest)(response response.UsersResponse) 
	CreateUsersForPublic(request request.UsersRequest)   (response response.UsersResponse) 
	CreateUsersForInternal(request request.UsersRequest) (response response.UsersResponse)
	CreateUsersForPrivate(request request.UsersRequest)  (response response.UsersResponse)
	UpdateUsersForPublic(request request.UsersRequest)   (response response.UsersResponse)
	UpdateUsersForInternal(request request.UsersRequest) (response response.UsersResponse)
	UpdateUsersForPrivate(request request.UsersRequest)  (response response.UsersResponse)
	DeleteUsersForPublic(request request.UsersRequest)   (response response.UsersResponse)
	DeleteUsersForInternal(request request.UsersRequest) (response response.UsersResponse)
	DeleteUsersForPrivate(request request.UsersRequest)  (response response.UsersResponse)
}

type usersRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (usersRepository usersRepository) BootstrapUsersForDB(request request.UsersRequest) (response response.UsersResponse) {
	fmt.Println("BootstrapUsersForDB")
	return response
}

//GET
func (usersRepository usersRepository) GetUsersForPublic(request request.UsersRequest) (response response.UsersResponse) {
	fmt.Println("GetUsersForPublic")
	return response
}

func (usersRepository usersRepository) GetUsersForInternal(request request.UsersRequest) (response response.UsersResponse ){
	fmt.Println("GetUsersForInternal")
	return response
}

func (usersRepository usersRepository) GetUsersForPrivate(request request.UsersRequest) (response response.UsersResponse) {
	fmt.Println("GetUsersForPrivate")
	return response
}

//CREATE
func (usersRepository usersRepository) CreateUsersForPublic(request request.UsersRequest) (response response.UsersResponse ){
	fmt.Println("CreateUsersForPublic")
	return response
}

func (usersRepository usersRepository) CreateUsersForInternal(request request.UsersRequest) (response response.UsersResponse) {
	fmt.Println("CreateUsersForInternal()")
	return response
}

func (usersRepository usersRepository) CreateUsersForPrivate(request request.UsersRequest) (response response.UsersResponse){
	fmt.Println("CreateUsersForPrivate()")
	return response
}

//UPDATE
func (usersRepository usersRepository) UpdateUsersForPublic(request request.UsersRequest) (response response.UsersResponse){
	fmt.Println("UpdateUsersForPublic()")
	return response
}

func (usersRepository usersRepository) UpdateUsersForInternal(request request.UsersRequest) (response response.UsersResponse){
	fmt.Println("UpdateUsersForInternal")
	return response
}

func (usersRepository usersRepository) UpdateUsersForPrivate(request request.UsersRequest) (response response.UsersResponse){
	fmt.Println("UpdateUsersForPrivate")
	return response
}

//DELETE
func (usersRepository usersRepository) DeleteUsersForPublic(request request.UsersRequest) (response response.UsersResponse){
	fmt.Println("DeleteUsersForPublic")
	return response
}

func (usersRepository usersRepository) DeleteUsersForInternal(request request.UsersRequest) (response response.UsersResponse ){
	fmt.Println("DeleteUsersForInternal")
	return response
}

func (usersRepository usersRepository) DeleteUsersForPrivate(request request.UsersRequest) (response response.UsersResponse){
	fmt.Println("DeleteUsersForPrivate")
	return response
}

func NewUsersRepository(conf config.BaseConfig) UsersRepository {
	return &usersRepository{BaseConfig: conf}
}