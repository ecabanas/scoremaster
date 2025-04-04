package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"scoremaster/backend/logs"
	"scoremaster/backend/services"
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
	router.HandleFunc(QuizBaseRoute, services.GetQuestionsWithAnswers).Methods("GET")
	router.HandleFunc(QuizBaseRoute, services.GetCategories).Methods("GET")

	// Category-related routes
	router.HandleFunc(CategoryBaseRoute, services.GetAllCategories).Methods("GET")
	router.HandleFunc(CategoryBaseRoute, services.CreateCategory).Methods("POST")

	// Question-related routes
	router.HandleFunc(QuestionBaseRoute, services.GetAllQuestions).Methods("GET")
	router.HandleFunc(QuestionBaseRoute, services.CreateQuestion).Methods("POST")

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
