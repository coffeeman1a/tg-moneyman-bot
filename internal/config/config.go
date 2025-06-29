package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	C Config
)

type Config struct {
	Port     string
	DSN      string
	BotToken string
	Debug    bool
}

func Init() {
	_ = godotenv.Load() // reading env variables from file if exists

	// logger
	logLevel := getEnv("LOG_LEVEL", "debug")
	logFormat := getEnv("LOG_FORMAT", "text")
	initLogger(logFormat, logLevel)

	// app
	port := getEnv("PORT", "8080")

	// db
	dbPort := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER_NAME", "moneyman")
	pass := getEnv("DB_PASS", "password")
	host := getEnv("HOST", "localhost")
	dbname := getEnv("DB_NAME", "moneyman")

	// data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		user, pass, host, dbPort, dbname,
	)

	// telegram
	botToken := getEnv("TG_BOT_TOKEN", "")

	// global config
	C = Config{
		Port:     port,
		DSN:      dsn,
		BotToken: botToken,
		Debug:    false,
	}

	if logLevel == "debug" {
		C.Debug = true
	}
}

func initLogger(logFormat, logLevel string) {
	switch logFormat {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
		log.WithField("logFormat", logFormat).Info("Setting log format to json")
	case "text":
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		log.WithField("logFormat", logFormat).Info("Setting log format to text")
	default:
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		log.WithField("logFormat", logFormat).Info("Undefined log format variable, setting text format by default")

	}

	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.WithField("logLevel", logLevel).Info("Setting log level to debug")
	case "info":
		log.SetLevel(log.InfoLevel)
		log.WithField("logLevel", logLevel).Info("Setting log level to info")
	case "warn":
		log.SetLevel(log.WarnLevel)
		log.WithField("logLevel", logLevel).Warn("Setting log level to warn")
	default:
		log.SetLevel(log.DebugLevel)
		log.WithField("logLevel", logLevel).Info("Undefined log level variable, setting debug level by default")
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
