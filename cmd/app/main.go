package main

import (
	"CTodo/internal/adapter/handlers"
	"CTodo/internal/adapter/repo"
	"CTodo/internal/config"
	"CTodo/internal/core/domain"
	"CTodo/internal/core/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var (
	userSrv *services.UserService
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	//TODO: init config and use it to setup server
	cfg := config.InitConfig()

	//TODO: init log & cover project with it

	//SQLite DB to test
	db, err := gorm.Open(sqlite.Open(cfg.Storage), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to %v", err)
		os.Exit(1)
	}

	//TODO: use Postgres
	// FOR POSTGRES
	//host := os.Getenv("DB_HOST")
	//port := os.Getenv("DB_PORT")
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")
	//
	//conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//
	//db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	//if err != nil {
	//	fmt.Printf("Failed to connect to %v", err)
	//	os.Exit(1)
	//}

	db.AutoMigrate(&domain.Task{}, &domain.User{}, &domain.TaskList{})

	//DB injection to storage
	storage := repo.NewStorage(db)

	//storage injection to service
	userSrv = services.NewUserService(storage)

	//server init
	//TODO: take out server to internal
	InitRoutes()
}

func InitRoutes() {
	router := gin.Default()

	userHandler := handlers.NewUserHandler(*userSrv)

	//TODO: make routes to tasks
	router.POST("/auth/register", userHandler.Register)
	router.POST("/auth/login", userHandler.Login)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//TODO: log all events & errors
	log.Println("starting server...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
