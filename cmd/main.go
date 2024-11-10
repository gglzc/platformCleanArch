package main

import (
	"os"
	"time"

	"github.com/gglzc/mqTest/config"
	"github.com/gglzc/mqTest/config/kafka"
	"github.com/gglzc/mqTest/controller"
	"github.com/gglzc/mqTest/repository"
	"github.com/gglzc/mqTest/routes"
	"github.com/gglzc/mqTest/service"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main(){
	
	db :=config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)
	// kafka setting
	// consumer:=config.SetupKafkaComsumer();
	c:=kafka.SetupKafkaProducer()
	
	server:=gin.Default()
	logger , _:=zap.NewProduction()
	// logger中間件
	server.Use(ginzap.Ginzap(logger , time.RFC3339 ,true))
	//f
	var (
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService service.UserService = service.NewUserService(userRepository ,c )
		userController controller.UserController = controller.NewUserController(userService)
	)

	routes.UserRoutes(server,userController)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	
	var serve string
	
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		// log.Fatalf("error running server: %v", err)
	}
}