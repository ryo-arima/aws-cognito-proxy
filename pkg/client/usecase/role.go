package usecase

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/repository"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
)

type RoleUsecase interface {
	BootstrapRoleForDB(request request.RoleRequest)
	GetRoleForPublic(request request.RoleRequest)
	GetRoleForInternal(request request.RoleRequest)
	GetRoleForPrivate(request request.RoleRequest)
	CreateRoleForPublic(request request.RoleRequest)
	CreateRoleForInternal(request request.RoleRequest)
	CreateRoleForPrivate(request request.RoleRequest)
	UpdateRoleForPublic(request request.RoleRequest)
	UpdateRoleForInternal(request request.RoleRequest)
	UpdateRoleForPrivate(request request.RoleRequest)
	DeleteRoleForPublic(request request.RoleRequest)
	DeleteRoleForInternal(request request.RoleRequest)
	DeleteRoleForPrivate(request request.RoleRequest)
}

type roleUsecase struct {
	RoleRepository   repository.RoleRepository
}

//Bootstrap
func (roleUsecase roleUsecase) BootstrapRoleForDB(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.BootstrapRoleForDB(request)
	fmt.Println(roles)
}

//GET
func (roleUsecase roleUsecase) GetRoleForPublic(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.GetRoleForPublic(request)
	fmt.Println(roles)
}

func (roleUsecase roleUsecase) GetRoleForInternal(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.GetRoleForInternal(request)
	fmt.Println(roles)
}

func (roleUsecase roleUsecase) GetRoleForPrivate(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.GetRoleForPrivate(request)
	fmt.Println(roles)
}

//CREATE
func (roleUsecase roleUsecase) CreateRoleForPublic(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.CreateRoleForPublic(request)
	fmt.Println(roles)
}

func (roleUsecase roleUsecase) CreateRoleForInternal(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.CreateRoleForInternal(request)
	fmt.Println(roles)
}

func (roleUsecase roleUsecase) CreateRoleForPrivate(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.CreateRoleForPrivate(request)
	fmt.Println(roles)
}

//UPDATE
func (roleUsecase roleUsecase) UpdateRoleForPublic(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.UpdateRoleForPublic(request)
	fmt.Println(roles)
}

func (roleUsecase roleUsecase) UpdateRoleForInternal(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.UpdateRoleForInternal(request)
	fmt.Println(roles)
}

func (roleUsecase roleUsecase) UpdateRoleForPrivate(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.UpdateRoleForPrivate(request)
	fmt.Println(roles)
}

//DELETE
func (roleUsecase roleUsecase) DeleteRoleForPublic(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.DeleteRoleForPublic(request)
	fmt.Println(roles)
}

func (roleUsecase roleUsecase) DeleteRoleForInternal(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.DeleteRoleForInternal(request)
	fmt.Println(roles)
}

func (roleUsecase roleUsecase) DeleteRoleForPrivate(request request.RoleRequest) {
	roles := roleUsecase.RoleRepository.DeleteRoleForPrivate(request)
	fmt.Println(roles)
}

func NewRoleUsecase(roleRepository repository.RoleRepository) RoleUsecase {
	return &roleUsecase{ RoleRepository: roleRepository}
}