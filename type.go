package main

type UserList struct {
	UserInfo
	LocalInfo
}

type UserInfo struct {
	StudentNumber string `json:"sn"`
	IDCard string `json:"idCard"`
}

type LocalInfo struct {
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
	Street   string `json:"street"`
}

type ReportInfo struct {
	DiagnosisName   string      `json:"diagnosisName"`
	RelationWithOwn string      `json:"relationWithOwn"`
	CurrentAddress  string `json:"currentAddress"`
	Remark          string      `json:"remark"`
	HealthInfo      string      `json:"healthInfo"`
	IsDiagnosis     int         `json:"isDiagnosis"`
	IsFever         int         `json:"isFever"`
	IsInSchool      string      `json:"isInSchool"`
	IsLeaveChengdu  int         `json:"isLeaveChengdu"`
	IsSymptom       string      `json:"isSymptom"`
	Temperature     string `json:"temperature"`
	Province        string `json:"province"`
	City            string `json:"city"`
	County          string `json:"county"`
}