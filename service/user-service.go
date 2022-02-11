package service

import (
	"log"

	"github.com/Kontribute/kontribute-web-backend/dto"
	"github.com/Kontribute/kontribute-web-backend/entity"
	repository "github.com/Kontribute/kontribute-web-backend/repository"
	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	CheckEmailInDb(email string) entity.User
	VerifyEmail(email string, otp string) entity.User
	CreateWebUser(user dto.UserWebRegisterDTO) entity.User
	SendWebOTP(email string) error
}

type userService struct {
	userRepository repository.UserRepository
}

//newUserService creates a new insatnace of userservice
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) CheckEmailInDb(email string) entity.User {
	//var user entity.User = service.userRepository.FindByEmail(email)
	return service.userRepository.FindByEmail(email)
}

func (service *userService) VerifyEmail(email string, otp string) entity.User {
	user := service.userRepository.VerifyOTP(email, otp)
	if user.Email != "" {
		return user
	}
	return entity.User{}
}

func (service *userService) CreateWebUser(user dto.UserWebRegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertWebUser(userToCreate)
	return res
}

func (service *userService) SendWebOTP(email string) error {
	return service.userRepository.SendOTP(email)
}
