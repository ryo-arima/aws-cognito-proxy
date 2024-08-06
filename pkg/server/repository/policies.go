package repository

import (
	"time"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type PoliciesRepository interface {
	GetPoliciess() []model.Policiess
	CreatePolicies(policies model.Policiess) *gorm.DB
	UpdatePolicies(policies model.Policiess) *gorm.DB
	DeletePolicies(uuid string) *gorm.DB
}

type policiesRepository struct {
	BaseConfig config.BaseConfig
}


func (policiesRepository policiesRepository) GetPoliciess() []model.Policiess {
	var policiess []model.Policiess
	policiesRepository.BaseConfig.DBConnection.Find(&policiess)
	return policiess
}

func (policiesRepository policiesRepository) CreatePolicies(policies model.Policiess) *gorm.DB {
	results := policiesRepository.BaseConfig.DBConnection.Create(&policies)
	return results
}

func (policiesRepository policiesRepository) UpdatePolicies(policies model.Policiess) *gorm.DB {
	results := policiesRepository.BaseConfig.DBConnection.Model(model.Policiess{}).Where("uuid = ?", policies.UUID).Updates(policies)
	return results
}

func (policiesRepository policiesRepository) DeletePolicies(uuid string) *gorm.DB {
	results := policiesRepository.BaseConfig.DBConnection.Model(model.Policiess{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewPoliciesRepository(conf config.BaseConfig) PoliciesRepository {
	return &policiesRepository{BaseConfig: conf}
}