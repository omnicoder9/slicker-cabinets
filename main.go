package main

import (
	"html/template"
	"log"
	"slicker-cabinets/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Add template functions
	r.SetFuncMap(template.FuncMap{
		"makeSlice": func(n int) []struct{} {
			return make([]struct{}, n)
		},
	})

	// Serve static files
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Routes
	r.GET("/", handlers.RenderHome)
	r.GET("/contact", handlers.RenderContact)
	r.GET("/services", handlers.RenderServices)
	r.GET("/reviews", handlers.RenderReviews)
	r.GET("/about", handlers.RenderAbout)
	r.GET("/login", handlers.RenderLogin)

	// API Routes
	api := r.Group("/api")
	{
		api.POST("/contact", handlers.HandleContactSubmit)
		api.POST("/login", handlers.HandleLoginSubmit)
	}

	log.Println("Server starting on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
