package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStudent(t *testing.T) {
	url := "http://localhost:8000/student"
	var jsonStr = []byte(`{"stdid":1111 , "fname":"Choida" , "lname": "Thukten" , "email":"thuks@gmail.com"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check for expected error response due to duplicate key violation
	if resp.StatusCode == http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		expectedError := `{"error":"pq: duplicate key value violates unique constraint \"student_email_key\""}`
		assert.JSONEq(t, expectedError, string(body))
		return
	}

	// If response status code is not as expected, fail the test
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	expResp := `{"status": "student added"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestGetStudent(t *testing.T) {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8000/student/1")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	expResp := `{"email":"td@gmail.com", "fname":"Thinley", "lname":"Dorji", "stdid":1}`
	assert.JSONEq(t, expResp, string(body))
}

func TestDeleteStudent(t *testing.T) { 
	url := "http://localhost:8000/student/3" 
	req, _ := http.NewRequest("DELETE", url, nil) 
	client := &http.Client{} 
	resp, err := client.Do(req)

	if err != nil { 
		panic(err)
	} 

	defer resp.Body.Close() 
	body, _ := io.ReadAll(resp.Body) 
	assert.Equal(t, http.StatusOK, resp.StatusCode) 
	expResp := `{"status":"deleted"}` 
	assert.JSONEq(t, expResp, string(body)) 
}

func TestStudentNotFound(t *testing.T) {
	assert := assert.New(t)
	c := http.Client{}
	r, _ := c.Get("http://localhost:8000/student/008")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(http.StatusNotFound, r.StatusCode)
	expResp := `{"error" : "Student not found"}`
	assert.JSONEq(expResp, string(body))
}
