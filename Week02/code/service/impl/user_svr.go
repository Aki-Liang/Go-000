package impl

import (
	"log"

	"github.com/Aki-Liang/Go-000/Week02/code/dao"
	"github.com/Aki-Liang/Go-000/Week02/code/model"
)

type UserService struct {
	dao *dao.UserDao
}

func CreateUserService() *UserService {
	return &UserService{
		dao: dao.CreateUserDao(),
	}
}

func (s *UserService) GetUser(userId int) *model.User {
	user, err := s.dao.GetUserInfo(userId)
	if err != nil {
		log.Println("get user info failed, err:", err.Error())
		return model.GenDefaultUser()
	}

	return user
}
