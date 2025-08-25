package auth

import (
	"ecoply/internal/database"
	"ecoply/internal/domain/execute"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/services"
	"ecoply/internal/mlog"
	"net/http"
)

func Login(request *LoginRequest) (*LoginResource, *merr.ResponseError) {
	var user *models.User
	var userExec execute.UserExecute = execute.NewUserExecute(database.Con)

	user, err := userExec.FindUserByCredentials(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	if err := userExec.PreloadUserType(user); err != nil {
		return nil, err
	}

	token, tokenErr := generateJwtToken(user)
	if tokenErr != nil {
		return nil, tokenErr
	}

	response := &LoginResource{
		Token: token,
		User:  newMeResource(user),
	}

	return response, nil
}

func generateJwtToken(user *models.User) (string, *merr.ResponseError) {
	jwtService := services.NewJwtService()
	token, tokenErr := jwtService.GenerateToken(user.ID, user.Email, user.UserType.Type)
	if tokenErr != nil {
		mlog.Log("Failed to generate JWT token: " + tokenErr.Error())
		return "", merr.NewResponseError(http.StatusInternalServerError, ErrFailedToGenerateToken)
	}
	return token, nil
}

func SignUp(request *SignUpRequest) (*LoginResource, *merr.ResponseError) {
	var userExec execute.UserExecute = execute.NewUserExecute(database.Con)

	user, err := userExec.CreateWithAddressAndType(execute.UserCreateWithAddressAndTypeParams{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		CpfCnpj:  request.CpfCnpj,
		UserType: request.UserType,
		Address: execute.AddressCreateParams{
			Cep:     request.Address.Cep,
			State:   request.Address.State,
			City:    request.Address.City,
			Country: request.Address.Country,
		},
	})

	if err != nil {
		return nil, err
	}

	if err := userExec.PreloadUserType(user); err != nil {
		return nil, err
	}

	token, tokenErr := generateJwtToken(user)
	if tokenErr != nil {
		return nil, tokenErr
	}

	response := &LoginResource{
		Token: token,
		User:  newMeResource(user),
	}

	return response, nil
}

func Me(userId uint) (*MeResource, *merr.ResponseError) {
	var userExec execute.UserExecute = execute.NewUserExecute(database.Con)
	user, err := userExec.FindById(userId)
	if err != nil {
		return nil, err
	}

	if err := userExec.PreloadUserType(user); err != nil {
		return nil, err
	}

	response := newMeResource(user)

	return &response, nil
}

func newMeResource(user *models.User) MeResource {
	return MeResource{
		Name:     user.Name,
		UserType: user.UserType.Type,
	}
}

func IsEmailAvailable(request *IsEmailAvailableRequest) (bool, *merr.ResponseError) {
	var userExec execute.UserExecute = execute.NewUserExecute(database.Con)
	user, err := userExec.FindByEmail(request.Email)
	if err != nil {
		if err.StatusCode == http.StatusNotFound {
			return true, nil
		}
		return false, err
	}

	if user != nil {
		return false, nil
	}

	return true, nil
}

func IsCpfCnpjAvailable(request *IsCpfCnpjAvailableRequest) (bool, *merr.ResponseError) {
	var userExec execute.UserExecute = execute.NewUserExecute(database.Con)
	user, err := userExec.FindByCpfCnpj(request.CpfCnpj)
	if err != nil {
		if err.StatusCode == http.StatusNotFound {
			return true, nil
		}
		return false, err
	}

	if user != nil {
		return false, nil
	}

	return true, nil
}
