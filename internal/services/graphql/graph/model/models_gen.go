// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddVisitRecordInput struct {
	UserID     string `json:"userId"`
	DoctorName string `json:"doctorName"`
	VisitDate  string `json:"visitDate"`
}

type AuthData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthPayload struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

type Doctor struct {
	DoctorID       string `json:"doctorId"`
	DoctorName     string `json:"doctorName"`
	Specialization string `json:"specialization"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateUserHealthDataInput struct {
	UserID   string   `json:"userId"`
	Age      *int     `json:"age,omitempty"`
	Height   *float64 `json:"height,omitempty"`
	Weight   *float64 `json:"weight,omitempty"`
	Pulse    *int     `json:"pulse,omitempty"`
	Pressure *string  `json:"pressure,omitempty"`
}

type UserAccountData struct {
	UserID string `json:"userId"`
}

type UserAccountPayload struct {
	Username      string   `json:"username"`
	Age           int      `json:"age"`
	Height        float64  `json:"height"`
	Weight        float64  `json:"weight"`
	Pulse         int      `json:"pulse"`
	Pressure      string   `json:"pressure"`
	DailyWater    *float64 `json:"dailyWater,omitempty"`
	BodyMassIndex *float64 `json:"bodyMassIndex,omitempty"`
}

type VisitRecord struct {
	RecordNumber   string `json:"recordNumber"`
	DoctorName     string `json:"doctorName"`
	Specialization string `json:"specialization"`
	VisitDate      string `json:"visitDate"`
}
