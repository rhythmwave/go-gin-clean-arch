// config/database.go
package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var db *gorm.DB
	var err error

	dbDriver := GetEnv("DB_DRIVER", "")
	dbUser := GetEnv("DB_USER", "")
	dbPassword := GetEnv("DB_PASSWORD", "")
	dbHost := GetEnv("DB_HOST", "")
	dbPort := GetEnv("DB_PORT", "")
	dbName := GetEnv("DB_NAME", "")

	switch dbDriver {
	case "mysql":
		dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mssql":
		dsn := "sqlserver://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "?database=" + dbName
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	default:
		log.Fatalf("Unsupported DB driver: %s", dbDriver)
	}

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Auto migrate the User entity
	// db.AutoMigrate(&entities.User{})

	return db
}
