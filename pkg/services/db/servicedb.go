package db

import (
	"ecommerce/pkg/services/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var initialized bool

var (
	db *gorm.DB
)

// Initialize obtiene los datos de configFile
//Inicializa la DB
func Initialize(host, port, userName, password, databaseName string) {

	log.Info("Inicializando Base de Datos")

	if initialized { //TODO: hacer chequeo de db por nil
		log.Panic("La Base de Datos ya fue inicializada")
	}

	initialized = true

	dsn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName + "?parseTime=true"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}) // TODO: ver como recibir la conexion
	if err != nil {
		log.Panic(err.Error())
	}

	db = d

	log.Info("Base de Datos inicializada")
}

// GetDB obtiene la base de datos
func GetDB() *gorm.DB {
	return db
}
