package repository

import (
	"fmt"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
)

type PoliciesRepository interface {
	BootstrapPoliciesForDB(request request.PoliciesRequest)(response response.PoliciesResponse) 
	GetPoliciesForPublic(request request.PoliciesRequest)(response response.PoliciesResponse) 
	GetPoliciesForInternal(request request.PoliciesRequest)(response response.PoliciesResponse) 
	GetPoliciesForPrivate(request request.PoliciesRequest)(response response.PoliciesResponse) 
	CreatePoliciesForPublic(request request.PoliciesRequest)   (response response.PoliciesResponse) 
	CreatePoliciesForInternal(request request.PoliciesRequest) (response response.PoliciesResponse)
	CreatePoliciesForPrivate(request request.PoliciesRequest)  (response response.PoliciesResponse)
	UpdatePoliciesForPublic(request request.PoliciesRequest)   (response response.PoliciesResponse)
	UpdatePoliciesForInternal(request request.PoliciesRequest) (response response.PoliciesResponse)
	UpdatePoliciesForPrivate(request request.PoliciesRequest)  (response response.PoliciesResponse)
	DeletePoliciesForPublic(request request.PoliciesRequest)   (response response.PoliciesResponse)
	DeletePoliciesForInternal(request request.PoliciesRequest) (response response.PoliciesResponse)
	DeletePoliciesForPrivate(request request.PoliciesRequest)  (response response.PoliciesResponse)
}

type policiesRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (policiesRepository policiesRepository) BootstrapPoliciesForDB(request request.PoliciesRequest) (response response.PoliciesResponse) {
	fmt.Println("BootstrapPoliciesForDB")
	return response
}

//GET
func (policiesRepository policiesRepository) GetPoliciesForPublic(request request.PoliciesRequest) (response response.PoliciesResponse) {
	fmt.Println("GetPoliciesForPublic")
	return response
}

func (policiesRepository policiesRepository) GetPoliciesForInternal(request request.PoliciesRequest) (response response.PoliciesResponse ){
	fmt.Println("GetPoliciesForInternal")
	return response
}

func (policiesRepository policiesRepository) GetPoliciesForPrivate(request request.PoliciesRequest) (response response.PoliciesResponse) {
	fmt.Println("GetPoliciesForPrivate")
	return response
}

//CREATE
func (policiesRepository policiesRepository) CreatePoliciesForPublic(request request.PoliciesRequest) (response response.PoliciesResponse ){
	fmt.Println("CreatePoliciesForPublic")
	return response
}

func (policiesRepository policiesRepository) CreatePoliciesForInternal(request request.PoliciesRequest) (response response.PoliciesResponse) {
	fmt.Println("CreatePoliciesForInternal()")
	return response
}

func (policiesRepository policiesRepository) CreatePoliciesForPrivate(request request.PoliciesRequest) (response response.PoliciesResponse){
	fmt.Println("CreatePoliciesForPrivate()")
	return response
}

//UPDATE
func (policiesRepository policiesRepository) UpdatePoliciesForPublic(request request.PoliciesRequest) (response response.PoliciesResponse){
	fmt.Println("UpdatePoliciesForPublic()")
	return response
}

func (policiesRepository policiesRepository) UpdatePoliciesForInternal(request request.PoliciesRequest) (response response.PoliciesResponse){
	fmt.Println("UpdatePoliciesForInternal")
	return response
}

func (policiesRepository policiesRepository) UpdatePoliciesForPrivate(request request.PoliciesRequest) (response response.PoliciesResponse){
	fmt.Println("UpdatePoliciesForPrivate")
	return response
}

//DELETE
func (policiesRepository policiesRepository) DeletePoliciesForPublic(request request.PoliciesRequest) (response response.PoliciesResponse){
	fmt.Println("DeletePoliciesForPublic")
	return response
}

func (policiesRepository policiesRepository) DeletePoliciesForInternal(request request.PoliciesRequest) (response response.PoliciesResponse ){
	fmt.Println("DeletePoliciesForInternal")
	return response
}

func (policiesRepository policiesRepository) DeletePoliciesForPrivate(request request.PoliciesRequest) (response response.PoliciesResponse){
	fmt.Println("DeletePoliciesForPrivate")
	return response
}

func NewPoliciesRepository(conf config.BaseConfig) PoliciesRepository {
	return &policiesRepository{BaseConfig: conf}
}