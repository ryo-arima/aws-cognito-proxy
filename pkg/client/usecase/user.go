package usecase

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/repository"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
)

type UserUsecase interface {
	BootstrapUserForDB(request request.UserRequest)
	GetUserForPublic(request request.UserRequest)
	GetUserForInternal(request request.UserRequest)
	GetUserForPrivate(request request.UserRequest)
	CreateUserForPublic(request request.UserRequest)
	CreateUserForInternal(request request.UserRequest)
	CreateUserForPrivate(request request.UserRequest)
	UpdateUserForPublic(request request.UserRequest)
	UpdateUserForInternal(request request.UserRequest)
	UpdateUserForPrivate(request request.UserRequest)
	DeleteUserForPublic(request request.UserRequest)
	DeleteUserForInternal(request request.UserRequest)
	DeleteUserForPrivate(request request.UserRequest)
}

type userUsecase struct {
	UserRepository   repository.UserRepository
}

//Bootstrap
func (userUsecase userUsecase) BootstrapUserForDB(request request.UserRequest) {
	users := userUsecase.UserRepository.BootstrapUserForDB(request)
	fmt.Println(users)
}

//GET
func (userUsecase userUsecase) GetUserForPublic(request request.UserRequest) {
	users := userUsecase.UserRepository.GetUserForPublic(request)
	fmt.Println(users)
}

func (userUsecase userUsecase) GetUserForInternal(request request.UserRequest) {
	users := userUsecase.UserRepository.GetUserForInternal(request)
	fmt.Println(users)
}

func (userUsecase userUsecase) GetUserForPrivate(request request.UserRequest) {
	users := userUsecase.UserRepository.GetUserForPrivate(request)
	fmt.Println(users)
}

//CREATE
func (userUsecase userUsecase) CreateUserForPublic(request request.UserRequest) {
	users := userUsecase.UserRepository.CreateUserForPublic(request)
	fmt.Println(users)
}

func (userUsecase userUsecase) CreateUserForInternal(request request.UserRequest) {
	users := userUsecase.UserRepository.CreateUserForInternal(request)
	fmt.Println(users)
}

func (userUsecase userUsecase) CreateUserForPrivate(request request.UserRequest) {
	users := userUsecase.UserRepository.CreateUserForPrivate(request)
	fmt.Println(users)
}

//UPDATE
func (userUsecase userUsecase) UpdateUserForPublic(request request.UserRequest) {
	users := userUsecase.UserRepository.UpdateUserForPublic(request)
	fmt.Println(users)
}

func (userUsecase userUsecase) UpdateUserForInternal(request request.UserRequest) {
	users := userUsecase.UserRepository.UpdateUserForInternal(request)
	fmt.Println(users)
}

func (userUsecase userUsecase) UpdateUserForPrivate(request request.UserRequest) {
	users := userUsecase.UserRepository.UpdateUserForPrivate(request)
	fmt.Println(users)
}

//DELETE
func (userUsecase userUsecase) DeleteUserForPublic(request request.UserRequest) {
	users := userUsecase.UserRepository.DeleteUserForPublic(request)
	fmt.Println(users)
}

func (userUsecase userUsecase) DeleteUserForInternal(request request.UserRequest) {
	users := userUsecase.UserRepository.DeleteUserForInternal(request)
	fmt.Println(users)
}

func (userUsecase userUsecase) DeleteUserForPrivate(request request.UserRequest) {
	users := userUsecase.UserRepository.DeleteUserForPrivate(request)
	fmt.Println(users)
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{ UserRepository: userRepository}
}