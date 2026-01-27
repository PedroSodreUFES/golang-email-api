package controller

import (
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
	users := router.Group("/user")
	{
		users.DELETE("/me", c.DeleteSelf)
		users.GET("/me", c.GetSelf)
		users.GET("/:id/email", c.GetUserByEmail)
		users.POST("", c.SignUp)
		users.POST("/login", c.Login)
		users.PUT("/photo", c.UpdateUserPhoto)
	}
}

func (c *UserController) GetSelf(ctx *gin.Context) {

}

func (c *UserController) GetUserByEmail(ctx *gin.Context) {

}

func (c *UserController) SignUp(ctx *gin.Context) {

}

func (c *UserController) UpdateUserPhoto(ctx *gin.Context) {

}

func (c *UserController) DeleteSelf(ctx *gin.Context) {

}

func (c *UserController) Login(ctx *gin.Context) {

}
