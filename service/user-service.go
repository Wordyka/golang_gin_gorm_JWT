package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/dto"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/entity"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/repository"
)

// interface UserService sebagai contract mengenai service yang dibuat pada User
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

// struct UserService untuk menginstansiasi interface dari UserRepository dari package repository
type userService struct {
	userRepository repository.UserRepository
}

// fungsi NewAuthService untuk memanggil struct authService dengan menginisialisasi variabel address userRepository
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

// Implementasi interface pada method Update untuk mengubah data user
func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updateUser := service.userRepository.UpdateUser(userToUpdate)
	return updateUser
}


// Implementasi interface pada method Profile untuk menampilkan semua data mengenai user tersebut
func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
