package controllers

import (
	"net/http"
	"models"
	"common"
)

func StudentDict(w http.ResponseWriter, r *http.Request) {
	confDict := models.GetConfDictFile()
	w.Write(common.NewResponseData(confDict, "").Encode())
}