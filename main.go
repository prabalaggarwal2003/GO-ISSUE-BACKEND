package main

import (
	"log"
	"net/http"

	"github.com/aryan/go-issue-backend/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterIssueRoutes(r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
