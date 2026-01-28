package controller

import (
	"errors"
	"main/internal/exceptions"
	"main/internal/jsonutils"
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
	var body requests.CreateUserRequest

	body, problems, err := jsonutils.DecodeValidJson[requests.CreateUserRequest](ctx.Request)
	if err != nil {
		if errors.Is(err, jsonutils.ErrFailedToDecodeJson) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": exceptions.ErrUnproccessableEntity.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, problems)
		return
	}

	response, err := c.userService.SignUp(ctx.Request.Context(), &body)
	if err != nil {
		// Erro de email não único
		if errors.Is(err, exceptions.ErrEmailShouldBeUnique) {
			ctx.JSON(http.StatusConflict, gin.H{"error": exceptions.ErrEmailShouldBeUnique.Error()})
			return
		}
		// Algum outro erro
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *UserController) updateUserPhoto(ctx *gin.Context) {

}

func (c *UserController) deleteSelf(ctx *gin.Context) {

}

func (c *UserController) login(ctx *gin.Context) {

}
