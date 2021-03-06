package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/dto"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/helper"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/service"
)

// pembuatan interface UserController yang berisi method dengan parameter berupa context HTTP web framework Gin
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

// struct UserController yang berisi service dari UserService dn JWTService
type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserController is creating new instance of UserController
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

// implementasi interface pada method Update untuk mengubah data user sesuai token yang di passing
func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}


// implementasi interface pada method Profile untuk mendapatkan data user itu sendiri 
func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
