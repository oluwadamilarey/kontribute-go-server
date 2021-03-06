package main

import (
	config "github.com/Kontribute/kontribute-web-backend/config"
	"github.com/Kontribute/kontribute-web-backend/controller"
	"github.com/Kontribute/kontribute-web-backend/middleware"
	repository "github.com/Kontribute/kontribute-web-backend/repository"
	"github.com/Kontribute/kontribute-web-backend/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//go:generate sqlboiler --wipe mysql

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)

	jwtService  service.JWTService  = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepository)

	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	userWebRoutes := r.Group("api/web")
	{
		userWebRoutes.POST("/user", userController.CheckEmailInDb)
		userWebRoutes.POST("/sendotp", userController.SendOTP)
		userWebRoutes.POST("/verifyotp", userController.VerifyOTP)
		userWebRoutes.POST("/register", userController.CreateUserFromWeb)
	}
	r.Run()
}
