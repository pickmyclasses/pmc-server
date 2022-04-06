package logic

import (
	"errors"
	"pmc_server/dao/postgres/user"
	"pmc_server/model/dto"

	"pmc_server/libs/jwt"
	libs "pmc_server/libs/snowflake"
	model "pmc_server/model"
)

func Register(param *model.RegisterParams) error {
	exist, err := dao.UserExist(param.Email)
	if err != nil {
		return err
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
		CollegeID: int64(param.CollegeID),
	})
}

func Login(param *model.LoginParams) (*dto.User, error) {
	user, err := dao.ReadUser(&model.User{
		Email:    param.Email,
		Password: param.Password,
	})
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenToken(user.UserID, user.FirstName, user.LastName)
	if err != nil {
		return nil, err
	}
	return &dto.User{
		ID:        user.ID,
		Token:     token,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		CollegeID: int32(user.CollegeID),
	}, nil
}

func GetUserHistoryCourseList(userID int64) ([]int64, error)  {
	historyCourseList, err := dao.GetUserHistoryCourseList(userID)
	if err != nil {
		return nil, err
	}
	return historyCourseList, nil
}

func AddUserCourseHistory(userID, courseID int64) error {
	err := dao.AddUserHistoryCourse(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveUserCourseHistory(userID, courseID int64) error {
	err := dao.RemoveUserHistoryCourse(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}