package repository

import (
	"time"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUserss() []model.Userss
	CreateUsers(users model.Userss) *gorm.DB
	UpdateUsers(users model.Userss) *gorm.DB
	DeleteUsers(uuid string) *gorm.DB
}

type usersRepository struct {
	BaseConfig config.BaseConfig
}


func (usersRepository usersRepository) GetUserss() []model.Userss {
	var userss []model.Userss
	usersRepository.BaseConfig.DBConnection.Find(&userss)
	return userss
}

func (usersRepository usersRepository) CreateUsers(users model.Userss) *gorm.DB {
	results := usersRepository.BaseConfig.DBConnection.Create(&users)
	return results
}

func (usersRepository usersRepository) UpdateUsers(users model.Userss) *gorm.DB {
	results := usersRepository.BaseConfig.DBConnection.Model(model.Userss{}).Where("uuid = ?", users.UUID).Updates(users)
	return results
}

func (usersRepository usersRepository) DeleteUsers(uuid string) *gorm.DB {
	results := usersRepository.BaseConfig.DBConnection.Model(model.Userss{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewUsersRepository(conf config.BaseConfig) UsersRepository {
	return &usersRepository{BaseConfig: conf}
}