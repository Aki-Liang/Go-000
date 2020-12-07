package dao

import (
	"errors"

	"github.com/Aki-Liang/Go-000/Week02/code/model"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("record not found")
)

type UserDao struct {
	db *gorm.DB
}

func CreateUserDao() *UserDao {
	return &UserDao{
		db: MustOpenDB(),
	}
}

func (dao *UserDao) GetUserInfo(userId int) (*model.User, error) {
	var userInfo model.User
	err := dao.db.Table("user_info").Where("user_id=?", userId).First(&userInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerrors.Wrap(ErrNotFound, "get user info failed")
		}
		return &userInfo, xerrors.Wrap(err, "get user info failed")
	}

	return &userInfo, nil
}

func MustOpenDB() *gorm.DB {
	db, err := gorm.Open("root:test@tcp(localhost:3306)/week02?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}
