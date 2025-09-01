package auth

import (
	"ecoply/internal/database"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/messages"
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
)

func login(request *LoginRequest) (*MeResource, *merr.ResponseError) {
	if err := checkEmail(request); err != nil {
		return nil, err
	}

	user, err := checkLogin(request)
	if err != nil {
		return nil, err
	}

	var response MeResource = MeResource{
		Name:     user.Name,
		Email:    user.Email,
		UserType: user.UserType.Type,
	}

	return &response, nil
}

func checkEmail(request *LoginRequest) *merr.ResponseError {
	var user models.User
	result := database.Con.Where("email = ?", request.Email).First(&user)
	if result.Error != nil {
		mlog.Log(result.Error.Error())
		return merr.NewResponseError(404, messages.UserWithEmailNotFound, nil)
	}
	return nil
}

func checkLogin(request *LoginRequest) (*models.User, *merr.ResponseError) {
	var user models.User
	// TODO implement hash
	var hashedPasswrod string = request.Password

	result := database.Con.Where("email = ? AND password = ?", request.Email, hashedPasswrod).First(&user)
	if result.Error != nil {
		mlog.Log(result.Error.Error())
		return nil, merr.NewResponseError(404, messages.UserWithEmailNotFound, nil)
	}
	return &user, nil
}
