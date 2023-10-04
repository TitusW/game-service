package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TitusW/game-service/config"
	auth_handler "github.com/TitusW/game-service/internal/handler/auth"
	user_handler "github.com/TitusW/game-service/internal/handler/user"
	token_repo "github.com/TitusW/game-service/internal/repo/token"
	user_repo "github.com/TitusW/game-service/internal/repo/user"
	auth_usecase "github.com/TitusW/game-service/internal/usecase/auth"
	user_usecase "github.com/TitusW/game-service/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func setupRedis(configData config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:  configData.Redis.Network,
		Addr:     configData.Redis.Addr,
		Password: configData.Redis.Password,
		DB:       configData.Redis.DB,
	})

	return rdb
}

func setupDB(configData config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		configData.Database.Host,
		configData.Database.Username,
		configData.Database.Password,
		configData.Database.Name,
		configData.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "api.",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	mainCfg := config.InitializeConfig()

	fmt.Println("Main Config: ", mainCfg)

	router := gin.Default()

	db := setupDB(mainCfg)
	rdb := setupRedis(mainCfg)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Healthy!",
		})
	})

	userRepo := user_repo.New(db)
	userUsecase := user_usecase.New(userRepo)
	userHandler := user_handler.New(userUsecase)

	tokenRepo := token_repo.New(rdb)
	authUsecase := auth_usecase.New(userRepo, tokenRepo)
	authHandler := auth_handler.New(authUsecase)

	router.POST("/users/register", userHandler.Register)
	router.GET("/users/")
	router.GET("/users/:ksuid")

	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/logout", authHandler.Logout)

	router.Run()
}
