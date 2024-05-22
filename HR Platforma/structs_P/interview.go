package structs_P

import (
	"time"
)

type Interview struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	VacancyID     string    `json:"vacancy_id"`
	RecruiterID   string    `json:"recruiter_id"`
	InterviewDate string    `json:"interview_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     int64     `json:"deleted_at"`
}

type Interv struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	VacancyID   string `json:"vacancy_id"`
	RecruiterID string `json:"recruiter_id"`
}

type InterviewCreate struct {
	ID          string    `json:"id"`
	UserID      User      `json:"user_id"`
	VacancyID   Vacancy   `json:"vacancy_id"`
	RecruiterID Recruiter `json:"recruiter_id"`
}

type InterviewAll struct {
	ID            string     `json:"id"`
	UserID        User       `json:"user"`
	VacancyID     VacancyAll `json:"vacancy"`
	RecruiterID   Recruiter  `json:"recruiter"`
	RecruiterComp Company    `json:"recruiter_comp"`
	InterviewDate string     `json:"interview_date"`
}

type InterviewUpdate struct {
	ID            string    `json:"id"`
	UserID        User      `json:"User"`
	VacancyID     Vacancy   `json:"Vacancy"`
	RecruiterID   Recruiter `json:"recruiter"`
	InterviewDate string    `json:"interview_date"`
}

type InterviewDeleted struct {
	ID string
}

type Interviews struct {
	Interviews []Interview `json:"interviews"`
	Count      int         `json:"count"`
}
