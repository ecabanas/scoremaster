package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"scoremaster/backend/logs"
	"scoremaster/backend/services"
)

func SetupRouter(router *mux.Router) {
	router.HandleFunc("/api/quiz", services.GetQuiz).Methods("GET")
}

func SetupServer(router *mux.Router) {
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})
	logs.Error.Println(http.ListenAndServe(":8080", handlers.CORS(credentials, methods, ttl, origins)(router)))
}
