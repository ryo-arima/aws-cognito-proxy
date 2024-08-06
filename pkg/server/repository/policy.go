package repository

import (
	"time"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type PolicyRepository interface {
	GetPolicys() []model.Policys
	CreatePolicy(policy model.Policys) *gorm.DB
	UpdatePolicy(policy model.Policys) *gorm.DB
	DeletePolicy(uuid string) *gorm.DB
}

type policyRepository struct {
	BaseConfig config.BaseConfig
}


func (policyRepository policyRepository) GetPolicys() []model.Policys {
	var policys []model.Policys
	policyRepository.BaseConfig.DBConnection.Find(&policys)
	return policys
}

func (policyRepository policyRepository) CreatePolicy(policy model.Policys) *gorm.DB {
	results := policyRepository.BaseConfig.DBConnection.Create(&policy)
	return results
}

func (policyRepository policyRepository) UpdatePolicy(policy model.Policys) *gorm.DB {
	results := policyRepository.BaseConfig.DBConnection.Model(model.Policys{}).Where("uuid = ?", policy.UUID).Updates(policy)
	return results
}

func (policyRepository policyRepository) DeletePolicy(uuid string) *gorm.DB {
	results := policyRepository.BaseConfig.DBConnection.Model(model.Policys{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewPolicyRepository(conf config.BaseConfig) PolicyRepository {
	return &policyRepository{BaseConfig: conf}
}