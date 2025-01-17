package repository

import (
	"time"

	"github.com/ryo-arima/aws-cognito-proxy/pkg/config"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() []model.Users
	CreateUser(user model.Users) *gorm.DB
	UpdateUser(user model.Users) *gorm.DB
	DeleteUser(uuid string) *gorm.DB
}

type userRepository struct {
	BaseConfig config.BaseConfig
}


func (userRepository userRepository) GetUsers() []model.Users {
	var users []model.Users
	userRepository.BaseConfig.DBConnection.Find(&users)
	return users
}

func (userRepository userRepository) CreateUser(user model.Users) *gorm.DB {
	results := userRepository.BaseConfig.DBConnection.Create(&user)
	return results
}

func (userRepository userRepository) UpdateUser(user model.Users) *gorm.DB {
	results := userRepository.BaseConfig.DBConnection.Model(model.Users{}).Where("uuid = ?", user.UUID).Updates(user)
	return results
}

func (userRepository userRepository) DeleteUser(uuid string) *gorm.DB {
	results := userRepository.BaseConfig.DBConnection.Model(model.Users{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewUserRepository(conf config.BaseConfig) UserRepository {
	return &userRepository{BaseConfig: conf}
}