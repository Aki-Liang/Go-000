package service

import (
	"github.com/Aki-Liang/Go-000/Week02/code/model"
	"github.com/Aki-Liang/Go-000/Week02/code/service/impl"
)

type IService interface {
	GetUser(userId int) *model.User
}

func CreateService() IService {
	return impl.CreateUserService()
}
