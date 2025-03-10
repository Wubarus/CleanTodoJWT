package main

import (
	"CTodo/internal/adapter/handlers"
	"CTodo/internal/adapter/repo"
	"CTodo/internal/adapter/routes"
	"CTodo/internal/config"
	"CTodo/internal/core/domain"
	"CTodo/internal/core/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	userSrv *services.UserService
)

func main() {
	var db *gorm.DB
	var err error

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := config.InitConfig()

	// DB depends on env variable in config
	switch cfg.Env {
	// case "test":
	case "dev":
		db, err = SqliteInit(cfg)
	case "prod":
		db, err = PostgresInit(cfg)
	default:
		panic("Config Env not matched")
	}

	if err != nil {
		panic(err.Error())
	}

	//TODO: init log & cover project with it

	// Data migration to DB
	err = db.AutoMigrate(&domain.Task{}, &domain.User{}, &domain.TaskList{})
	if err != nil {
		panic(err.Error())
	}

	// DB injection to storage
	storage := repo.NewStorage(db)

	router := gin.Default()

	// init user Service, Handler, Router with dep injection
	//TODO: create supergroup with users, separated auth, tasks
	userService := services.NewUserService(storage)
	userHandler := handlers.NewUserHandler(userService)
	routes.NewUserRouter(router, userHandler)

	// server init
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	log.Println("Starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

func PostgresInit(cfg *config.Config) (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SqliteInit(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Storage), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
