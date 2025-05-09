package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"scoremaster/backend/controllers"
	"scoremaster/backend/logs"
)

var (
	ServerPort                 = ":8080"
	AllowedMethods             = []string{"GET","POST","PUT","DELETE","OPTIONS"}
	AllowedOrigins             = []string{"http://localhost:8080", "http://localhost:5173"}
	allowedHeaders				= []string{"Content-Type", "Authorization", "Origin", "Accept", "X-Requested-With"}
	QuizBaseRoute              = "/api/quizzes"
	CategoryBaseRoute          = "/api/categories"
	QuestionBaseRoute          = "/api/questions"
	ParticipantBaseRoute       = "/api/participants"
	ParticipantAnswerBaseRoute = "/api/participants-answers"
)

func SetupRouter(router *mux.Router) {
	// Create a new Participant
	router.HandleFunc(ParticipantBaseRoute, controllers.CreateParticipant).Methods("POST")
	router.HandleFunc(ParticipantBaseRoute+"/check-email", controllers.CheckEmailExists).Methods("POST")


	//Quiz-related routes
	router.HandleFunc(QuizBaseRoute, controllers.GetQuestionsWithAnswers).Methods("GET")
	router.HandleFunc(QuizBaseRoute, controllers.GetCategories).Methods("GET")

	// Category-related routes
	router.HandleFunc(CategoryBaseRoute, controllers.GetCategories).Methods("GET")
	router.HandleFunc(CategoryBaseRoute, controllers.CreateCategory).Methods("POST")

	// Question-related routes
	router.HandleFunc(QuestionBaseRoute, controllers.GetQuestionsWithAnswers).Methods("GET")
	router.HandleFunc(QuestionBaseRoute, controllers.CreateQuestion).Methods("POST")
	router.HandleFunc(QuestionBaseRoute + "/{id}", controllers.GetQuestionByID).Methods("GET")

	// ParticipantAnswer-related routes
	router.HandleFunc(ParticipantAnswerBaseRoute, controllers.CreateParticipantAnswer).Methods("POST")

}

func StartServer(router http.Handler) {
	logs.Info.Printf("Starting server on port %s...", ServerPort)

	err := http.ListenAndServe(ServerPort, router)
	logs.Error.FatalIf(err, "Failed to start server")
}

func SetupServer(router *mux.Router) http.Handler {
	corsMiddleware := configureCORS(
		AllowedMethods,
		AllowedOrigins,
	allowedHeaders)
	return corsMiddleware(router)
}

func configureCORS(allowedMethods []string, allowedOrigins []string, allowedHeaders []string) func(http.Handler) http.Handler {
	
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods(allowedMethods)
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins(allowedOrigins)
	headers := handlers.AllowedHeaders(allowedHeaders)

	return handlers.CORS(credentials, methods, ttl, origins, headers)

}
