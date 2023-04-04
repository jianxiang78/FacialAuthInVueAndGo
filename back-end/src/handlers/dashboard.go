package handlers

import (
	"../conf"
	"../model"
	"encoding/json"
	"io"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	AccessToken := r.Header.Get("Access-Token")

	jsonResult := new(model.JsonResult)

	if conf.GetEnv().CryptoKey != AccessToken {
		jsonResult.Code = conf.GetEnv().CodeErr
		jsonResult.Msg = conf.GetEnv().MsgErr
	}else {
		user := new(model.User)
		user.LoginName=conf.GetEnv().LoginName
		user.FacialId=conf.GetEnv().LoginUserFacialId

		jsonResult.Code = conf.GetEnv().CodeSuccess
		jsonResult.Msg = conf.GetEnv().MsgSuccess
		jsonResult.Result = user
	}

	msg, _ := json.Marshal(jsonResult)
	w.Header().Add("Content-Type","application/json")
	_, _ = io.WriteString(w, string(msg))
}
