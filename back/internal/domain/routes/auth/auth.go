package auth

import (
	"ecoply/internal/database"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"ecoply/internal/domain/services"
	"ecoply/internal/mlog"
	"net/http"

	"gorm.io/gorm"
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
	var err *merr.ResponseError
	var user *models.User

	user, err = createUser(request)
	if err != nil {
		return nil, err
	}

	var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)
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
	case "cnpj":
		return IsCnpjAvailable(request.Value)
	case "ccee":
		return IsCceeCodeAvailable(request.Value)
	default:
		return false, merr.NewResponseError(http.StatusBadRequest, ErrInvalidAvailabilityType)
	}
}

func IsEmailAvailable(email string) (bool, *merr.ResponseError) {
	return verifyIfAvailable(email, func(value string) (any, *merr.ResponseError) {
		var userRepo repository.UserRepository = repository.NewUserRepository(database.Con)
		return userRepo.FindByEmail(value)
	})
}

func IsCnpjAvailable(cnpj string) (bool, *merr.ResponseError) {
	return verifyIfAvailable(cnpj, func(value string) (any, *merr.ResponseError) {
		var agentRepo repository.AgentRepository = repository.NewAgentRepository(database.Con)
		return agentRepo.FindByCnpj(value)
	})
}

func IsCceeCodeAvailable(cceeCode string) (bool, *merr.ResponseError) {
	return verifyIfAvailable(cceeCode, func(value string) (any, *merr.ResponseError) {
		var agentRepo repository.AgentRepository = repository.NewAgentRepository(database.Con)
		return agentRepo.FindByCceeCode(value)
	})
}

func verifyIfAvailable(value string, findFunc func(string) (any, *merr.ResponseError)) (bool, *merr.ResponseError) {
	exists, err := findFunc(value)
	if err != nil {
		if err.StatusCode == http.StatusNotFound {
			return true, nil
		}
		return false, err
	}

	if exists != nil {
		return false, nil
	}

	return true, nil
}

func RefreshToken(oldToken string) (string, *merr.ResponseError) {
	jwtService := services.NewJwtService()
	newToken, err := jwtService.RefreshToken(oldToken)
	if err != nil {
		mlog.Log("Failed to refresh JWT token: " + err.Error())
		return "", merr.NewResponseError(http.StatusInternalServerError, ErrFailedToGenerateToken)
	}
	return newToken, nil
}

func createUser(request *SignUpRequest) (*models.User, *merr.ResponseError) {
	var responseError *merr.ResponseError
	var user *models.User

	database.Con.Transaction(func(tx *gorm.DB) error {
		var addressRepo repository.AddressRepository = repository.NewAddressRepository(tx)
		addressModel, err := addressRepo.Create(repository.AddressCreateParams{
			Cep:           request.Address.Cep,
			Number:        request.Address.Number,
			Complement:    request.Address.Complement,
			Street:        request.Address.Street,
			Neighborhood:  request.Address.Neighborhood,
			City:          request.Address.City,
			State:         request.Address.State,
			StateInitials: request.Address.StateInitials,
		})
		if err != nil {
			responseError = err
			return err.Error
		}

		var submarketRepo repository.SubmarketRepository = repository.NewSubmarketRepository(tx)
		submarketModel, err := submarketRepo.FindByName(request.Agent.SubmarketName)
		if err != nil {
			responseError = err
			return err.Error
		}

		var agentRepo repository.AgentRepository = repository.NewAgentRepository(tx)
		agentModel, err := agentRepo.Create(repository.AgentCreateParams{
			Cnpj:        request.Agent.Cnpj,
			CompanyName: request.Agent.CompanyName,
			CceeCode:    request.Agent.CceeCode,
			Submarket:   submarketModel,
			Address:     addressModel,
		})
		if err != nil {
			responseError = err
			return err.Error
		}

		var userTypeRepo repository.UserTypeRepository = repository.NewUserTypeRepository(tx)
		userTypeModel, err := userTypeRepo.FindByName(request.UserType)
		if err != nil {
			responseError = err
			return err.Error
		}

		var userRepo repository.UserRepository = repository.NewUserRepository(tx)
		user, err = userRepo.Create(repository.UserCreateParams{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
			UserType: userTypeModel,
			Agent:    agentModel,
		})
		if err != nil {
			responseError = err
			return err.Error
		}

		return nil
	})

	if responseError != nil {
		return nil, responseError
	}

	return user, nil
}
