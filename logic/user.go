package logic

import (
	"errors"

	dao "pmc_server/dao/user"
	"pmc_server/libs/jwt"
	libs "pmc_server/libs/snowflake"
	model "pmc_server/model"
)

func Register(param *model.RegisterParams) error {
	exist, err := dao.UserExist(param.Email)
	if err != nil {
		return errors.New("failed to find user info")
	}

	if exist {
		return errors.New("user already exist")
	}

	userID := libs.GenID()
	return dao.InsertUser(&model.User{
		UserID:    userID,
		Email:     param.Email,
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Password:  param.Password,
		//FIXME: this needs to be changed in the future
		College: model.College{Name: "University of Utah"},
	})
}

func Login(param *model.LoginParams) (string, error) {
	user, err := dao.ReadUser(&model.User{
		Email:    param.Email,
		Password: param.Password,
	})
	if err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.FirstName, user.LastName)
}
