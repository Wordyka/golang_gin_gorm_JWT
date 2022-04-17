package repository

import (
	"log"

	"github.com/Wordyka/golang_gin_gorm_JWT.git/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// interface UserRepository yang berisi method dengan parameter berupa nilai dari entity User
type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
}

// struct userConnection yang berisi koneksi dari database
type userConnection struct {
	connection *gorm.DB
}

// fungsi NewUserRepository untuk membuat repository baru pada database yang telah terkoneksi
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

// Implementasi interface pada method InsertUser untuk menambahkan data user baru
func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}


// Implementasi interface pada method UpdateUser untuk mengubah data user berdasarkan id nya
// password dienkripsi dengan mengimplementasikan method hashAndSalt
func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}


// Implementasi interface pada method VerifyCredential untuk memverifikasi email dan password user
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

// Implementasi interface pada method IsDuplicateEmail untuk mengecek apakah email yang diinput user sama dengan user lainnya
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

// Implementasi interface pada method FindByEmail untuk mencari email dari user tersebut
func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

// Implementasi interface pada method ProfileUser untuk mendapatkan data user tersebut
func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}


// method hashAndSalt untuk mengenkripsi password user yang diinput dengan hash bcrypt.MinCost
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Print(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}




