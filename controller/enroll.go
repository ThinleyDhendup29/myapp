package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/date"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Enroll(w http.ResponseWriter, r *http.Request) {
	var e model.Enroll
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	e.Date_Enrolled = date.GetData()
	defer r.Body.Close()

	saveErr := e.EnrollStud()

	if saveErr != nil {
		//Contains check if "duplicate keys" words are there in the error
		if strings.Contains(saveErr.Error(), "Duplicate key") {
			httpResp.ResponseWithError(w, http.StatusBadRequest, "Duplicate entries")
			return
		} else {
			httpResp.ResponseWithError(w, http.StatusInternalServerError, saveErr.Error())
		}
	}
	// no error
	httpResp.ResponseWithJSON(w, http.StatusCreated, map[string]string{"status": "enrolled"})
}

func GetEnroll(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]

	//string sid to int type
	stdid, _ := strconv.ParseInt(sid, 10, 64)

	e := model.Enroll{
		StdId:    stdid,
		CourseID: cid,
	}
	getErr := e.Get()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, "No such enrollments")
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.ResponseWithJSON(w, http.StatusOK, e)
}

func GetEnrolls(w http.ResponseWriter, r *http.Request) {
	enrolls, getErr := model.GetAllEnrolls()
	if getErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.ResponseWithJSON(w, http.StatusOK, enrolls)
}

func DeleteEnroll(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]

	//string sid to int type
	stdid, _ := strconv.ParseInt(sid, 10, 64)

	e := model.Enroll{
		StdId:    stdid,
		CourseID: cid,
	}
	err := e.Delete()
	if err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"Status": "Deleted"})
}

