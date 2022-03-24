package autocheck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadConfig(UserMap map[string]*UserList, jsonPath string) {
	f, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&UserMap)
	if err != nil {
		fmt.Printf("json decode has error:%v\n", err)
	}
}

func GetSessionData(UserMap map[string]*UserList, sessionPath string) (session map[string]string) {
	session = make(map[string]string)
	url := "https://yjsxx.whut.edu.cn/wx/api/login/checkBind"
	f, err := os.OpenFile(sessionPath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err.Error())
	}
	for userName, v := range UserMap {
		tempName := userName
		tempV := v
		tempUser := UserInfo{
			tempV.StudentNumber,
			tempV.IDCard,
		}
		payload, err := json.Marshal(tempUser)
		if err != nil {
			fmt.Println("Session请求体构建时Json数据错误:", err.Error())
			os.Exit(1)
		}
		reader := bytes.NewReader(payload)
		req, _ := http.NewRequest("POST", url, reader)
		for header, headerV := range SessionHeader {
			req.Header.Add(header, headerV)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("得到Session失败")
			os.Exit(1)
		}
		body, _ := ioutil.ReadAll(resp.Body)

		data := make(map[string]interface{})
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("教务处系统更新，重新抓包")
			os.Exit(1)
		}
		session[tempName] = data["data"].(map[string]interface{})["sessionId"].(string)
		_ = resp.Body.Close()
		_, err = f.Write([]byte(tempName + ":" + session[tempName] + "\n"))
	}
	_ = f.Close()
	return session
}

func BindUserInfo(UserMap map[string]*UserList, SessionMap map[string]string) {
	fmt.Println("--------绑定阶段--------")
	url := "https://yjsxx.whut.edu.cn/wx/api/login/bindUserInfo"
	for userName, v := range UserMap {
		tempName := userName
		tempV := v
		tempUser := UserInfo{
			tempV.StudentNumber,
			tempV.IDCard,
		}
		payload, err := json.Marshal(tempUser)
		if err != nil {
			fmt.Println("绑定User请求体构建时Json数据错误:", err.Error())
			os.Exit(1)
		}
		reader := bytes.NewReader(payload)
		req, _ := http.NewRequest("POST", url, reader)
		for header, headerV := range BindUserHeader {
			req.Header.Add(header, headerV)
		}
		cookie := fmt.Sprintf("JSESSIONID=%s", SessionMap[tempName])
		//fmt.Println(cookie)
		req.Header.Add("Cookie", cookie)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("绑定User失败")
			os.Exit(1)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		data := make(map[string]interface{})
		_ = json.Unmarshal(body, &data)
		fmt.Println(data)
		_ = resp.Body.Close()

	}
	fmt.Println("--------绑定阶段--------\n")
}

func Report(UserMap map[string]*UserList, SessionMap map[string]string) {
	fmt.Println("--------报送阶段--------")
	url := "https://yjsxx.whut.edu.cn/wx/./monitorRegister"
	for userName, v := range UserMap {
		tempName := userName
		tempV := v
		currentAddress := tempV.Province + tempV.City + tempV.County + tempV.Street
		tempReportJson := ReportInfo{
			DiagnosisName:   "",
			RelationWithOwn: "",
			CurrentAddress:  currentAddress,
			Remark:          "无",
			HealthInfo:      "正常",
			IsDiagnosis:     0,
			IsFever:         0,
			IsInSchool:      "0",
			IsLeaveChengdu:  0,
			IsSymptom:       "0",
			Temperature:     "36.5°C~36.9°C",
			Province:        tempV.Province,
			City:            tempV.City,
			County:          tempV.County,
		}
		payload, err := json.Marshal(tempReportJson)
		if err != nil {
			fmt.Println("体温请求体构建时Json数据错误:", err.Error())
			os.Exit(1)
		}
		reader := bytes.NewReader(payload)
		req, _ := http.NewRequest("POST", url, reader)
		for header, headerV := range ReportHeader {
			req.Header.Add(header, headerV)
		}
		cookie := fmt.Sprintf("JSESSIONID=%s", SessionMap[tempName])
		req.Header.Add("Cookie", cookie)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("报送失败")
			os.Exit(1)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		data := make(map[string]interface{})
		_ = json.Unmarshal(body, &data)
		fmt.Println(data)
		_ = resp.Body.Close()
	}
	fmt.Println("--------报送阶段--------\n")
}

func CancelBind(UserMap map[string]*UserList, SessionMap map[string]string) {
	fmt.Println("--------解绑阶段--------")
	url := "https://yjsxx.whut.edu.cn/wx/api/login/cancelBind"
	for userName, _ := range UserMap {
		tempName := userName
		req, _ := http.NewRequest("POST", url, nil)
		for header, headerV := range CancelBindHeader {
			req.Header.Add(header, headerV)
		}
		cookie := fmt.Sprintf("JSESSIONID=%s", SessionMap[tempName])
		req.Header.Add("Cookie", cookie)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("解绑失败")
			os.Exit(1)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		data := make(map[string]interface{})
		_ = json.Unmarshal(body, &data)
		fmt.Println(data)
		_ = resp.Body.Close()
	}
	fmt.Println("--------解绑阶段--------")
}
