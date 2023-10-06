package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TitusW/game-service/config"
	auth_handler "github.com/TitusW/game-service/internal/handler/auth"
	ledger_handler "github.com/TitusW/game-service/internal/handler/ledger"
	user_handler "github.com/TitusW/game-service/internal/handler/user"
	userbankaccount_handler "github.com/TitusW/game-service/internal/handler/user-bank-account"
	"github.com/TitusW/game-service/internal/middleware"
	ledger_repo "github.com/TitusW/game-service/internal/repo/ledger"
	token_repo "github.com/TitusW/game-service/internal/repo/token"
	user_repo "github.com/TitusW/game-service/internal/repo/user"
	userbankaccount_repo "github.com/TitusW/game-service/internal/repo/user-bank-account"
	wallet_repo "github.com/TitusW/game-service/internal/repo/wallet"
	auth_usecase "github.com/TitusW/game-service/internal/usecase/auth"
	ledger_usecase "github.com/TitusW/game-service/internal/usecase/ledger"
	user_usecase "github.com/TitusW/game-service/internal/usecase/user"
	userbankaccount_usecase "github.com/TitusW/game-service/internal/usecase/user-bank-account"
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
			TablePrefix: "game.",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	mainCfg := config.InitializeConfig()

	router := gin.Default()

	db := setupDB(mainCfg)
	rdb := setupRedis(mainCfg)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Healthy!",
		})
	})

	walletRepo := wallet_repo.New(db)

	userBankAccountRepo := userbankaccount_repo.New(db)
	userBankAccountUC := userbankaccount_usecase.New((userBankAccountRepo))
	userBankAccountHandler := userbankaccount_handler.New(userBankAccountUC)

	ledgerRepo := ledger_repo.New(db)
	ledgerUsecase := ledger_usecase.New(ledgerRepo, walletRepo, db)
	ledgerHandler := ledger_handler.New(ledgerUsecase)

	userRepo := user_repo.New(db)
	userUsecase := user_usecase.New(userRepo, walletRepo, userBankAccountRepo, db)
	userHandler := user_handler.New(userUsecase)

	tokenRepo := token_repo.New(rdb)
	authUsecase := auth_usecase.New(userRepo, tokenRepo)
	authHandler := auth_handler.New(authUsecase)

	router.POST("/users/register", userHandler.Register)
	router.GET("/users/", middleware.Authentication(), userHandler.GetUsers)
	router.GET("/users/:ksuid", middleware.Authentication(), userHandler.GetUserDetails)

	router.POST("/user-bank-accounts/register", middleware.Authentication(), userBankAccountHandler.Register)

	router.POST("/ledgers/topup", middleware.Authentication(), ledgerHandler.Topup)

	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/logout", middleware.Authentication(), authHandler.Logout)

	router.Run()
}
