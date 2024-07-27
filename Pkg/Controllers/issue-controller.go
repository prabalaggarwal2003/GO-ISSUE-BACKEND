package controllers

import(
	"encoding/json"
	
	"github.com/gorilla/mux"
	"net/http"

	"github.com/aryan/go-issue-backend/Pkg/Utils"
	"github.com/aryan/go-issue-backend/Pkg/Models"
)

var NewIssue models.Issue
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func GetIssue(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	newIssue:=models.GetAllIssue()
	res, _ :=json.Marshal(newIssue)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetIssueById(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	vars := mux.Vars(r)
	complaintId := vars["complaintId"]
	
	issueDetails, _:= models.GetIssueById(complaintId)
	res, _ := json.Marshal(issueDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetIssueByEnro(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	vars := mux.Vars(r)
	enro := vars["enro"]
	
	issueDetails, _:= models.GetIssueByEnro(enro)
	res, _ := json.Marshal(issueDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateIssue(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	CreateIssue := &models.Issue{}
	utils.ParseBody(r, CreateIssue)
	b:= CreateIssue.CreateIssue()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateIssue(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	var updateIssue = &models.Issue{}
	utils.ParseBody(r, updateIssue)
	vars := mux.Vars(r)
	issueId := vars["issueId"]
	
	issueDetails, db:=models.GetIssueById(issueId)
	if updateIssue.Status == "Pending"{
		issueDetails.Status = "Complete"
	}

	db.Save(&issueDetails)
	res, _ := json.Marshal(issueDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}