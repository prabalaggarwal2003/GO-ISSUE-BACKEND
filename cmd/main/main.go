package main

import (
	"log"
	"net/http"

	"github.com/aryan/go-issue-backend/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"os"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterIssueRoutes(r)
	http.Handle("/", r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:", r))
}
