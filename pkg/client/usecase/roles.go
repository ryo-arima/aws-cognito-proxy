package usecase

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/repository"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
)

type RolesUsecase interface {
	BootstrapRolesForDB(request request.RolesRequest)
	GetRolesForPublic(request request.RolesRequest)
	GetRolesForInternal(request request.RolesRequest)
	GetRolesForPrivate(request request.RolesRequest)
	CreateRolesForPublic(request request.RolesRequest)
	CreateRolesForInternal(request request.RolesRequest)
	CreateRolesForPrivate(request request.RolesRequest)
	UpdateRolesForPublic(request request.RolesRequest)
	UpdateRolesForInternal(request request.RolesRequest)
	UpdateRolesForPrivate(request request.RolesRequest)
	DeleteRolesForPublic(request request.RolesRequest)
	DeleteRolesForInternal(request request.RolesRequest)
	DeleteRolesForPrivate(request request.RolesRequest)
}

type rolesUsecase struct {
	RolesRepository   repository.RolesRepository
}

//Bootstrap
func (rolesUsecase rolesUsecase) BootstrapRolesForDB(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.BootstrapRolesForDB(request)
	fmt.Println(roless)
}

//GET
func (rolesUsecase rolesUsecase) GetRolesForPublic(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.GetRolesForPublic(request)
	fmt.Println(roless)
}

func (rolesUsecase rolesUsecase) GetRolesForInternal(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.GetRolesForInternal(request)
	fmt.Println(roless)
}

func (rolesUsecase rolesUsecase) GetRolesForPrivate(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.GetRolesForPrivate(request)
	fmt.Println(roless)
}

//CREATE
func (rolesUsecase rolesUsecase) CreateRolesForPublic(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.CreateRolesForPublic(request)
	fmt.Println(roless)
}

func (rolesUsecase rolesUsecase) CreateRolesForInternal(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.CreateRolesForInternal(request)
	fmt.Println(roless)
}

func (rolesUsecase rolesUsecase) CreateRolesForPrivate(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.CreateRolesForPrivate(request)
	fmt.Println(roless)
}

//UPDATE
func (rolesUsecase rolesUsecase) UpdateRolesForPublic(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.UpdateRolesForPublic(request)
	fmt.Println(roless)
}

func (rolesUsecase rolesUsecase) UpdateRolesForInternal(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.UpdateRolesForInternal(request)
	fmt.Println(roless)
}

func (rolesUsecase rolesUsecase) UpdateRolesForPrivate(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.UpdateRolesForPrivate(request)
	fmt.Println(roless)
}

//DELETE
func (rolesUsecase rolesUsecase) DeleteRolesForPublic(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.DeleteRolesForPublic(request)
	fmt.Println(roless)
}

func (rolesUsecase rolesUsecase) DeleteRolesForInternal(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.DeleteRolesForInternal(request)
	fmt.Println(roless)
}

func (rolesUsecase rolesUsecase) DeleteRolesForPrivate(request request.RolesRequest) {
	roless := rolesUsecase.RolesRepository.DeleteRolesForPrivate(request)
	fmt.Println(roless)
}

func NewRolesUsecase(rolesRepository repository.RolesRepository) RolesUsecase {
	return &rolesUsecase{ RolesRepository: rolesRepository}
}