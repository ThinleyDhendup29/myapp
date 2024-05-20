// model/model.go
// to interact with the database, helps to add the data in database
package model

import (
	"myapp/dataStore/postgres"
)

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

// Course struct represents a course record
type Course struct {
	Cid        string `json:"cid"`
	Coursename string `json:"coursename"`
}

// $1, .... is a place holder that will take the inputed values
const queryInsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES ($1,$2,$3,$4);"

// Create() func adds the data in database, have a return type of error
func (s *Student) Create() error {
	// used _ to ignore the return type, in case we dont have any return
	// use of postgres to access the database
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}

const queryGetUser = "SELECT stdid, firstname, lastname, email FROM student WHERE stdid=$1;"

func (s *Student) Read() error {
	return postgres.Db.QueryRow(queryGetUser, s.StdId).Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

const queryUpdate = "UPDATE student SET stdid=$1, firstname=$2, lastname=$3, email=$4 WHERE stdid=$5 RETURNING stdid;"

func (s *Student) Update(oldID int64) error {
	err := postgres.Db.QueryRow(queryUpdate, s.StdId, s.FirstName, s.LastName, s.Email, oldID).Scan(&s.StdId)
	return err
}

const queryDeleteUser = "DELETE FROM student WHERE stdid=$1 RETURNING stdid"

func (s *Student) Delete() error {
	if err := postgres.Db.QueryRow(queryDeleteUser, s.StdId).Scan(&s.StdId); err != nil {
		return err
	}
	return nil
}

func GetAllStudents() ([]Student, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM student;")
	if getErr != nil {
		return nil, getErr
	}

	// Create a slice of type student
	students := []Student{}

	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return nil, dbErr
		}
		students = append(students, s)
	}
	rows.Close()
	return students, nil
}

// This code defines functions to interact with a "course" table in a PostgreSQL database.

// Constant SQL statements for course operations
const (
	queryInsertCourse = "INSERT INTO course (cid, coursename) VALUES ($1, $2);"                // Insert a new course
	queryGetCourse    = "SELECT cid, coursename FROM course WHERE cid=$1;"                    // Get a course by ID
	queryUpdateC      = "UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid;" // Update a course
	queryDeleteCourse = "DELETE FROM course WHERE cid=$1 RETURNING cid;"                      // Delete a course
)

// Create method inserts a new course into the database
func (s *Course) Create() error {
	// Execute the insert query using prepared statement with parameters
	_, err := postgres.Db.Exec(queryInsertCourse, s.Cid, s.Coursename)
	return err
}

// Read method retrieves a course from the database by its ID
func (s *Course) Read() error {
	// Execute the query row using prepared statement with the course ID parameter
	return postgres.Db.QueryRow(queryGetCourse, s.Cid).Scan(&s.Cid, &s.Coursename)
}

// UpdateC method updates a course's details in the database
func (s *Course) UpdateC(oldID string) error {
	// Execute the update query row using prepared statement with new and old course IDs
	err := postgres.Db.QueryRow(queryUpdateC, s.Cid, s.Coursename, oldID).Scan(&s.Cid)
	return err
}

// DeleteC method removes a course from the database
func (s *Course) DeleteC() error {
	// Execute the delete query row using prepared statement with the course ID parameter
	if err := postgres.Db.QueryRow(queryDeleteCourse, s.Cid).Scan(&s.Cid); err != nil {
		return err
	}
	return nil
}

// GetAllCourses function retrieves all courses from the database
func GetAllCourses() ([]Course, error) {
	// Execute the query to get all courses
	rows, err := postgres.Db.Query("SELECT * from course;")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed even in case of errors

	courses := []Course{} // Initialize an empty slice to store courses

	for rows.Next() {
		// Create a new Course instance for each row
		var c Course
		// Scan the row values into the Course struct fields
		dbErr := rows.Scan(&c.Cid, &c.Coursename)
		if dbErr != nil {
			return nil, dbErr
		}
		// Add the course to the slice
		courses = append(courses, c)
	}

	return courses, nil
}
