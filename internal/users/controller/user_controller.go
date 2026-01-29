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
	jwt_secret  []byte
}

func NewUserController(userService models.UserService, jwt_secret []byte) *UserController {
	return &UserController{
		userService: userService,
		jwt_secret:  jwt_secret,
	}
}

func (c *UserController) RegisterRoutes(router *gin.Engine) {
	public := router.Group("/user")
	{
		public.POST("", c.signUp)
		public.POST("/login", c.login)
	}
	private := router.Group("/user")
	private.Use(middlewares.JWTMiddleware(c.jwt_secret))
	{
		private.DELETE("/me", c.deleteSelf)
		private.GET("/me", c.getSelf)
		private.PUT("/me/photo", c.updateUserPhoto)
	}
}

func (c *UserController) getSelf(ctx *gin.Context) {
	val, ok := ctx.Get(middlewares.IDKey)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	userID, ok := val.(int32)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	user, err := c.userService.GetMe(ctx, userID)
	if err != nil {
		if errors.Is(err, exceptions.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": exceptions.ErrUserNotFound.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
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
	val, ok := ctx.Get(middlewares.IDKey)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	userID, ok := val.(int32)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	err := c.userService.DeleteUser(ctx, userID)
	if err != nil {
		if errors.Is(err, exceptions.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": exceptions.ErrUserNotFound.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UserController) login(ctx *gin.Context) {
	var body requests.LoginRequest

	body, problems, err := jsonutils.DecodeValidJson[requests.LoginRequest](ctx.Request)
	if err != nil {
		if errors.Is(err, jsonutils.ErrFailedToDecodeJson) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": exceptions.ErrUnproccessableEntity.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, problems)
		return
	}

	response, err := c.userService.AuthenticateUser(ctx.Request.Context(), &body)
	if err != nil {
		// Erro de usuário não encontrado ou credencial não bate
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": exceptions.ErrInvalidCredentials.Error()})
			return
		}
		// Algum outro erro
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
