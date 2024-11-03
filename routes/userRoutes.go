package routes

import (
	"github.com/gglzc/mqTest/controller"
	"github.com/gin-gonic/gin"
)
// jwt 程式完成在放入
//, jwtService service.JWTService

func UserRoutes(route *gin.Engine, userController controller.UserController) {
	routes := route.Group("/api/user")
	{
		// User
		routes.POST("updateBalance", userController.UpdateBalance)
		routes.GET("", userController.GetUser)
		// routes.POST("/login", userController.Login)
		// routes.DELETE("", middleware.Authenticate(jwtService), userController.Delete)
		// routes.PATCH("", middleware.Authenticate(jwtService), userController.Update)
		// routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		// routes.POST("/verify_email", userController.VerifyEmail)
		// routes.POST("/send_verification_email", userController.SendVerificationEmail)
	}
}