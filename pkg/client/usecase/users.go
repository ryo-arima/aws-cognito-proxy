package usecase

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/repository"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
)

type UsersUsecase interface {
	BootstrapUsersForDB(request request.UsersRequest)
	GetUsersForPublic(request request.UsersRequest)
	GetUsersForInternal(request request.UsersRequest)
	GetUsersForPrivate(request request.UsersRequest)
	CreateUsersForPublic(request request.UsersRequest)
	CreateUsersForInternal(request request.UsersRequest)
	CreateUsersForPrivate(request request.UsersRequest)
	UpdateUsersForPublic(request request.UsersRequest)
	UpdateUsersForInternal(request request.UsersRequest)
	UpdateUsersForPrivate(request request.UsersRequest)
	DeleteUsersForPublic(request request.UsersRequest)
	DeleteUsersForInternal(request request.UsersRequest)
	DeleteUsersForPrivate(request request.UsersRequest)
}

type usersUsecase struct {
	UsersRepository   repository.UsersRepository
}

//Bootstrap
func (usersUsecase usersUsecase) BootstrapUsersForDB(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.BootstrapUsersForDB(request)
	fmt.Println(userss)
}

//GET
func (usersUsecase usersUsecase) GetUsersForPublic(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.GetUsersForPublic(request)
	fmt.Println(userss)
}

func (usersUsecase usersUsecase) GetUsersForInternal(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.GetUsersForInternal(request)
	fmt.Println(userss)
}

func (usersUsecase usersUsecase) GetUsersForPrivate(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.GetUsersForPrivate(request)
	fmt.Println(userss)
}

//CREATE
func (usersUsecase usersUsecase) CreateUsersForPublic(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.CreateUsersForPublic(request)
	fmt.Println(userss)
}

func (usersUsecase usersUsecase) CreateUsersForInternal(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.CreateUsersForInternal(request)
	fmt.Println(userss)
}

func (usersUsecase usersUsecase) CreateUsersForPrivate(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.CreateUsersForPrivate(request)
	fmt.Println(userss)
}

//UPDATE
func (usersUsecase usersUsecase) UpdateUsersForPublic(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.UpdateUsersForPublic(request)
	fmt.Println(userss)
}

func (usersUsecase usersUsecase) UpdateUsersForInternal(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.UpdateUsersForInternal(request)
	fmt.Println(userss)
}

func (usersUsecase usersUsecase) UpdateUsersForPrivate(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.UpdateUsersForPrivate(request)
	fmt.Println(userss)
}

//DELETE
func (usersUsecase usersUsecase) DeleteUsersForPublic(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.DeleteUsersForPublic(request)
	fmt.Println(userss)
}

func (usersUsecase usersUsecase) DeleteUsersForInternal(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.DeleteUsersForInternal(request)
	fmt.Println(userss)
}

func (usersUsecase usersUsecase) DeleteUsersForPrivate(request request.UsersRequest) {
	userss := usersUsecase.UsersRepository.DeleteUsersForPrivate(request)
	fmt.Println(userss)
}

func NewUsersUsecase(usersRepository repository.UsersRepository) UsersUsecase {
	return &usersUsecase{ UsersRepository: usersRepository}
}