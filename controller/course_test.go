package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	url := "http://localhost:8000/course"
	var jsonStr = []byte(`{"cid":"202" , "coursename":"Dzongkha" }`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	expResp := `{"Status": "Course Added"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestGetCourse(t *testing.T) {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8000/course/103")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	expResp := `{"cid":"103", "coursename":"Programming"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestDeleteCourse(t *testing.T) {
	url := "http://localhost:8000/course/101"
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

func TestCourseNotFound(t *testing.T) {
	assert := assert.New(t)
	c := http.Client{}
	r, _ := c.Get("http://localhost:8000/course/107")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(http.StatusNotFound, r.StatusCode)
	expResp := `{"error" : "Course not found"}`
	assert.JSONEq(expResp, string(body))
}
