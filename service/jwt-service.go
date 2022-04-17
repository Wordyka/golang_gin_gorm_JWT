package service

import (
	"fmt"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// interface JWTService sebagai contract mengenai service yang dibuat
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

// struct jwtCustomClaim yang berisi variabel yang digunakan untuk mengklaim jwt berdasarkan user_id 
type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// struct jwtService yang berisi variabel untuk menyimpan secretKey dan issuer
type jwtService struct {
	secretKey string
	issuer    string
}

//fungsi NewJWTService untuk membuat instansiasi pada JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "rudychandra",
		secretKey: getSecretKey(),
	}
}

// fungsi getSecretKey untuk mendapatkan secret key berupa string 
func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "rudychandra"
	}
	return secretKey
}


// fungsi GenerateToken untuk mendapatkan / mengenerate token berdasarkan user_id nya 
func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}


// fungsi ValidateToken untuk memvalidasi token berdasarkan token dari jwt
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
