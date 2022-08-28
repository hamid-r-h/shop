package db

import (
	"log"
	"os"

	"shop/src/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "host=pg user=hamid password=heidari dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect ", err.Error())
		os.Exit(2)
	}
	log.Println("connected to the database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migration")
	db.SetupJoinTable(&models.User{}, "Products", &models.UserProduct{})
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Comment{})

	Database = DbInstance{Db: db}
}
