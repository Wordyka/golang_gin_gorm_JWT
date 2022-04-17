package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/dto"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/entity"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/repository"
	"golang.org/x/crypto/bcrypt"
)

// digunakan untuk mengerjakan task tertentu yang berhubungan langsung dengan layer presentation (interface/front end client).
	

// interface AuthService sebagai contract mengenai service yang dibuat pada autentikasi
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

// struct authService untuk menginstansiasi interface dari UserRepository dari package repository
type authService struct {
	userRepository repository.UserRepository
}


// fungsi NewAuthService untuk memanggil struct authService dengan menginisialisasi variabel address userRepository
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}


// fungsi VerifiyCredential sebagai verifikasi email dan password pada userRepository
func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparePassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparePassword {
			return res
		}
		return false
	}
	return res
}


// fungsi CreateUser untuk membuat user baru
func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

// fungsi FindByEmail untuk mencaru user melalui email user tersebut
func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

// fungsi isDuplicateEmail untuk mengecek email user tersebut apakah duplicate dengan user lain
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

// fungsi comparePassword untuk membandingkan password yang dienkripsi dengan inputan user
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false
	}
	return true
}




