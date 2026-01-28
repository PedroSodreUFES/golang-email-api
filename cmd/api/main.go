package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/users/controller"
	"main/internal/users/repositories"
	"main/internal/users/service"
	"os"

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

	usersRepository := repositories.NewPostgreUserRepository(pool)
	userService := service.NewUserService(usersRepository)
	userController := controller.NewUserController(userService)

	userController.RegisterRoutes(router)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}