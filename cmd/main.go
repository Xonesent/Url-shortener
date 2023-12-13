package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	server "url-shortener"
	"url-shortener/pkg/handler"
	postgres "url-shortener/pkg/repository"
	"url-shortener/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := Init_Config(); err != nil {
		logrus.Fatalf("config error - %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("godotenv error - %s", err.Error())
	}

	db, err := postgres.New_Postgres(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("db initialize error - %s", err.Error())
	}

	repos := postgres.New_Repository(db)
	services := service.New_Service(repos)
	handlers := handler.New_Handler(services)

	server := new(server.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.Init()); err != nil {
			logrus.Fatalf("http server error - %s", err.Error())
		}
	}()

	logrus.Print("Url-shortener Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Url-shortener Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("db connection close: %s", err.Error())
	}
}

func Init_Config() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
