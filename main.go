package main

import (
	"log"
	"net/http"
	"url_shortener/db"
	"url_shortener/routes"
)


func main()  {

	// connect to MongoDB
	db.ConnectMongoDB()
	
	// register all routes
	routes.RegisterRoutes()

	log.Println("Server started on http://localhost:8080")
	

	err := http.ListenAndServe(":8080",nil)
	if(err != nil){
		log.Fatal(err)
	}



}

