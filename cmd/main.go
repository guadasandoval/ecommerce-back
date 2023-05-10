package main

import (
	"ecommerce/pkg/libs/config"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/keyvalue"
	"ecommerce/pkg/services/log"
	"ecommerce/router"
	"os"
	"go.uber.org/zap"
)

func main() {

	configFile := os.Getenv("CONFIGFILE")

	if configFile == "" {
		panic("CONFIGFILE not defined")
	}

	// Read Config
	cfg := config.NewConfig()
	err := cfg.ReadConfigFile(configFile)
	if err != nil {
		panic(err)
	}

	cfgData := cfg.Data()
	name := cfg.Name()
	version := cfg.Version()

	cfgLogger := cfgData.GetParamAsData("log")
	inicializarLogger(cfgLogger)

	cfgDb := cfgData.GetParamAsData("db")
	inicializarBD(cfgDb)

	cfgRedis := cfgData.GetParamAsData("redis")
	inicializarRedis(cfgRedis)

	cfgServer := cfgData.GetParamAsData("server")
	inicializarRouter(cfgServer)

	log.Info("Server started", zap.String("name", name), zap.String("version", version))
}

func inicializarLogger(config *config.Data) {
	devMode := config.GetParamAsBool("devMode")
	log.Initialize(devMode)
}

func inicializarBD(config *config.Data) {

	dbHost := config.GetParamAsString("host")
	dbPort := config.GetParamAsString("port")
	dbUserName := config.GetParamAsString("username")
	dbPassword := config.GetParamAsString("password")
	dbDatabase := config.GetParamAsString("database")

	db.Initialize(dbHost, dbPort, dbUserName, dbPassword, dbDatabase)
}

func inicializarRedis(config *config.Data) {

	redisHost := config.GetParamAsString("host")
	redisPort := config.GetParamAsInt("port")
	redisDatabase := config.GetParamAsInt("database")
	redisPassword := config.GetParamAsString("password")
	keyvalue.Initialize(redisHost, redisPort, redisDatabase, redisPassword)
}

func inicializarRouter(config *config.Data) {
	puerto := config.GetParamAsString("port")

	router.Initialize(puerto)
}
