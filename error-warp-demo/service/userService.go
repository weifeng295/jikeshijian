package service

import (
	"github.com/pkg/errors"
	"warp-demo/models"
)

var (
	UserModel = models.NewUsers()
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// GetUserById 根据用户id查询
func (us *UserService) GetUserById(userid int64) (*models.Users, error) {
	user, err := UserModel.GetUserById(userid)
	return user, errors.Wrap(err, "userService.go GetUserById get user failed\n")
}
