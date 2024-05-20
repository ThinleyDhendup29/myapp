package model

import "myapp/dataStore/postgres"

type Enroll struct {
	StdId         int64  `json:"stdid"`
	CourseID      string `json:"cid"`
	Date_Enrolled string `json:"date"`
}

const (
	queryEnrollStd    = "INSERT INTO enroll(std_id, course_id, date_enrolled) VALUES($1, $2, $3) RETURNING std_id;"
	queryGetEnroll    = "SELECT std_id, course_id, date_enrolled FROM enroll WHERE std_id=$1 and course_id=$2;"
	queryDeleteEnroll = "DELETE FROM enroll where std_id=$1 and course_id=$2 RETURNING std_id;"
)

func (e *Enroll) EnrollStud() error {
	result := postgres.Db.QueryRow(queryEnrollStd, e.StdId, e.CourseID, e.Date_Enrolled)
	return result.Scan(&e.StdId)
}

func (e *Enroll) Get() error {
	return postgres.Db.QueryRow(queryGetEnroll, e.StdId, e.CourseID).Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
}

func GetAllEnrolls() ([]Enroll, error) {
	rows, getErr := postgres.Db.Query("SELECT std_id, course_id, date_enrolled from enroll;")
	if getErr != nil {
		return nil, getErr
	}
	//create slice of type Course
	enrolls := []Enroll{}

	for rows.Next() {
		var e Enroll
		dbErr := rows.Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
		if dbErr != nil {
			return nil, dbErr
		}
		enrolls = append(enrolls, e)
	}
	rows.Close()
	return enrolls, nil
}

func (e *Enroll) Delete() error {
	return postgres.Db.QueryRow(queryDeleteEnroll, e.StdId, e.CourseID).Scan(&e.StdId)
}