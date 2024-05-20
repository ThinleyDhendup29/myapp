package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEnroll(t *testing.T) {
	url := "http://localhost:8000/enroll"
	var jsonStr = []byte(`{"stdid": 1002, "cid": "202"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check status code
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Read response body
	body, _ := io.ReadAll(resp.Body)

	// Check response body
	expResp := `{"status":"enrolled"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestGetEnroll(t *testing.T) {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8000/enroll/1002/202")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	expResp := `{"stdid":1002, "cid":"202","date":"2024-05-20T14:36:53Z"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestDeleteEnroll(t *testing.T) {
	url := "http://localhost:8000/enroll/1002/202"
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expResp := `{"Status": "Deleted"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestEnrollNotFound(t *testing.T) {
	assert := assert.New(t)
	c := http.Client{}
	r, _ := c.Get("http://localhost:8000/enroll/1002/202")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(http.StatusNotFound, r.StatusCode)
	expResp := `{"error" : "No such enrollments"}`
	assert.JSONEq(expResp, string(body))
}
