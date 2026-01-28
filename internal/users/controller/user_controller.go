package controller

import (
	"main/internal/exceptions"
	"main/internal/middlewares"
	"main/internal/users/DTO/requests"
	"main/internal/users/models"
	"net/http"

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
		public.POST("", c.signUp)
		public.POST("/login", c.login)
	}
	private := router.Group("/user")
	private.Use(middlewares.JWTMiddleware())
	{
		private.DELETE("/me", c.deleteSelf)
		private.GET("/me", c.getSelf)
		private.PUT("/photo", c.updateUserPhoto)
	}
}

func (c *UserController) getSelf(ctx *gin.Context) {
	
}

func (c *UserController) signUp(ctx *gin.Context) {
	var request requests.CreateUserRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": exceptions.ErrUnproccessableEntity.Error()})
		return
	}

	user, err := c.userService.SignUp(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) updateUserPhoto(ctx *gin.Context) {

}

func (c *UserController) deleteSelf(ctx *gin.Context) {

}

func (c *UserController) login(ctx *gin.Context) {

}
