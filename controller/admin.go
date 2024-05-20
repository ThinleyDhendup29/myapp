package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request){
	var admin model.Admin

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil{
		httpResp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	saveErr := admin.Create()
	if saveErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	// no error
	httpResp.ResponseWithJSON(w, http.StatusCreated, map[string]string{"status": "admin added"})
}

func Login(w http.ResponseWriter, r *http.Request){
	var admin model.Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	getErr := admin.Get()
	if getErr != nil {
		httpResp.ResponseWithError(w, http.StatusUnauthorized, getErr.Error())
		return
	}
	cookie := http.Cookie{
		Name: "my-cookie",
		Value: "my-value",
		Expires: time.Now().Add(30 * time.Minute),
		Secure: true,
	}
	http.SetCookie(w, &cookie)
	
	httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"message":"login successful"})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Expires: time.Now(),
	})
	httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"message": "cookie deleted"})
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	// Retrive the "mu-cookie" cookie from the request
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		if err == http.ErrNoCookie{
			// No cookie found, redirect to the login page or return an error
			httpResp.ResponseWithError(w, http.StatusSeeOther, "cookie not found")
			return false
		}
		// some other error occured 
		httpResp.ResponseWithError(w, http.StatusInternalServerError,"internal server error")
		return false
	}
	// verify the cookie value
	if cookie.Value != "my-value" {
		// Invalid cookie value, redirect to the login page or return an error
		httpResp.ResponseWithError(w, http.StatusUnauthorized, "cookie does not match")
		return false
	}
	return true
}
