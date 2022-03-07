package dao

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
	"go.uber.org/zap"

	"pmc_server/init/postgres"
	"pmc_server/model"
)

// UserExist check if a user exists in the database
func UserExist(email string) (exist bool, err error) {
	var user model.User
	result := postgres.DB.Where(&model.User{Email: email}).Find(&user)
	if result.Error != nil {
		return true, result.Error
	}
	return result.RowsAffected != 0, err
}

// InsertUser creates a new user in the database
func InsertUser(user *model.User) error {
	user.Password = EncryptPassword(user.Password)

	result := postgres.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ReadUser basically checks the if the user has input the correct username and password for logging in
func ReadUser(user *model.User) (*model.User, error) {
	var u model.User
	result := postgres.DB.Where(&model.User{Email: user.Email}).Find(&u)
	if result.Error != nil {
		return nil, errors.New("failed to find user")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user does not exist")
	}

	zap.L().Info(u.Password)
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(u.Password, "$")
	check := password.Verify(user.Password, passwordInfo[2], passwordInfo[3], options)
	if !check {
		return nil, errors.New("user info does not match")
	}
	return &u, nil
}

// EncryptPassword encrypts a password with sha512
func EncryptPassword(pwd string) string {
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(pwd, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

// GetUserByID get a user info by given ID
func GetUserByID(userID int64) (*model.User, error) {
	var user model.User
	result := postgres.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
