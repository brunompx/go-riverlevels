package database

import (
	"fmt"
	"log"

	"github.com/brunompx/go-riverlevels/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDatabase returns a database instance.
func GetDatabase() *gorm.DB {

	//user := envvar.DBUser()
	//password := envvar.DBPassword()
	//dbname := envvar.DBName()
	//dbhost := envvar.DBHost()
	//dbport := envvar.DBPort()

	//dsn := fmt.Sprintf(
	//	"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
	//	user, password, dbname, dbhost, dbport)
	//db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db.AutoMigrate(&model.User{})

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "postgres", "riverlevels", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	db.AutoMigrate(&model.Forecast{})
	db.AutoMigrate(&model.ForecastSet{})
	db.AutoMigrate(&model.ForecastLevel{})
	db.AutoMigrate(&model.Measure{})
	db.AutoMigrate(&model.MeasureLevel{})

	return db
}

func initStorage(db *gorm.DB) {
	genericDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	pingErr := genericDB.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database!")
}
