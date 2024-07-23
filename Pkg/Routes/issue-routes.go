package routes

import (
	"github.com/gorilla/mux"
	"github.com/aryan/go-issue-backend/Pkg/Controllers"
)

var RegisterIssueRoutes = func(router *mux.Router){
	router.HandleFunc("/issue/", controllers.CreateIssue).Methods("POST")
	router.HandleFunc("/issue/", controllers.GetIssue).Methods("GET")
	router.HandleFunc("/issue/{issueId}", controllers.GetIssueById).Methods("GET")
	router.HandleFunc("/issues/{enro}", controllers.GetIssueByEnro).Methods("GET")
	router.HandleFunc("/issue/{val}", controllers.UpdateIssue).Methods("PUT")
}