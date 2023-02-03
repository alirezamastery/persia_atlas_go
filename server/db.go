package server

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"persia_atlas/server/models"
)

func (server *Server) connectDatabase() {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	var err error
	server.DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		CreateBatchSize: 1000,
		Logger:          logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("error in connecting to database:", err)
		return
	} else {
		log.Println("connected to database")
	}

	sqlDB, err := server.DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func (server *Server) migrateDatabase() {
	err := server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Profile{},

		&models.Brand{},
		&models.ActualProduct{},
		&models.VariantSelectorType{},
		&models.ProductType{},
		&models.Product{},
		&models.VariantSelector{},
		&models.Variant{},

		&models.CostType{},
		&models.Cost{},
		&models.Income{},
		&models.ProductCost{},
		&models.Invoice{},
		&models.InvoiceItem{},
	)
	if err != nil {
		log.Fatalln("error in db migration:", err.Error())
		return
	}
}
