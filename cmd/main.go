package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	api "ims-staff-api"
	"ims-staff-api/pkg/handler"
	"ims-staff-api/pkg/repository"
	"ims-staff-api/pkg/service"
)

//	@title			ims-staff-api
//	@version		1.0
//	@description	API server for IMS application

//	@host		localhost:8000
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//  @in header
//  @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Это можно убрать, если вы не используете .env файл
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"), // Исправлено на viper
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(api.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error of init server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml") // Убедитесь, что ваш файл конфигурации имеет правильное расширение
	viper.AutomaticEnv()

	// Привязка переменных окружения к viper ключам
	viper.BindEnv("db.host", "DB_HOST")
	viper.BindEnv("db.port", "DB_PORT")
	viper.BindEnv("db.username", "DB_USER")
	viper.BindEnv("db.dbname", "DB_NAME")
	viper.BindEnv("db.sslmode", "DB_SSLMODE")
	viper.BindEnv("db.password", "DB_PASSWORD")
	viper.BindEnv("port", "PORT")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
