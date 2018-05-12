package controllers

import (
	"net/http"
	"common"
	"models"
)

func Counts(w http.ResponseWriter, r *http.Request) {
	statistic := models.GetStatisticCount()
	w.Write(common.NewResponseData(statistic, "").Encode())
}