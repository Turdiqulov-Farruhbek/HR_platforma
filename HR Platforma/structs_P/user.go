package structs_P

import "time"

type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Birthday    string    `json:"birthday"`
	Gender      string    `json:"gender"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   int64     `json:"deleted_at"`
}


type UserAll struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Birthday    string    `json:"birthday"`
	Gender      string    `json:"gender"`
}

type UserCreate struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
}

type UserUpdate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
}

type UserDeleted struct {
	ID string
}

type Users struct {
	Users []User
	Count int
}
