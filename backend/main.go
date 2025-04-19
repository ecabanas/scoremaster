package main

import (
	"github.com/gorilla/mux"

	"scoremaster/backend/database"
	"scoremaster/backend/router"
)

func main() {
	r := mux.NewRouter()

	database.InitDatabase()
	router.SetupRouter(r)
	wrappedRouter := router.SetupServer(r)
	router.StartServer(wrappedRouter)

}
