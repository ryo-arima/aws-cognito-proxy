package repository

import (
	"time"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type GroupsRepository interface {
	GetGroupss() []model.Groupss
	CreateGroups(groups model.Groupss) *gorm.DB
	UpdateGroups(groups model.Groupss) *gorm.DB
	DeleteGroups(uuid string) *gorm.DB
}

type groupsRepository struct {
	BaseConfig config.BaseConfig
}


func (groupsRepository groupsRepository) GetGroupss() []model.Groupss {
	var groupss []model.Groupss
	groupsRepository.BaseConfig.DBConnection.Find(&groupss)
	return groupss
}

func (groupsRepository groupsRepository) CreateGroups(groups model.Groupss) *gorm.DB {
	results := groupsRepository.BaseConfig.DBConnection.Create(&groups)
	return results
}

func (groupsRepository groupsRepository) UpdateGroups(groups model.Groupss) *gorm.DB {
	results := groupsRepository.BaseConfig.DBConnection.Model(model.Groupss{}).Where("uuid = ?", groups.UUID).Updates(groups)
	return results
}

func (groupsRepository groupsRepository) DeleteGroups(uuid string) *gorm.DB {
	results := groupsRepository.BaseConfig.DBConnection.Model(model.Groupss{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewGroupsRepository(conf config.BaseConfig) GroupsRepository {
	return &groupsRepository{BaseConfig: conf}
}