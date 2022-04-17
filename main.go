package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/config"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/controller"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/middleware"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/repository"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/service"
	"gorm.io/gorm"
)

// deklarasi variabel dengan meng-assign valuenya berupa instansiasi file pada package nya
var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

// fungsi main() sebagai entry point dalam menjalankan server
func main() {
	// melakukan delay ketika eksekusi method CloseDatabaseConnection untuk menutup koneksi database
	defer config.CloseDatabaseConnection(db)	

	// instansiasi gin secara default pada variabel r
	r := gin.Default()

	// Route untuk melakukan autentikasi
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	// Route untuk mengakses profile dengan meng-authorisasi JWT nya terlebih dahulu
	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	// Route untuk mengakses data book dengan meng-authorisasi JWT nya terlebih dahulu
	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	// menjalankan gin
	r.Run()
}
