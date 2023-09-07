package user

import (
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/model/user"
)

type UserService struct{}

func (s *UserService) Add(u user.User) (err error) {
	err = global.DB.Create(&u).Error
	return err
}

// APP登录
func (s *UserService) AppLogin() {

}

// Web后台登录
func (s *UserService) WebLogin() {

}

// 小程序登录
func (s *UserService) MiniProgramLogin() {

}

// 用户注册
func (s *UserService) Register() {

}
