package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/auth"
	"main/internal/users/controller"
	"main/internal/users/repositories"
	"main/internal/users/service"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOEMAIL_DATABASE_USER"),
		os.Getenv("GOEMAIL_DATABASE_PASSWORD"),
		os.Getenv("GOEMAIL_DATABASE_HOST"),
		os.Getenv("GOEMAIL_DATABASE_PORT"),
		os.Getenv("GOEMAIL_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	secret := os.Getenv("GOEMAIL_JWT_KEY")
	jwtMaker := auth.JWTMaker{
		Secret: []byte(secret),
		Duration: time.Hour * 2,
	}
	usersRepository := repositories.NewPostgreUserRepository(pool)
	userService := service.NewUserService(usersRepository, jwtMaker)
	userController := controller.NewUserController(userService, jwtMaker.Secret)

	userController.RegisterRoutes(router)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}