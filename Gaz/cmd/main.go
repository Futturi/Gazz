package main

import (
	"fmt"
	"github.com/Futturi/Gaz/internal/handler"
	"github.com/Futturi/Gaz/internal/repo"
	"github.com/Futturi/Gaz/internal/service"
	"github.com/Futturi/Gaz/pkg"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"time"
)

// @title Notification app
// @version 1.0
// @description API Server 4 Notification Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
var (
	port = os.Getenv("APP_PORT")
)

func main() {
	fmt.Println(time.Now())
	logg := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logg)
	err := InitCfg()
	if err != nil {
		slog.Error("error with config", "error", err)
		return
	}
	cfg := pkg.NewConfig(viper.GetString("host"),
		viper.GetString("port"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("dbname"),
		viper.GetString("sslmode"))
	db, err := pkg.Init(cfg)
	if err != nil {
		slog.Error("error with initializing db", "error", err)
		return
	}
	err = pkg.Migrat(viper.GetString("host"), viper.GetString("port"), viper.GetString("dbname"), viper.GetString("username"), viper.GetString("password"))
	if err != nil {
		slog.Info("error", "err", err)
	}
	c := cron.New()
	repo := repo.NewRepository(db)
	service := service.NewSerivce(repo)
	handler := handler.NewHandler(service, c)

	defer c.Stop()
	if err := handler.Init().Listen(":" + port); err != nil {
		slog.Error("error with initializing server", "error", err)
	}
}

func InitCfg() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}
