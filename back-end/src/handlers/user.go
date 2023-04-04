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

func UserEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	AccessToken := r.Header.Get("Access-Token")

	jsonResult := new(model.JsonResult)

	if conf.GetEnv().CryptoKey != AccessToken {
		jsonResult.Code = conf.GetEnv().CodeErr
		jsonResult.Msg = conf.GetEnv().MsgErr
	}else {
		body, _ := ioutil.ReadAll(r.Body)
		var inputUser model.User
		err := json.Unmarshal(body, &inputUser)
		if err != nil {
			fmt.Println("error in translating,", err.Error())
			return
		}

		user := new(model.User)
		user.LoginName=conf.GetEnv().LoginName

		user.Name = inputUser.Name
		user.FacialId = inputUser.FacialId
		conf.GetEnv().LoginUserFacialId=inputUser.FacialId

		jsonResult.Code = conf.GetEnv().CodeSuccess
		jsonResult.Msg = conf.GetEnv().MsgSuccess
		jsonResult.Result = user
	}

	msg, _ := json.Marshal(jsonResult)
	w.Header().Add("Content-Type","application/json")
	_, _ = io.WriteString(w, string(msg))
}
