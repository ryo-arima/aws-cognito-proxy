package repository

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
)

type PolicyRepository interface {
	BootstrapPolicyForDB(request request.PolicyRequest)(response response.PolicyResponse) 
	GetPolicyForPublic(request request.PolicyRequest)(response response.PolicyResponse) 
	GetPolicyForInternal(request request.PolicyRequest)(response response.PolicyResponse) 
	GetPolicyForPrivate(request request.PolicyRequest)(response response.PolicyResponse) 
	CreatePolicyForPublic(request request.PolicyRequest)   (response response.PolicyResponse) 
	CreatePolicyForInternal(request request.PolicyRequest) (response response.PolicyResponse)
	CreatePolicyForPrivate(request request.PolicyRequest)  (response response.PolicyResponse)
	UpdatePolicyForPublic(request request.PolicyRequest)   (response response.PolicyResponse)
	UpdatePolicyForInternal(request request.PolicyRequest) (response response.PolicyResponse)
	UpdatePolicyForPrivate(request request.PolicyRequest)  (response response.PolicyResponse)
	DeletePolicyForPublic(request request.PolicyRequest)   (response response.PolicyResponse)
	DeletePolicyForInternal(request request.PolicyRequest) (response response.PolicyResponse)
	DeletePolicyForPrivate(request request.PolicyRequest)  (response response.PolicyResponse)
}

type policyRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (policyRepository policyRepository) BootstrapPolicyForDB(request request.PolicyRequest) (response response.PolicyResponse) {
	fmt.Println("BootstrapPolicyForDB")
	return response
}

//GET
func (policyRepository policyRepository) GetPolicyForPublic(request request.PolicyRequest) (response response.PolicyResponse) {
	fmt.Println("GetPolicyForPublic")
	return response
}

func (policyRepository policyRepository) GetPolicyForInternal(request request.PolicyRequest) (response response.PolicyResponse ){
	fmt.Println("GetPolicyForInternal")
	return response
}

func (policyRepository policyRepository) GetPolicyForPrivate(request request.PolicyRequest) (response response.PolicyResponse) {
	fmt.Println("GetPolicyForPrivate")
	return response
}

//CREATE
func (policyRepository policyRepository) CreatePolicyForPublic(request request.PolicyRequest) (response response.PolicyResponse ){
	fmt.Println("CreatePolicyForPublic")
	return response
}

func (policyRepository policyRepository) CreatePolicyForInternal(request request.PolicyRequest) (response response.PolicyResponse) {
	fmt.Println("CreatePolicyForInternal()")
	return response
}

func (policyRepository policyRepository) CreatePolicyForPrivate(request request.PolicyRequest) (response response.PolicyResponse){
	fmt.Println("CreatePolicyForPrivate()")
	return response
}

//UPDATE
func (policyRepository policyRepository) UpdatePolicyForPublic(request request.PolicyRequest) (response response.PolicyResponse){
	fmt.Println("UpdatePolicyForPublic()")
	return response
}

func (policyRepository policyRepository) UpdatePolicyForInternal(request request.PolicyRequest) (response response.PolicyResponse){
	fmt.Println("UpdatePolicyForInternal")
	return response
}

func (policyRepository policyRepository) UpdatePolicyForPrivate(request request.PolicyRequest) (response response.PolicyResponse){
	fmt.Println("UpdatePolicyForPrivate")
	return response
}

//DELETE
func (policyRepository policyRepository) DeletePolicyForPublic(request request.PolicyRequest) (response response.PolicyResponse){
	fmt.Println("DeletePolicyForPublic")
	return response
}

func (policyRepository policyRepository) DeletePolicyForInternal(request request.PolicyRequest) (response response.PolicyResponse ){
	fmt.Println("DeletePolicyForInternal")
	return response
}

func (policyRepository policyRepository) DeletePolicyForPrivate(request request.PolicyRequest) (response response.PolicyResponse){
	fmt.Println("DeletePolicyForPrivate")
	return response
}

func NewPolicyRepository(conf config.BaseConfig) PolicyRepository {
	return &policyRepository{BaseConfig: conf}
}