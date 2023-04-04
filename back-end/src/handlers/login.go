package handlers

import (
	"../conf"
	"../model"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	user :=new(model.User)
	user.LoginName=conf.GetEnv().LoginName
	user.LoginPassWord=conf.GetEnv().LoginPassWord

	body, _ := ioutil.ReadAll(r.Body)

	loginForm := make(map[string]interface{})
	if err := json.Unmarshal(body, &loginForm); err == nil {
		fmt.Println(loginForm)
	} else {
		fmt.Println(err)
	}

	jsonResult := new(model.JsonResult)

	if user.LoginName==loginForm["loginName"] && user.LoginPassWord==loginForm["passWord"] {
		Result := make(map[string]string)
		Result["accessToken"]=conf.GetEnv().CryptoKey
		jsonResult.Code = conf.GetEnv().CodeSuccess
		jsonResult.Msg = conf.GetEnv().MsgSuccess
		jsonResult.Result=Result
	}else {
		jsonResult.Code = conf.GetEnv().CodeErr
		jsonResult.Msg = conf.GetEnv().MsgErr
	}

	msg, _ := json.Marshal(jsonResult)

	_, _ = io.WriteString(w, string(msg))
}


func LoginWithFace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	body, _ := ioutil.ReadAll(r.Body)

	var loginUser model.User
	err := json.Unmarshal(body, &loginUser)
	if err != nil {
		fmt.Println("error in translating,", err.Error())
		return
	}

	jsonResult := new(model.JsonResult)

	if loginUser.FacialId==conf.GetEnv().LoginUserFacialId {
		Result := make(map[string]string)
		Result["accessToken"]=conf.GetEnv().CryptoKey
		jsonResult.Code = conf.GetEnv().CodeSuccess
		jsonResult.Msg = conf.GetEnv().MsgSuccess
		jsonResult.Result=Result
	}else {
		jsonResult.Code = conf.GetEnv().CodeErr
		jsonResult.Msg = conf.GetEnv().MsgErr
	}

	msg, _ := json.Marshal(jsonResult)

	_, _ = io.WriteString(w, string(msg))
}
