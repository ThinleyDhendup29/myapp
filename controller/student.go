package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// // Validate cookie
	// if !VerifyCookie(w, r){
	// 	return
	// }
	var stud model.Student
	decoder := json.NewDecoder(r.Body)

	// if err := decoder.Decode(&stud); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Invalid json data"))
	// 	return
	// }
	// if err := decoder.Decode(&stud); err != nil {
	// 	response, _ := json.Marshal(map[string]string{"error": "invalid json body"})
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(response)
	// 	return
	// }
	if err := decoder.Decode(&stud); err != nil {

		httpResp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	defer r.Body.Close()
	saveErr := stud.Create()
	if saveErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	// if err := decoder.Decode(&stud); err != nil{
	// 	httpResp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
	// 	return
	// }
	// if saveErr != nil {
	// 	response, _ := json.Marshal(map[string]string{"error": saveErr.Error()})

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(response)
	// 	return
	// }
	// if err := stud.Create(); err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Database error"))
	// 	return
	// }

	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte("Student added successfully"))
	// response, _ := json.Marshal(map[string]string{"status": "student added"})
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// w.Write(response)

	httpResp.ResponseWithJSON(w, http.StatusCreated, map[string]string{"status": "student added"})
}

func GetStud(w http.ResponseWriter, r *http.Request) {
	//get url parameter
	sid := mux.Vars(r)["sid"]
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdId}
	getErr := s.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, "Student not found")
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.ResponseWithJSON(w, http.StatusOK, s)
}

// convert string sid to int
func getUserId(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

func UpdateStud(w http.ResponseWriter, r *http.Request) {
	old_sid := mux.Vars(r)["sid"]
	old_stdId, idErr := getUserId(old_sid)
	if idErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	defer r.Body.Close()

	updateErr := stud.Update(old_stdId)

	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, "student not found")
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.ResponseWithJSON(w, http.StatusOK, stud)
	}
}

func DeleteStud(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdId}
	if err := s.Delete(); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	students, getErr := model.GetAllStudents()
	if getErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.ResponseWithJSON(w, http.StatusOK, students)
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
	// This function handles adding a new course
	// First, it creates a new instance of the Course model struct
	var cour model.Course

	// Next, it creates a new decoder to decode the JSON request body into the Course struct
	decoder := json.NewDecoder(r.Body)

	// It attempts to decode the request body into the Course struct.
	// If there is an error during decoding, it returns a bad request error with a message indicating that the JSON body is invalid.
	if err := decoder.Decode(&cour); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}

	// Close the request body to release resources
	defer r.Body.Close()

	// Call the Create method on the course instance to save the course to the database
	// If there is an error saving the course, it returns a bad request error with the error message from the Create method.
	saveErr := cour.Create()
	if saveErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}

	// If everything is successful, it returns a created status code with a JSON response indicating that the course has been added
	httpResp.ResponseWithJSON(w, http.StatusCreated, map[string]string{"Status": "Course Added"})
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	// This function retrieves a course based on the provided course ID

	// Extract the course ID (cid) from the request URL using mux.Vars
	cid := mux.Vars(r)["cid"]

	// Create a new Course object with the provided course ID
	c := model.Course{Cid: cid}

	// Call the Read method on the Course object to retrieve the course from the database
	getErr := c.Read()

	// Handle potential errors during retrieval
	if getErr != nil {
		// Check for specific error types
		switch getErr {
		case sql.ErrNoRows:
			// If sql.ErrNoRows is encountered, the course wasn't found
			httpResp.ResponseWithError(w, http.StatusNotFound, "Course not found")
		default:
			// For other errors, return an internal server error with the error message
			httpResp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}

	// If successful, return the retrieved course data as JSON with a status code of OK
	httpResp.ResponseWithJSON(w, http.StatusOK, c)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	// This function updates an existing course

	// Extract the course ID (old_cid) from the request URL using mux.Vars
	old_cid := mux.Vars(r)["cid"]

	// Create a new Course object to hold the updated data
	var cour model.Course

	// Decode the JSON request body into the Course object
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cour); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}

	// Close the request body to release resources
	defer r.Body.Close()

	// Call the UpdateC method on the Course object to update the course in the database
	updateErr := cour.UpdateC(old_cid)

	// Handle potential errors during update
	if updateErr != nil {
		// Check for specific error types
		switch updateErr {
		case sql.ErrNoRows:
			// If sql.ErrNoRows is encountered, the course wasn't found
			httpResp.ResponseWithError(w, http.StatusNotFound, "Course Not Found")
		default:
			// For other errors, return an internal server error with the error message
			httpResp.ResponseWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		// If successful, return the updated course data as JSON with a status code of OK
		httpResp.ResponseWithJSON(w, http.StatusOK, cour)
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// This function deletes a course

	// Extract the course ID (cid) from the request URL using mux.Vars
	cid := mux.Vars(r)["cid"]

	// Create a new Course object with the provided course ID
	s := model.Course{Cid: cid}

	// Call the DeleteC method on the Course object to delete the course from the database
	if err := s.DeleteC(); err != nil {
		// Handle errors during deletion
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// If successful, return a status code of OK with a JSON message indicating deletion
	httpResp.ResponseWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	// This function retrieves all courses

	// Call the GetAllCourses method (likely from your model package) to get all courses from the database
	courses, getErr := model.GetAllCourses()

	// Handle potential errors during retrieval
	if getErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}

	// If successful, return all retrieved courses data as JSON with a status code of OK
	httpResp.ResponseWithJSON(w, http.StatusOK, courses)
}
