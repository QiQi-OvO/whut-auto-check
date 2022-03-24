package main

import (
	"github.com/spf13/pflag"
	"strings"
	"whut-auto-check/autocheck"
)

func main() {
	var jsonPath string
	pflag.StringVar(&jsonPath, "path", "./res/userInfo.json", "path for userInfo.json")
	pflag.Parse()
	pathPrefix := strings.Split(jsonPath, "/")
	sessionPath := pathPrefix[0] + "/res/session_backup.txt"
	UserMap := make(map[string]*autocheck.UserList)
	autocheck.ReadConfig(UserMap, jsonPath)
	SeesionMap := autocheck.GetSessionData(UserMap, sessionPath)
	autocheck.BindUserInfo(UserMap, SeesionMap)
	autocheck.Report(UserMap, SeesionMap)
	autocheck.CancelBind(UserMap, SeesionMap)
}
