package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/resources"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(request *requests.Login) (*resources.Login, *merr.ResponseError)
	SignUp(request *requests.SignUp) (*resources.Login, *merr.ResponseError)
	Me(userUuid string) (*resources.Me, *merr.ResponseError)
	Availability(request *requests.Availability) (bool, *merr.ResponseError)
	RefreshToken(oldToken string) (string, *merr.ResponseError)
}

type authService struct {
	userRepo      repository.UserRepository
	agentRepo     repository.AgentRepository
	addressRepo   repository.AddressRepository
	userTypeRepo  repository.UserTypeRepository
	submarketRepo repository.SubmarketRepository
	jwtService    *JwtService
	db            *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{
		userRepo:      repository.NewUserRepository(db),
		agentRepo:     repository.NewAgentRepository(db),
		addressRepo:   repository.NewAddressRepository(db),
		userTypeRepo:  repository.NewUserTypeRepository(db),
		submarketRepo: repository.NewSubmarketRepository(db),
		jwtService:    NewJwtService(),
		db:            db,
	}
}

func (s *authService) makeMeResource(user *models.User) (*resources.Me, *merr.ResponseError) {
	var err error
	if err = s.db.Preload("Agent").
		Preload("Agent.Address").
		Preload("Agent.Address.Street").
		Preload("Agent.Address.Street.Neighborhood").
		Preload("Agent.Address.Street.Neighborhood.City").
		Preload("Agent.Address.Street.Neighborhood.City.State").
		Preload("Agent.Submarket").
		Preload("UserType").
		First(user).Error; err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	var address models.Address = user.Agent.Address
	var agent models.Agent = user.Agent

	var me resources.Me = resources.Me{
		Name:     user.Name,
		Email:    user.Email,
		UserType: user.UserType.Type,
		Address: resources.Address{
			Cep:          address.Cep,
			Street:       address.Street.Street,
			Number:       address.Number,
			Complement:   address.Complement,
			Neighborhood: address.Street.Neighborhood.Neighborhood,
			StateInitial: address.Street.Neighborhood.City.State.Initials,
			City:         address.Street.Neighborhood.City.City,
			State:        address.Street.Neighborhood.City.State.State,
		},
		Agent: resources.Agent{
			Cnpj:          agent.Cnpj,
			CceeCode:      agent.CceeCode,
			SubmarketName: agent.Submarket.Name,
			CompanyName:   agent.CompanyName,
		},
	}

	return &me, nil
}

func (s *authService) Login(request *requests.Login) (*resources.Login, *merr.ResponseError) {
	var err error
	var errResponse *merr.ResponseError

	var user *models.User
	var token string
	var meResource *resources.Me
	var response *resources.Login

	user, err = s.userRepo.FindByEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidCredentials)
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if user.Password != Hash256String(request.Password) {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrIncorrectPassword)
	}

	token, errResponse = s.generateJwtToken(user)
	if errResponse != nil {
		return nil, errResponse
	}

	meResource, errResponse = s.makeMeResource(user)
	if errResponse != nil {
		return nil, errResponse
	}

	response = &resources.Login{
		Token: token,
		User:  *meResource,
	}

	return response, nil
}

func (s *authService) SignUp(request *requests.SignUp) (*resources.Login, *merr.ResponseError) {
	var errResponse *merr.ResponseError
	var user *models.User
	var token string
	var meResource *resources.Me
	var response *resources.Login

	user, errResponse = s.createUser(request)
	if errResponse != nil {
		return nil, errResponse
	}

	token, errResponse = s.generateJwtToken(user)
	if errResponse != nil {
		return nil, errResponse
	}

	meResource, errResponse = s.makeMeResource(user)
	if errResponse != nil {
		return nil, errResponse
	}

	response = &resources.Login{
		Token: token,
		User:  *meResource,
	}

	return response, nil
}

func (s *authService) Me(userUuid string) (*resources.Me, *merr.ResponseError) {
	var response *resources.Me
	var err error
	var errResponse *merr.ResponseError
	var user *models.User

	user, err = s.userRepo.FindByUuid(userUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrUserNotFound)
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	response, errResponse = s.makeMeResource(user)
	if errResponse != nil {
		return nil, errResponse
	}

	return response, nil
}

func (s *authService) Availability(request *requests.Availability) (bool, *merr.ResponseError) {
	switch request.Type {
	case "email":
		return s.isAvailable(request.Value, func(value string) (any, error) {
			return s.userRepo.FindByEmail(value)
		})
	case "cnpj":
		return s.isAvailable(request.Value, func(value string) (any, error) {
			return s.agentRepo.FindByCnpj(value)
		})
	case "ccee":
		return s.isAvailable(request.Value, func(value string) (any, error) {
			return s.agentRepo.FindByCceeCode(value)
		})
	default:
		return false, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidAvailabilityType)
	}
}

func (s *authService) RefreshToken(oldToken string) (string, *merr.ResponseError) {
	newToken, err := s.jwtService.RefreshToken(oldToken)
	if err != nil {
		return "", merr.NewResponseError(http.StatusInternalServerError, ErrFailedToGenerateToken)
	}
	return newToken, nil
}

func (s *authService) generateJwtToken(user *models.User) (string, *merr.ResponseError) {
	token, err := s.jwtService.GenerateToken(user.Uuid, user.Email, user.UserType.Type)
	if err != nil {
		return "", merr.NewResponseError(http.StatusInternalServerError, ErrFailedToGenerateToken)
	}
	return token, nil
}

func (s *authService) isAvailable(value string, findFunc func(string) (any, error)) (bool, *merr.ResponseError) {
	var exists any
	var err error

	exists, err = findFunc(value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if exists != nil {
		return false, nil
	}

	return true, nil
}

func (s *authService) createUser(request *requests.SignUp) (*models.User, *merr.ResponseError) {
	var responseError *merr.ResponseError
	var user *models.User

	s.db.Transaction(func(tx *gorm.DB) error {
		var addressRepo = repository.NewAddressRepository(tx)
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
			responseError = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			return err
		}

		var submarketRepo = repository.NewSubmarketRepository(tx)
		submarketModel, err := submarketRepo.FindByName(request.Agent.SubmarketName)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				responseError = merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidSubmarket)
			} else {
				responseError = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			}
			return err
		}

		var agentRepo = repository.NewAgentRepository(tx)
		agentModel, err := agentRepo.Create(&models.Agent{
			Cnpj:        request.Agent.Cnpj,
			CompanyName: request.Agent.CompanyName,
			CceeCode:    request.Agent.CceeCode,
			SubmarketId: submarketModel.ID,
			AddressId:   addressModel.ID,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				responseError = merr.NewResponseError(http.StatusUnprocessableEntity, ErrAgentAlreadyExists)
			} else {
				responseError = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			}
			return err
		}

		var userTypeRepo = repository.NewUserTypeRepository(tx)
		userTypeModel, err := userTypeRepo.FindByName(request.UserType)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				responseError = merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidUserType)
			} else {
				responseError = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			}
			return err
		}

		var userRepo = repository.NewUserRepository(tx)
		user, err = userRepo.Create(repository.UserCreateParams{
			Uuid:     NewUuidV7String(),
			Name:     request.Name,
			Email:    request.Email,
			Password: Hash256String(request.Password),
			UserType: userTypeModel,
			Agent:    agentModel,
		})

		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				responseError = merr.NewResponseError(http.StatusUnprocessableEntity, ErrUserAlreadyExists)
			} else {
				responseError = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			}
			return err
		}

		return nil
	})

	if responseError != nil {
		return nil, responseError
	}

	return user, nil
}
