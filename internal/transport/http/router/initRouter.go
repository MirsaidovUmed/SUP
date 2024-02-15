package router

import (
	"SUP/internal/transport/http/handlers"
	"SUP/internal/transport/http/middleware"
	"SUP/pkg/http"

	"github.com/gorilla/mux"
)

func InitRouter(handlers *handlers.Handler, mw middleware.MiddlewareInterface) *mux.Router {
	router := http.NewRouter()
	router.Use(mw.TimeDuration)
	privateRouter := router.NewRoute().Subrouter()
	privateRouter.Use(mw.JWT)

	router.HandleFunc("/api/registration", handlers.Registration).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")

	privateRouter.HandleFunc("/api/updateUser", handlers.UpdateUser).Methods("POST")

	privateRouter.HandleFunc("/api/product/createProject", handlers.CreateProject).Methods("POST")
	privateRouter.HandleFunc("/api/product/updateProject", handlers.UpdateProject).Methods("POST")

	privateRouter.HandleFunc("/api/getProject", handlers.GetProject).Methods("GET")

	privateRouter.HandleFunc("/api/project/createTask", handlers.CreateTask).Methods("POST")
	privateRouter.HandleFunc("/api/project/updateTask", handlers.UpdateTask).Methods("POST")
	privateRouter.HandleFunc("/api/getTask", handlers.GetTask).Methods("GET")

	privateRouter.HandleFunc("/api/product/createProjectParticipant", handlers.CreateProjectParticipant).Methods("POST")
	privateRouter.HandleFunc("/api/product/deleteProjectParticipant", handlers.DeleteProjectParticipant).Methods("DELETE")

	return router
}
