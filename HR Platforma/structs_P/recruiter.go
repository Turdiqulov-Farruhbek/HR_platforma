package structs_P

import (
	"time"
)

type Recruiter struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Birthday    string    `json:"birthday"`
	Gender      string    `json:"gender"`
	CompanyID   string    `json:"company_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   int64     `json:"deleted_at"`
}

type RecruiterAll struct {
	ID          string    `json:"-"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Birthday    string    `json:"birthday"`
	Gender      string    `json:"gender"`
	CreatedAt   time.Time `json:"created_at"`
	CompanyID   Company   `json:"Company"`
}

type RecruiterCreate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
	CompanyID   string `json:"company_id"`
}

type RecruiterUpdate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
	CompanyID   string `json:"company_id"`
}

type RecruiterDeleted struct {
	ID string
}

type Recruiters struct {
	Recruiters []Recruiter
	Count      int
}
