package repository

import (
	"time"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoles() []model.Roles
	CreateRole(role model.Roles) *gorm.DB
	UpdateRole(role model.Roles) *gorm.DB
	DeleteRole(uuid string) *gorm.DB
}

type roleRepository struct {
	BaseConfig config.BaseConfig
}


func (roleRepository roleRepository) GetRoles() []model.Roles {
	var roles []model.Roles
	roleRepository.BaseConfig.DBConnection.Find(&roles)
	return roles
}

func (roleRepository roleRepository) CreateRole(role model.Roles) *gorm.DB {
	results := roleRepository.BaseConfig.DBConnection.Create(&role)
	return results
}

func (roleRepository roleRepository) UpdateRole(role model.Roles) *gorm.DB {
	results := roleRepository.BaseConfig.DBConnection.Model(model.Roles{}).Where("uuid = ?", role.UUID).Updates(role)
	return results
}

func (roleRepository roleRepository) DeleteRole(uuid string) *gorm.DB {
	results := roleRepository.BaseConfig.DBConnection.Model(model.Roles{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewRoleRepository(conf config.BaseConfig) RoleRepository {
	return &roleRepository{BaseConfig: conf}
}