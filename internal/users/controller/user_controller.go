package controller

import (
	"main/internal/middlewares"
	"main/internal/users/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService models.UserService
}

func NewUserController(userService models.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) RegisterRoutes(router *gin.Engine) {
	public := router.Group("/user")
	{
		public.POST("", c.SignUp)
		public.POST("/login", c.Login)
	}
	private := router.Group("/user")
	private.Use(middlewares.JWTMiddleware())
	{
		private.DELETE("/me", c.DeleteSelf)
		private.GET("/me", c.GetSelf)
		private.PUT("/photo", c.UpdateUserPhoto)
	}
}

func (c *UserController) GetSelf(ctx *gin.Context) {

}

func (c *UserController) SignUp(ctx *gin.Context) {

}

func (c *UserController) UpdateUserPhoto(ctx *gin.Context) {

}

func (c *UserController) DeleteSelf(ctx *gin.Context) {

}

func (c *UserController) Login(ctx *gin.Context) {

}
