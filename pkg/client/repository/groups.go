package repository

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
)

type GroupsRepository interface {
	BootstrapGroupsForDB(request request.GroupsRequest)(response response.GroupsResponse) 
	GetGroupsForPublic(request request.GroupsRequest)(response response.GroupsResponse) 
	GetGroupsForInternal(request request.GroupsRequest)(response response.GroupsResponse) 
	GetGroupsForPrivate(request request.GroupsRequest)(response response.GroupsResponse) 
	CreateGroupsForPublic(request request.GroupsRequest)   (response response.GroupsResponse) 
	CreateGroupsForInternal(request request.GroupsRequest) (response response.GroupsResponse)
	CreateGroupsForPrivate(request request.GroupsRequest)  (response response.GroupsResponse)
	UpdateGroupsForPublic(request request.GroupsRequest)   (response response.GroupsResponse)
	UpdateGroupsForInternal(request request.GroupsRequest) (response response.GroupsResponse)
	UpdateGroupsForPrivate(request request.GroupsRequest)  (response response.GroupsResponse)
	DeleteGroupsForPublic(request request.GroupsRequest)   (response response.GroupsResponse)
	DeleteGroupsForInternal(request request.GroupsRequest) (response response.GroupsResponse)
	DeleteGroupsForPrivate(request request.GroupsRequest)  (response response.GroupsResponse)
}

type groupsRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (groupsRepository groupsRepository) BootstrapGroupsForDB(request request.GroupsRequest) (response response.GroupsResponse) {
	fmt.Println("BootstrapGroupsForDB")
	return response
}

//GET
func (groupsRepository groupsRepository) GetGroupsForPublic(request request.GroupsRequest) (response response.GroupsResponse) {
	fmt.Println("GetGroupsForPublic")
	return response
}

func (groupsRepository groupsRepository) GetGroupsForInternal(request request.GroupsRequest) (response response.GroupsResponse ){
	fmt.Println("GetGroupsForInternal")
	return response
}

func (groupsRepository groupsRepository) GetGroupsForPrivate(request request.GroupsRequest) (response response.GroupsResponse) {
	fmt.Println("GetGroupsForPrivate")
	return response
}

//CREATE
func (groupsRepository groupsRepository) CreateGroupsForPublic(request request.GroupsRequest) (response response.GroupsResponse ){
	fmt.Println("CreateGroupsForPublic")
	return response
}

func (groupsRepository groupsRepository) CreateGroupsForInternal(request request.GroupsRequest) (response response.GroupsResponse) {
	fmt.Println("CreateGroupsForInternal()")
	return response
}

func (groupsRepository groupsRepository) CreateGroupsForPrivate(request request.GroupsRequest) (response response.GroupsResponse){
	fmt.Println("CreateGroupsForPrivate()")
	return response
}

//UPDATE
func (groupsRepository groupsRepository) UpdateGroupsForPublic(request request.GroupsRequest) (response response.GroupsResponse){
	fmt.Println("UpdateGroupsForPublic()")
	return response
}

func (groupsRepository groupsRepository) UpdateGroupsForInternal(request request.GroupsRequest) (response response.GroupsResponse){
	fmt.Println("UpdateGroupsForInternal")
	return response
}

func (groupsRepository groupsRepository) UpdateGroupsForPrivate(request request.GroupsRequest) (response response.GroupsResponse){
	fmt.Println("UpdateGroupsForPrivate")
	return response
}

//DELETE
func (groupsRepository groupsRepository) DeleteGroupsForPublic(request request.GroupsRequest) (response response.GroupsResponse){
	fmt.Println("DeleteGroupsForPublic")
	return response
}

func (groupsRepository groupsRepository) DeleteGroupsForInternal(request request.GroupsRequest) (response response.GroupsResponse ){
	fmt.Println("DeleteGroupsForInternal")
	return response
}

func (groupsRepository groupsRepository) DeleteGroupsForPrivate(request request.GroupsRequest) (response response.GroupsResponse){
	fmt.Println("DeleteGroupsForPrivate")
	return response
}

func NewGroupsRepository(conf config.BaseConfig) GroupsRepository {
	return &groupsRepository{BaseConfig: conf}
}