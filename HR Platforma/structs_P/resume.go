package structs_P

import (
	"time"
)

type Resume struct {
	ID          string    `json:"id"`
	Position    string    `json:"position"`
	Experience  int       `json:"experience"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   int64     `json:"deleted_at"`
}

type ResumeAll struct {
	ID          string  `json:"id"`
	Position    string  `json:"position"`
	Experience  int     `json:"experience"`
	Description string  `json:"description"`
	UserID      UserAll `json:"User"`
}

type ResumeCreate struct {
	Position    string `json:"position"`
	Experience  int    `json:"experience"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

type ResumeUpdate struct {
	ID          string `json:"id"`
	Position    string `json:"position"`
	Experience  int    `json:"experience"`
	Description string `json:"description"`
}

type ResumeDeleted struct {
	ID string
}

type Resumes struct {
	Resumes []Resume `json:"resumes"`
	Count   int      `json:"count"`
}
