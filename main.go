package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// var tpl = template.Must(template.ParseFiles("views/index.html"))

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	tpl.Execute(w, nil)
// }

// func testRequest(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Hello World!</h1>"))
// 	fmt.Println(r.GetBody())
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// mux := http.NewServeMux()

	// mux.HandleFunc("/", indexHandler)
	// mux.HandleFunc("/test", testRequest)
	// http.ListenAndServe(":"+port, mux)

	/* TESTING WITH GIN */

	// Set the router as the default router shipped with gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello",
			})
		})
	}

	// Start and run the server
	router.Run(port)
}
