package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"scoremaster/backend/controllers"
	"scoremaster/backend/logs"
)

const (
	ServerPort        = ":8080"
	AllowedMethods    = "GET"
	AllowedOrigins    = "*"
	QuizBaseRoute     = "/api/quizzes"
	CategoryBaseRoute = "/api/categories"
	QuestionBaseRoute = "/api/questions"
)

func SetupRouter(router *mux.Router) {
	//Quiz-related routes
	router.HandleFunc(QuizBaseRoute, controllers.GetQuestionsWithAnswers).Methods("GET")
	router.HandleFunc(QuizBaseRoute, controllers.GetCategories).Methods("GET")

	// Category-related routes
	router.HandleFunc(CategoryBaseRoute, controllers.GetCategories).Methods("GET")
	router.HandleFunc(CategoryBaseRoute, controllers.CreateCategory).Methods("POST")

	// Question-related routes
	router.HandleFunc(QuestionBaseRoute, controllers.GetQuestionsWithAnswers).Methods("GET")
	router.HandleFunc(QuestionBaseRoute, controllers.CreateQuestion).Methods("POST")

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
