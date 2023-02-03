package server

import (
	"github.com/gin-contrib/cors"
	apcontroller "persia_atlas/server/controllers/actual_product"
	brandcontroller "persia_atlas/server/controllers/brand"
	productcontroller "persia_atlas/server/controllers/product"
	ptypecontroller "persia_atlas/server/controllers/product_type"
	"persia_atlas/server/controllers/user"
	variantcontroller "persia_atlas/server/controllers/variant"
	apsrvc "persia_atlas/server/services/actual_product"
	brandsrvc "persia_atlas/server/services/brand"
	productsrvc "persia_atlas/server/services/product"
	ptypesrvc "persia_atlas/server/services/product_type"
	variantsrvc "persia_atlas/server/services/variant"
	"persia_atlas/server/websocket"
)

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

	brandService := brandsrvc.NewBrandService(server.DB)
	brandController := brandcontroller.NewBrandController(brandService, server.DB)
	brandController.RegisterRoutes(server.Router)

	ptService := ptypesrvc.NewProductTypeService(server.DB)
	ptController := ptypecontroller.NewProductTypeController(ptService, server.DB)
	ptController.RegisterRoutes(server.Router)

	apService := apsrvc.NewActualProductService(server.DB)
	apController := apcontroller.NewActualProductController(apService, server.DB)
	apController.RegisterRoutes(server.Router)

	productService := productsrvc.NewProductService(server.DB)
	productController := productcontroller.NewProductController(productService, server.DB)
	productController.RegisterRoutes(server.Router)

	variantService := variantsrvc.NewVariantService(server.DB)
	variantController := variantcontroller.NewVariantController(variantService, server.DB)
	variantController.RegisterRoutes(server.Router)
}
