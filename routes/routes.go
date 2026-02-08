package routes

import (
	
	"net/http"
	"url_shortener/handlers"
	"url_shortener/middlewares"
)

func RegisterRoutes() {

	// auth routes
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler)
	http.HandleFunc("/reset-password", handlers.ResetPasswordHandler)
	http.HandleFunc(
		"/shorten",
		middlewares.AuthMiddleware(handlers.ShortenURLHandler),
	)
	http.HandleFunc("/", handlers.RedirectHandler)
	http.HandleFunc("/my-urls",middlewares.AuthMiddleware(handlers.MyURLsHandler),
)
	
		
	
}

