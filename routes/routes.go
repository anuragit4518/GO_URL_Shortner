package routes

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w," url shortner server is running")
}