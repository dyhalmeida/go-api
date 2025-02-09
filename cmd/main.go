package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dyhalmeida/go-apis/configs"
	_ "github.com/dyhalmeida/go-apis/docs"
	"github.com/dyhalmeida/go-apis/internal/entity"
	"github.com/dyhalmeida/go-apis/internal/infra/database"
	"github.com/dyhalmeida/go-apis/internal/infra/websever/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Expert API Example
// @version 1.0
// @description Product API with Authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Diego Almeida
// @contact.url https://www.linkedin.com/in/dyhalmeida
// @contact.email dyhalmeida@gmail.com

// @host localhost:3333
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config := configs.NewConfig()

	db, err := gorm.Open(sqlite.Open(("test.db")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProductDb(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUserDb(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwt", config.GetTokenAuth()))
	r.Use(middleware.WithValue("jwtExpiresIn", config.GetJwtExpiresIn()))

	// r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.GetTokenAuth()))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/token", userHandler.GetJwtToken)

	swaggerURL := fmt.Sprintf("http://localhost:%s/docs/doc.json", config.GetServerPort())
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(swaggerURL)))

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
}

// Exemplo de custom middleware
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("request: %s %s", req.Method, req.URL.Path)
		next.ServeHTTP(res, req)
	})
}
