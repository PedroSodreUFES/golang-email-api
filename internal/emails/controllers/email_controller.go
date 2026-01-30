package controller

import (
	"errors"
	"main/internal/emails/DTO/requests"
	"main/internal/emails/models"
	"main/internal/exceptions"
	"main/internal/jsonutils"
	"main/internal/middlewares"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmailController struct {
	emailService models.EmailService
	jwt_secret   []byte
}

func NewEmailController(emailService models.EmailService, jwt_secret []byte) *EmailController {
	return &EmailController{
		emailService: emailService,
		jwt_secret:   jwt_secret,
	}
}

func (c *EmailController) RegisterRoutes(router *gin.Engine) {
	private := router.Group("/emails")
	private.Use(middlewares.JWTMiddleware(c.jwt_secret))
	{
		private.DELETE("/:id", c.deleteEmail)
		private.POST("/", c.sendEmail)
		private.GET("/sent", c.getMySentEmails)
		private.GET("/received", c.getMyReceivedEmails)
		private.GET("/:id", c.getEmailById)
	}
}

func (c *EmailController) deleteEmail(ctx *gin.Context) {
	val, ok := ctx.Get(middlewares.IDKey)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}
	userID, ok := val.(int32)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}

	emailIdParam := ctx.Param("id")
	emailId, err := strconv.ParseInt(emailIdParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email id"})
		return
	}

	err = c.emailService.DeleteEmail(ctx.Request.Context(), int32(emailId), userID)
	if err != nil {
		if errors.Is(err, exceptions.ErrEmailNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": exceptions.ErrEmailNotFound.Error()})
			return
		}
		if errors.Is(err, exceptions.ErrNotAllowed) {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": exceptions.ErrNotAllowed.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *EmailController) sendEmail(ctx *gin.Context) {
	val, ok := ctx.Get(middlewares.IDKey)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}
	userID, ok := val.(int32)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}

	var body requests.SendEmailRequest
	body, problems, err := jsonutils.DecodeValidJson[requests.SendEmailRequest](ctx.Request)
	if err != nil {
		if errors.Is(err, jsonutils.ErrFailedToDecodeJson) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": exceptions.ErrUnproccessableEntity.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, problems)
		return
	}

	email, err := c.emailService.SendEmail(ctx.Request.Context(), &body, userID)
	if err != nil {
		if errors.Is(err, exceptions.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": exceptions.ErrUserNotFound.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, email)
}

func (c *EmailController) getMyReceivedEmails(ctx *gin.Context) {
	val, ok := ctx.Get(middlewares.IDKey)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}
	userID, ok := val.(int32)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}

	myReceivedEmails, err := c.emailService.GetMyReceivedEmails(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, myReceivedEmails)
}

func (c *EmailController) getMySentEmails(ctx *gin.Context) {
	val, ok := ctx.Get(middlewares.IDKey)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}
	userID, ok := val.(int32)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}

	mySentEmails, err := c.emailService.GetMySentEmails(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mySentEmails)
}

func (c *EmailController) getEmailById(ctx *gin.Context) {
	val, ok := ctx.Get(middlewares.IDKey)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}
	userID, ok := val.(int32)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": exceptions.ErrUnauthorized.Error()})
		return
	}

	emailIdParam := ctx.Param("id")
	emailId, err := strconv.ParseInt(emailIdParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email id"})
		return
	}

	email, err := c.emailService.GetEmailById(ctx.Request.Context(), userID, int32(emailId))
	if err != nil {
		if errors.Is(err, exceptions.ErrEmailNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, exceptions.ErrNotAllowed) {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": exceptions.ErrInternalServerError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, email)
}
