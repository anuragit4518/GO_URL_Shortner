package main

import (
	"log"
	"net/http"
	"url_shortener/db"
	"url_shortener/routes"
	"url_shortener/middlewares"

)


func main()  {

	// connect to MongoDB
	db.ConnectMongoDB()
	
	// register all routes
	routes.RegisterRoutes()

	log.Println("Server started on http://localhost:8080")
	


// ...

	handler := middlewares.CORSMiddleware(http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8080", handler))

	



}

