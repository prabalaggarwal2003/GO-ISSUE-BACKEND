package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/aryan/go-issue-backend/pkg/utils"
	"github.com/aryan/go-issue-backend/pkg/models"
)

var NewIssue models.Issue

func GetIssue(w http.ResponseWriter, r *http.Request){
	newIssue:=models.GetAllIssue()
	res, _ :=json.Marshal(newIssue)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetIssueById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	issueId := vars["issueId"]
	ID, err:= strconv.ParseInt(issueId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	issueDetails, _:= models.GetIssueById(ID)
	res, _ := json.Marshal(issueDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateIssue(w http.ResponseWriter, r *http.Request){
	CreateIssue := &models.Issue{}
	utils.ParseBody(r, CreateIssue)
	b:= CreateIssue.CreateIssue()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateIssue(w http.ResponseWriter, r *http.Request){
	var updateIssue = &models.Issue{}
	utils.ParseBody(r, updateIssue)
	vars := mux.Vars(r)
	issueId := vars["issueId"]
	ID, err := strconv.ParseInt(issueId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	issueDetails, db:=models.GetIssueById(ID)
	if updateIssue.Status == "Pending"{
		issueDetails.Status = "Complete"
	}

	db.Save(&issueDetails)
	res, _ := json.Marshal(issueDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}