package usecase

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/client/repository"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
)

type PolicyUsecase interface {
	BootstrapPolicyForDB(request request.PolicyRequest)
	GetPolicyForPublic(request request.PolicyRequest)
	GetPolicyForInternal(request request.PolicyRequest)
	GetPolicyForPrivate(request request.PolicyRequest)
	CreatePolicyForPublic(request request.PolicyRequest)
	CreatePolicyForInternal(request request.PolicyRequest)
	CreatePolicyForPrivate(request request.PolicyRequest)
	UpdatePolicyForPublic(request request.PolicyRequest)
	UpdatePolicyForInternal(request request.PolicyRequest)
	UpdatePolicyForPrivate(request request.PolicyRequest)
	DeletePolicyForPublic(request request.PolicyRequest)
	DeletePolicyForInternal(request request.PolicyRequest)
	DeletePolicyForPrivate(request request.PolicyRequest)
}

type policyUsecase struct {
	PolicyRepository   repository.PolicyRepository
}

//Bootstrap
func (policyUsecase policyUsecase) BootstrapPolicyForDB(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.BootstrapPolicyForDB(request)
	fmt.Println(policys)
}

//GET
func (policyUsecase policyUsecase) GetPolicyForPublic(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.GetPolicyForPublic(request)
	fmt.Println(policys)
}

func (policyUsecase policyUsecase) GetPolicyForInternal(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.GetPolicyForInternal(request)
	fmt.Println(policys)
}

func (policyUsecase policyUsecase) GetPolicyForPrivate(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.GetPolicyForPrivate(request)
	fmt.Println(policys)
}

//CREATE
func (policyUsecase policyUsecase) CreatePolicyForPublic(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.CreatePolicyForPublic(request)
	fmt.Println(policys)
}

func (policyUsecase policyUsecase) CreatePolicyForInternal(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.CreatePolicyForInternal(request)
	fmt.Println(policys)
}

func (policyUsecase policyUsecase) CreatePolicyForPrivate(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.CreatePolicyForPrivate(request)
	fmt.Println(policys)
}

//UPDATE
func (policyUsecase policyUsecase) UpdatePolicyForPublic(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.UpdatePolicyForPublic(request)
	fmt.Println(policys)
}

func (policyUsecase policyUsecase) UpdatePolicyForInternal(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.UpdatePolicyForInternal(request)
	fmt.Println(policys)
}

func (policyUsecase policyUsecase) UpdatePolicyForPrivate(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.UpdatePolicyForPrivate(request)
	fmt.Println(policys)
}

//DELETE
func (policyUsecase policyUsecase) DeletePolicyForPublic(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.DeletePolicyForPublic(request)
	fmt.Println(policys)
}

func (policyUsecase policyUsecase) DeletePolicyForInternal(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.DeletePolicyForInternal(request)
	fmt.Println(policys)
}

func (policyUsecase policyUsecase) DeletePolicyForPrivate(request request.PolicyRequest) {
	policys := policyUsecase.PolicyRepository.DeletePolicyForPrivate(request)
	fmt.Println(policys)
}

func NewPolicyUsecase(policyRepository repository.PolicyRepository) PolicyUsecase {
	return &policyUsecase{ PolicyRepository: policyRepository}
}