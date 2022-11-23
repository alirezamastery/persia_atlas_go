package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"persia_atlas/server/controllers/brand"
	"persia_atlas/server/controllers/user"
	"persia_atlas/server/models"
	"persia_atlas/server/services/brand"
	"persia_atlas/server/websocket"
)

type Server struct {
	DB      *gorm.DB
	Router  *gin.Engine
	WsHub   *websocket.WsHub
	RedisDB *redis.Client
}

func (server *Server) connectDatabase() {
	var err error

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	server.DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		CreateBatchSize: 1000,
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

func (server *Server) addMiddlewares() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:9200", // quasar live server
		"http://127.0.0.1:5500", // vscode live server
	}
	server.Router.Use(cors.New(config))
}

func (server *Server) setupRoutes() {
	websocket.RegisterRoutes(server.Router, server.DB, server.WsHub, server.RedisDB)
	user.RegisterRoutes(server.Router, server.DB)

	brandService := brandservice.NewBrandService(server.DB)
	brandController := brandcontroller.NewBrandController(brandService, server.DB)
	brandController.RegisterRoutes(server.Router)
}

func (server *Server) Initialize() {
	server.connectDatabase()
	server.migrateDatabase()

	server.Router = gin.Default()

	server.addMiddlewares()

	server.setupRoutes()
}

func (server *Server) Run(addr string) {
	err := server.Router.Run(addr)
	if err != nil {
		log.Fatalln("error in running router:", err)
		return
	}
}
