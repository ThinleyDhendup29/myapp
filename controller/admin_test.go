package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test case to check admin login success
func TestAdmLogin(t *testing.T) {
	url := "http://localhost:8000/login"
	// create a json data
	var jsonStr = []byte(`{"email":"12230065.gcit@rub.edu.bt" , "password":"Tshering@1238"}`)
	//create http request object
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-type", "application/json") // header set bay
	//creating a client
	client := &http.Client{} // this particular client is going to send the request
	//send api login request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//defer the closing of the response body
	defer resp.Body.Close() // to access the resp body at a later stage
	// data is actual response
	data, _ := io.ReadAll(resp.Body)
	//expected response in json string
	expResp := `{"message":"login successful"}`
	// to see if the actual response matches the expected response
	assert.JSONEq(t, expResp, string(data))
	// compare response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// pass an email that doesnt exist
func TestAdmUserNotExist(t *testing.T) {
	url := "http://localhost:8000/login"
	var data = []byte(`{"email": "tc@gmail.com", "password": "kinlll"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// Check status code
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	// Check error message
	assert.JSONEq(t, `{"error": "sql: no rows in result set"}`, string(body))
}