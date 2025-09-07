package auth

import (
	"ecoply/internal/database"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"ecoply/internal/domain/services"
	"ecoply/internal/mlog"
	"net/http"
)

func Login(request *LoginRequest) (*LoginResource, *merr.ResponseError) {
	var user *models.User
	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)

	user, err := userRepo.FindUserByCredentials(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	if err := userRepo.PreloadUserType(user); err != nil {
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
	token, tokenErr := jwtService.GenerateToken(user.Uuid, user.Email, user.UserType.Type)
	if tokenErr != nil {
		mlog.Log("Failed to generate JWT token: " + tokenErr.Error())
		return "", merr.NewResponseError(http.StatusInternalServerError, ErrFailedToGenerateToken)
	}
	return token, nil
}

func SignUp(request *SignUpRequest) (*LoginResource, *merr.ResponseError) {
	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)

	user, err := userRepo.CreateWithAddressAndType(repository.UserCreateWithAddressAndTypeParams{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		CpfCnpj:  request.CpfCnpj,
		UserType: request.UserType,
		Address: repository.AddressCreateParams{
			Cep:     request.Address.Cep,
			State:   request.Address.State,
			City:    request.Address.City,
			Country: request.Address.Country,
		},
	})

	if err != nil {
		return nil, err
	}

	if err := userRepo.PreloadUserType(user); err != nil {
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

func Me(userUuid string) (*MeResource, *merr.ResponseError) {
	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)
	user, err := userRepo.FindByUuid(userUuid)
	if err != nil {
		return nil, err
	}

	if err := userRepo.PreloadUserType(user); err != nil {
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

func Availability(request *AvailabilityRequest) (bool, *merr.ResponseError) {
	switch request.Type {
	case "email":
		return IsEmailAvailable(request.Value)
	case "cpf", "cnpj":
		return IsCpfCnpjAvailable(request.Value)
	default:
		return false, merr.NewResponseError(http.StatusBadRequest, ErrInvalidAvailabilityType)
	}
}

func IsEmailAvailable(email string) (bool, *merr.ResponseError) {
	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)
	user, err := userRepo.FindByEmail(email)
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

func IsCpfCnpjAvailable(cpfCnpj string) (bool, *merr.ResponseError) {
	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)
	user, err := userRepo.FindByCpfCnpj(cpfCnpj)
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
