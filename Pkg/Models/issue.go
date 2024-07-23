package models

import (
	"github.com/aryan/go-issue-backend/Pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Issue struct {
	gorm.Model
	Complaintid      string `json:"complaintid"`
	Designation      string `json:"designation"`
	Name             string `json:"name"`
	Enrollmentno     string `json:"enrollmentno"`
	Location         string `json:"location"`
	Area             string `json:"area"`
	Floorno          string `json:"floorno"`
	Roomno           string `json:"roomno"`
	Itemtype         string `json:"itemtype"`
	Equipmentid      string `json:"equipmentid"`
	Issuedescription string `json:"issuedescription"`
	Status           string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Issue{})
}

func (b *Issue) CreateIssue() *Issue {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllIssue() []Issue {
	var Issues []Issue
	db.Find(&Issues)
	return Issues
}

func GetIssueById(Id int64) (*Issue, *gorm.DB) {
	var getIssue Issue
	db := db.Where("ID=?", Id).Find(&getIssue)
	return &getIssue, db
}

func GetIssueByEnro(Enro int64) (*Issue, *gorm.DB) {
	var getenro Issue
	db := db.Where("enrollmentno=?", Enro).Find(&getenro)
	return &getenro, db
}
