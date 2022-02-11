package controller

import (
	"fmt"
	"net/http"
	"strconv"

	dto "github.com/Kontribute/kontribute-web-backend/dto"
	"github.com/Kontribute/kontribute-web-backend/entity"
	"github.com/Kontribute/kontribute-web-backend/helper"
	"github.com/Kontribute/kontribute-web-backend/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//UserController is a ....
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	CheckEmailInDb(context *gin.Context)
	CreateUserFromWeb(context *gin.Context)
	SendOTP(context *gin.Context)
	VerifyOTP(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) SendOTP(context *gin.Context) {
	var SendWebOTPDTO dto.SendWebOTPDTO
	errDTO := context.ShouldBind(&SendWebOTPDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err := c.userService.SendWebOTP(SendWebOTPDTO.Email)
	if err != nil {
		response := helper.BuildResponse(true, "OK", err)
		context.JSON(http.StatusCreated, response)
		return
	}
	response := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusCreated, response)
}

func (c *userController) CreateUserFromWeb(context *gin.Context) {
	var UserFromWebDTO dto.UserWebRegisterDTO
	errDTO := context.ShouldBind(&UserFromWebDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	result := c.userService.CreateWebUser(UserFromWebDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *userController) VerifyOTP(context *gin.Context) {
	var UserWebVerifyDTO dto.UserWebVerify
	errDTO := context.ShouldBind(&UserWebVerifyDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	user := c.userService.VerifyEmail(UserWebVerifyDTO.Email, UserWebVerifyDTO.OTP)
	if user.Email != "" {
		res := helper.BuildResponse(true, "OK", user)
		context.JSON(http.StatusOK, res)
		return
	}
	res := helper.BuildErrorResponse("invalid otp", errDTO.Error(), helper.EmptyObj{})
	context.JSON(http.StatusBadRequest, res)

}

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

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}

func (c *userController) CheckEmailInDb(context *gin.Context) {
	var UserCheckDTO dto.UserCheckDTO
	errDTO := context.ShouldBind(&UserCheckDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		var user entity.User = c.userService.CheckEmailInDb(UserCheckDTO.Email)
		if (user == entity.User{}) {
			res := helper.BuildErrorResponse("User not found", "No data with given Email, But An OTP as been sent to your Email", helper.EmptyObj{})
			context.JSON(http.StatusNotFound, res)
		} else {
			res := helper.BuildResponse(true, "OK", user)
			context.JSON(http.StatusOK, res)
		}
	}
}
