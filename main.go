package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"strings"
	"time"
	"whut-auto-check/autocheck"
)

func main() {
	fmt.Println("报送时间:",time.Now())
	var jsonPath string
	pflag.StringVar(&jsonPath, "path", "./res/userInfo.json", "path for userInfo.json")
	pflag.Parse()
	pathPrefix := strings.Split(jsonPath, "userInfo.json")
	sessionPath := pathPrefix[0] + "session_backup.txt"
	UserMap := make(map[string]*autocheck.UserList)
	autocheck.ReadConfig(UserMap, jsonPath)
	SeesionMap := autocheck.GetSessionData(UserMap, sessionPath)
	autocheck.BindUserInfo(UserMap, SeesionMap)
	autocheck.Report(UserMap, SeesionMap)
	autocheck.CancelBind(UserMap, SeesionMap)
}
