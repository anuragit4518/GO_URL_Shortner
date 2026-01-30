package routes

import (
	"fmt"
	"net/http"
	"url_shortener/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", homeHandler)
	// auth routes
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler)
	http.HandleFunc("/reset-password", handlers.ResetPasswordHandler)
}

func homeHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w," url shortner server is running")
}