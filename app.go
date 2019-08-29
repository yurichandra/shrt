package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yurichandra/shrt/db"
	"github.com/yurichandra/shrt/handler"
	"github.com/yurichandra/shrt/service"
)

func initRoutes() chi.Router {
	routes := chi.NewRouter()

	routes.Use(render.SetContentType(render.ContentTypeJSON))

	routes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]interface{}{
			"name":    "shrt-api",
			"version": "1",
		}

		res, _ := json.Marshal(payload)
		w.Write(res)
	})

	routes.Mount("/auth", handler.NewAuthHandler(authService).GetRoutes())
	routes.Mount("/shorten", handler.NewShortenerHandler(shortenerService).GetRoutes())

	return routes
}

func serveHTTP() {
	router := initRoutes()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("App running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func runMigration() {
	fmt.Println("Running migration")
	db.Migrate()
	fmt.Println("Migration completed!")
}

func seedKeys() {
	fmt.Println("Seeding keys...")
	redisService := service.InitRedisService(redisClient)
	redisService.Init()
	fmt.Println("Seeding completed...")
}
