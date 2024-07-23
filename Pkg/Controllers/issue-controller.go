package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	issueId := vars["issueId"]
	ID, err:= strconv.ParseInt(issueId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	issueDetails, _:= models.GetIssueById(ID)
	res, _ := json.Marshal(issueDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetIssueByEnro(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	vars := mux.Vars(r)
	enro := vars["enro"]
	enrollmentno, err:= strconv.ParseInt(enro,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	issueDetails, _:= models.GetIssueByEnro(enrollmentno)
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
	ID, err:= strconv.ParseInt(issueId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	issueDetails, db:=models.GetIssueById(ID)
	if updateIssue.Status == "Pending"{
		issueDetails.Status = "Complete"
	}

	db.Save(&issueDetails)
	res, _ := json.Marshal(issueDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}