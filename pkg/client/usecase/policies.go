package usecase

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/repository"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
)

type PoliciesUsecase interface {
	BootstrapPoliciesForDB(request request.PoliciesRequest)
	GetPoliciesForPublic(request request.PoliciesRequest)
	GetPoliciesForInternal(request request.PoliciesRequest)
	GetPoliciesForPrivate(request request.PoliciesRequest)
	CreatePoliciesForPublic(request request.PoliciesRequest)
	CreatePoliciesForInternal(request request.PoliciesRequest)
	CreatePoliciesForPrivate(request request.PoliciesRequest)
	UpdatePoliciesForPublic(request request.PoliciesRequest)
	UpdatePoliciesForInternal(request request.PoliciesRequest)
	UpdatePoliciesForPrivate(request request.PoliciesRequest)
	DeletePoliciesForPublic(request request.PoliciesRequest)
	DeletePoliciesForInternal(request request.PoliciesRequest)
	DeletePoliciesForPrivate(request request.PoliciesRequest)
}

type policiesUsecase struct {
	PoliciesRepository   repository.PoliciesRepository
}

//Bootstrap
func (policiesUsecase policiesUsecase) BootstrapPoliciesForDB(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.BootstrapPoliciesForDB(request)
	fmt.Println(policiess)
}

//GET
func (policiesUsecase policiesUsecase) GetPoliciesForPublic(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.GetPoliciesForPublic(request)
	fmt.Println(policiess)
}

func (policiesUsecase policiesUsecase) GetPoliciesForInternal(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.GetPoliciesForInternal(request)
	fmt.Println(policiess)
}

func (policiesUsecase policiesUsecase) GetPoliciesForPrivate(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.GetPoliciesForPrivate(request)
	fmt.Println(policiess)
}

//CREATE
func (policiesUsecase policiesUsecase) CreatePoliciesForPublic(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.CreatePoliciesForPublic(request)
	fmt.Println(policiess)
}

func (policiesUsecase policiesUsecase) CreatePoliciesForInternal(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.CreatePoliciesForInternal(request)
	fmt.Println(policiess)
}

func (policiesUsecase policiesUsecase) CreatePoliciesForPrivate(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.CreatePoliciesForPrivate(request)
	fmt.Println(policiess)
}

//UPDATE
func (policiesUsecase policiesUsecase) UpdatePoliciesForPublic(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.UpdatePoliciesForPublic(request)
	fmt.Println(policiess)
}

func (policiesUsecase policiesUsecase) UpdatePoliciesForInternal(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.UpdatePoliciesForInternal(request)
	fmt.Println(policiess)
}

func (policiesUsecase policiesUsecase) UpdatePoliciesForPrivate(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.UpdatePoliciesForPrivate(request)
	fmt.Println(policiess)
}

//DELETE
func (policiesUsecase policiesUsecase) DeletePoliciesForPublic(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.DeletePoliciesForPublic(request)
	fmt.Println(policiess)
}

func (policiesUsecase policiesUsecase) DeletePoliciesForInternal(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.DeletePoliciesForInternal(request)
	fmt.Println(policiess)
}

func (policiesUsecase policiesUsecase) DeletePoliciesForPrivate(request request.PoliciesRequest) {
	policiess := policiesUsecase.PoliciesRepository.DeletePoliciesForPrivate(request)
	fmt.Println(policiess)
}

func NewPoliciesUsecase(policiesRepository repository.PoliciesRepository) PoliciesUsecase {
	return &policiesUsecase{ PoliciesRepository: policiesRepository}
}