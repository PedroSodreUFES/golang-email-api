package controller

import (
	"errors"
	"main/internal/exceptions"
	"main/internal/jsonutils"
	"main/internal/middlewares"
	"main/internal/users/DTO/requests"
	"main/internal/users/models"
	"net/http"
	"path/filepath"
	"strings"

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

	user, err := c.userService.GetMe(ctx.Request.Context(), userID)
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

	// pegar arquivo
	fileHeader, err := ctx.FormFile("photo") // nome do campo no form-data
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "photo is required"})
		return
	}

	// validar arquivo
	// -- extensão
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".webp" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type"})
		return
	}
	// -- tamanho
	const maxSize = 8 << 20 // 8 MB
	if fileHeader.Size > maxSize {
		ctx.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file too large (max 8MB)"})
		return
	}

	// service + resposta
	if err := c.userService.UpdateUserPhoto(ctx.Request.Context(), userID, fileHeader); err != nil {
		if errors.Is(err, exceptions.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": exceptions.ErrUserNotFound.Error()})
			return
		}
		if errors.Is(err, exceptions.ErrTimeoutExceeded) {
			ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": exceptions.ErrTimeoutExceeded.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
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

	err := c.userService.DeleteUser(ctx.Request.Context(), userID)
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
