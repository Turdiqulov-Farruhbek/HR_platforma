package structs_P

import (
	"time"
)

type Vacancy struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	MinExp      int       `json:"min_exp"`
	CompanyID   string    `json:"company_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   int64     `json:"deleted_at"`
}


type VacancyAll struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	MinExp      int       `json:"min_exp"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CompanyID   Company   `json:"Company"`
}

type VacancyCreated struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	MinExp      int    `json:"min_exp"`
	CompanyID   string `json:"company_id"`
	Description string `json:"description"`
}

type VacancyUpdate struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	MinExp      int       `json:"min_exp"`
	CompanyID   string    `json:"company_id"`
	Description string    `json:"description"`
}

type VacancyDeleted struct {
	ID string	
}

type Vacancies struct {
	Vacancies []Vacancy `json:"vacancies"`
	Count     int       `json:"count"`
}
