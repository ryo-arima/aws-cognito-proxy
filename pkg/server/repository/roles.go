package repository

import (
	"time"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type RolesRepository interface {
	GetRoless() []model.Roless
	CreateRoles(roles model.Roless) *gorm.DB
	UpdateRoles(roles model.Roless) *gorm.DB
	DeleteRoles(uuid string) *gorm.DB
}

type rolesRepository struct {
	BaseConfig config.BaseConfig
}


func (rolesRepository rolesRepository) GetRoless() []model.Roless {
	var roless []model.Roless
	rolesRepository.BaseConfig.DBConnection.Find(&roless)
	return roless
}

func (rolesRepository rolesRepository) CreateRoles(roles model.Roless) *gorm.DB {
	results := rolesRepository.BaseConfig.DBConnection.Create(&roles)
	return results
}

func (rolesRepository rolesRepository) UpdateRoles(roles model.Roless) *gorm.DB {
	results := rolesRepository.BaseConfig.DBConnection.Model(model.Roless{}).Where("uuid = ?", roles.UUID).Updates(roles)
	return results
}

func (rolesRepository rolesRepository) DeleteRoles(uuid string) *gorm.DB {
	results := rolesRepository.BaseConfig.DBConnection.Model(model.Roless{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewRolesRepository(conf config.BaseConfig) RolesRepository {
	return &rolesRepository{BaseConfig: conf}
}