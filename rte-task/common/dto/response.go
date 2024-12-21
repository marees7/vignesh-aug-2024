package dto

import "time"

type Response struct {
	Message   string      `json:"message,omitempty"`
	Error     string      `json:"error,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Total     int64       `json:"total,omitempty"`
	Limit     int         `json:"limit,omitempty"`
	Offset    int         `json:"offset,omitempty"`
	Vacancy   int         `json:"vacancy,omitempty"`
	JobStatus string      `json:"job_status,omitempty"`
}

type ApplicantDetail struct {
	UserID      int       `json:"user_id,omitempty"`
	JobID       int       `json:"job_id,omitempty"`
	Experience  int       `json:"experience,omitempty"`
	Skills      string    `json:"skills,omitempty"`
	Language    string    `json:"language,omitempty"`
	Country     string    `json:"country,omitempty"`
	JobRole     string    `json:"job_role,omitempty"`
	CreatedAt   int       `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
}

type ErrorResponse struct {
	Error      error `json:"error,omitempty"`
	StatusCode int   `json:"status_code,omitempty"`
}

type LoginUser struct {
	Message  string `json:"message,omitempty"`
	Token    string `json:"token,omitempty"`
	ID       int    `json:"user_id,omitempty"`
	RoleType string `json:"role_type,omitempty"`
}
