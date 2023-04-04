package handlers

import (
	"../conf"
	"../model"
	"encoding/json"
	"io"
	"net/http"
)

func Config(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	result := make(map[string]string)
	result["FaceIOAppCode"] = conf.GetEnv().FaceIOAppCode

	jsonResult := new(model.JsonResult)
	jsonResult.Code = conf.GetEnv().CodeSuccess
	jsonResult.Msg = conf.GetEnv().MsgSuccess
	jsonResult.Result = result

	msg, _ := json.Marshal(jsonResult)

	_, _ = io.WriteString(w, string(msg))

}