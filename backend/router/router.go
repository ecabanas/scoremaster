package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"scoremaster/backend/logs"
	"scoremaster/backend/services"
)

const (
	ServerPort     = ":8080"
	AllowedMethods = "GET"
	AllowedOrigins = "*"
)

func SetupRouter(router *mux.Router) {
	router.HandleFunc("/api/quiz", services.GetQuestionsWithAnswers).Methods("GET")
	router.HandleFunc("/api/quiz", services.GetCategories).Methods("GET")
}

func StartServer(router http.Handler) {
	logs.Info.Printf("Starting server on port %s...", ServerPort)
	err := http.ListenAndServe(ServerPort, router)
	logs.Error.FatalIf(err, "Failed to start server")
}

func SetupServer(router *mux.Router) http.Handler {
	corsMiddleware := configureCORS([]string{AllowedMethods}, []string{AllowedOrigins})
	return corsMiddleware(router)
}

func configureCORS(allowedMethods []string, allowedOrigins []string) func(http.Handler) http.Handler {
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods(allowedMethods)
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins(allowedOrigins)

	return handlers.CORS(credentials, methods, ttl, origins)

}
