package usecase

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/repository"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
)

type GroupsUsecase interface {
	BootstrapGroupsForDB(request request.GroupsRequest)
	GetGroupsForPublic(request request.GroupsRequest)
	GetGroupsForInternal(request request.GroupsRequest)
	GetGroupsForPrivate(request request.GroupsRequest)
	CreateGroupsForPublic(request request.GroupsRequest)
	CreateGroupsForInternal(request request.GroupsRequest)
	CreateGroupsForPrivate(request request.GroupsRequest)
	UpdateGroupsForPublic(request request.GroupsRequest)
	UpdateGroupsForInternal(request request.GroupsRequest)
	UpdateGroupsForPrivate(request request.GroupsRequest)
	DeleteGroupsForPublic(request request.GroupsRequest)
	DeleteGroupsForInternal(request request.GroupsRequest)
	DeleteGroupsForPrivate(request request.GroupsRequest)
}

type groupsUsecase struct {
	GroupsRepository   repository.GroupsRepository
}

//Bootstrap
func (groupsUsecase groupsUsecase) BootstrapGroupsForDB(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.BootstrapGroupsForDB(request)
	fmt.Println(groupss)
}

//GET
func (groupsUsecase groupsUsecase) GetGroupsForPublic(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.GetGroupsForPublic(request)
	fmt.Println(groupss)
}

func (groupsUsecase groupsUsecase) GetGroupsForInternal(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.GetGroupsForInternal(request)
	fmt.Println(groupss)
}

func (groupsUsecase groupsUsecase) GetGroupsForPrivate(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.GetGroupsForPrivate(request)
	fmt.Println(groupss)
}

//CREATE
func (groupsUsecase groupsUsecase) CreateGroupsForPublic(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.CreateGroupsForPublic(request)
	fmt.Println(groupss)
}

func (groupsUsecase groupsUsecase) CreateGroupsForInternal(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.CreateGroupsForInternal(request)
	fmt.Println(groupss)
}

func (groupsUsecase groupsUsecase) CreateGroupsForPrivate(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.CreateGroupsForPrivate(request)
	fmt.Println(groupss)
}

//UPDATE
func (groupsUsecase groupsUsecase) UpdateGroupsForPublic(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.UpdateGroupsForPublic(request)
	fmt.Println(groupss)
}

func (groupsUsecase groupsUsecase) UpdateGroupsForInternal(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.UpdateGroupsForInternal(request)
	fmt.Println(groupss)
}

func (groupsUsecase groupsUsecase) UpdateGroupsForPrivate(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.UpdateGroupsForPrivate(request)
	fmt.Println(groupss)
}

//DELETE
func (groupsUsecase groupsUsecase) DeleteGroupsForPublic(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.DeleteGroupsForPublic(request)
	fmt.Println(groupss)
}

func (groupsUsecase groupsUsecase) DeleteGroupsForInternal(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.DeleteGroupsForInternal(request)
	fmt.Println(groupss)
}

func (groupsUsecase groupsUsecase) DeleteGroupsForPrivate(request request.GroupsRequest) {
	groupss := groupsUsecase.GroupsRepository.DeleteGroupsForPrivate(request)
	fmt.Println(groupss)
}

func NewGroupsUsecase(groupsRepository repository.GroupsRepository) GroupsUsecase {
	return &groupsUsecase{ GroupsRepository: groupsRepository}
}