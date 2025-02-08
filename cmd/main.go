package main

import (
	"fmt"
	"net/http"

	"github.com/dyhalmeida/go-apis/configs"
	"github.com/dyhalmeida/go-apis/internal/entity"
	"github.com/dyhalmeida/go-apis/internal/infra/database"
	"github.com/dyhalmeida/go-apis/internal/infra/websever/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
	userHandler := handlers.NewUserHandler(userDB, config.GetTokenAuth(), config.GetJwtExpiresIn())

	r := chi.NewRouter()
	r.Use(middleware.Logger)

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

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
}
