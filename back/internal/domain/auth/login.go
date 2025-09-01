package auth

import (
	"ecoply/internal/database"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/messages"
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
	"ecoply/internal/services"
	"errors"

	"gorm.io/gorm"
)

func login(request *LoginRequest) (*LoginResource, *merr.ResponseError) {
	user, err := authenticateUser(request)
	if err != nil {
		return nil, err
	}

	jwtService := services.NewJWTService()
	token, tokenErr := jwtService.GenerateToken(user.ID, user.Email, user.UserType.Type)
	if tokenErr != nil {
		mlog.Log("Failed to generate JWT token: " + tokenErr.Error())
		return nil, merr.NewResponseError(500, messages.AuthFailedToGenerateToken, nil)
	}

	response := &LoginResource{
		Token: token,
		User: MeResource{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			UserType: user.UserType.Type,
		},
	}

	return response, nil
}

func authenticateUser(request *LoginRequest) (*models.User, *merr.ResponseError) {
	var user models.User

	result := database.Con.Where("email = ?", request.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, merr.NewResponseError(401, messages.AuthEmailNotFound, nil)
		}
		mlog.Log(result.Error.Error())
		return nil, merr.NewResponseError(500, messages.DatabaseError, nil)
	}

	if user.Password != request.Password {
		return nil, merr.NewResponseError(401, messages.AuthIncorrectPassword, nil)
	}

	return &user, nil
}
