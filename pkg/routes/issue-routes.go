package routes

import (
	"github.com/gorilla/mux"
	"github.com/aryan/go-issue-backend/pkg/controllers"
)

var RegisterIssueRoutes = func(router *mux.Router){
	router.HandleFunc("/issue/", controllers.CreateIssue).Methods("POST")
	router.HandleFunc("/issue/", controllers.GetIssue).Methods("GET")
	router.HandleFunc("/issue/{issueId}", controllers.GetIssueById).Methods("GET")
	router.HandleFunc("/issue/{issueId}", controllers.UpdateIssue).Methods("PUT")
}