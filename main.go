package main

import (
	"log"
	"net/http"

	"github.com/aryan/go-issue-backend/Pkg/Routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/handlers"
)



func main() {
	r := mux.NewRouter()
	routes.RegisterIssueRoutes(r)
	http.Handle("/", r)
	
	log.Fatal(http.ListenAndServe("localhost:9010", 
	handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET","POST","PUT"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)))
}

