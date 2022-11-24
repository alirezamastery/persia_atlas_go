package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	brandcontroller "persia_atlas/server/controllers/brand"
	"persia_atlas/server/models"
	brandservice "persia_atlas/server/services/brand"
	"persia_atlas/server/websocket"
	"persia_atlas/server/websocket/commands"
	"testing"

	srv "persia_atlas/server"
)

var (
	server = srv.Server{}

	brandController = brandcontroller.BrandController{}
)

func TestMain(m *testing.M) {
	fmt.Println("Start Test")
	err := godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		msg := fmt.Sprintf("error in loading .env file: %s", err.Error())
		log.Fatalf(msg)
	} else {
		log.Println(".env loaded")
	}

	setupDatabase()
	migrateDatabase()
	setupWs()
	setupControllers()

	exitVal := m.Run()

	os.Exit(exitVal)
}

func GetNewRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func setupDatabase() {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_TEST_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	var err error
	server.DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	} else {
		fmt.Printf("We are connected to the database\n")
	}
}

func migrateDatabase() {
	err := server.DB.Migrator().DropTable(
		&models.User{},
		&models.Profile{},

		&models.Brand{},
		&models.ActualProduct{},
		&models.VariantSelectorType{},
		&models.ProductType{},
		&models.Product{},
		&models.VariantSelector{},
		&models.ProductVariant{},

		&models.CostType{},
		&models.Cost{},
		&models.Income{},
		&models.ProductCost{},
		&models.Invoice{},
		&models.InvoiceItem{},
	)
	if err != nil {
		log.Fatalln("error in dropping tables:", err.Error())
		return
	}

	err = server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Profile{},

		&models.Brand{},
		&models.ActualProduct{},
		&models.VariantSelectorType{},
		&models.ProductType{},
		&models.Product{},
		&models.VariantSelector{},
		&models.ProductVariant{},

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

func setupWs() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	wsCommands := commands.GetWsCommands()
	wsHub := websocket.NewWsHub(wsCommands)
	go wsHub.Run()
	server.WsHub = wsHub
	server.RedisDB = rdb
}

func setupControllers() {
	brandService := brandservice.NewBrandService(server.DB)
	brandController.BrandService = brandService
	brandController.DB = server.DB
}
