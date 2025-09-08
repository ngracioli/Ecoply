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
	var err *merr.ResponseError
	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)

	user, err = userRepo.FindUserByCredentials(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	if err = userRepo.PreloadUserType(user); err != nil {
		return nil, err
	}

	token, err := generateJwtToken(user)
	if err != nil {
		return nil, err
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
	var addressData *services.AddressData
	var err *merr.ResponseError

	if addressData, err = services.LoadAddressByCep(request.Address.Cep); err != nil {
		return nil, err
	}

	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)

	user, err := userRepo.CreateWithAddressAndType(repository.UserCreateWithAddressAndTypeParams{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Cnpj:     request.Cnpj,
		UserType: request.UserType,
		Address: repository.AddressCreateParams{
			Cep:           request.Address.Cep,
			Number:        request.Address.Number,
			Complement:    request.Address.Complement,
			State:         addressData.State,
			City:          addressData.City,
			Neighborhood:  addressData.Neighborhood,
			Street:        addressData.Street,
			StateInitials: addressData.StateInitials,
		},
	})

	if err != nil {
		return nil, err
	}

	if err := userRepo.PreloadUserType(user); err != nil {
		return nil, err
	}

	token, err := generateJwtToken(user)
	if err != nil {
		return nil, err
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
		return IsCnpjAvailable(request.Value)
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

func IsCnpjAvailable(cnpj string) (bool, *merr.ResponseError) {
	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)
	user, err := userRepo.FindByCnpj(cnpj)
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
