package postgres_P

import (
	"database/sql"
	"fmt"
	"log/slog"
	"root/structs_P"
)

type ResumeManager struct {
	con *sql.DB
}

func NewResumeManager(conn *sql.DB) *ResumeManager {
	return &ResumeManager{con: conn}
}

func (rm *ResumeManager) Create(resume *structs_P.ResumeCreate) error {
	query = `INSERT INTO resume (position, experience, description, user_id) VALUES ($1, $2, $3, $4)`
	_, err = rm.con.Exec(query, resume.Position, resume.Experience, resume.Description, resume.UserID)
	if err != nil {
		slog.Error("error creating resume", "err:", err)
		return err
	}
	return nil
}

func (rm *ResumeManager) GetByID(resumeID string) (*structs_P.ResumeAll, error) {
	resume := structs_P.ResumeAll{}
	query = `SELECT r.id, r.position, r.experience, r.description, u.name, u.email, u.phone_number, u.birthday, u.gender 
	FROM resume r 
	JOIN  users u on r.user_id = u.id 
	WHERE r.id=$1 AND r.deleted_at=0 AND u.deleted_at=0`
	err = rm.con.QueryRow(query, resumeID).Scan(&resume.ID, &resume.Position, &resume.Experience, &resume.Description, &resume.UserID.Name,
		&resume.UserID.Email, &resume.UserID.PhoneNumber, &resume.UserID.Birthday, &resume.UserID.Gender)
	if err != nil {
		slog.Error("rm.con.QueryRow", "err", err)
		return nil, err
	}
	return &resume, nil
}

func (rm *ResumeManager) GetAllFiltrResume(position, experience, from, to string) ([]*structs_P.ResumeAll, error) {
	var resumes []*structs_P.ResumeAll
	query = `SELECT r.id, r.position, r.experience, r.description, u.name, u.email, u.phone_number, u.birthday, u.gender 
	FROM resume r 
	JOIN users u on r.user_id = u.id 
	WHERE r.deleted_at=0 AND u.deleted_at=0`

	args := []interface{}{}

	paramIndex := 1

	if position != "" {
		query += fmt.Sprintf(" AND position = $%d", paramIndex)
		args = append(args, position)
		paramIndex++
	}
	if experience != "" {
		query += fmt.Sprintf(" AND experience = $%d", paramIndex)
		args = append(args, experience)
		paramIndex++
	}
	if from != "" {
		query += fmt.Sprintf(" AND experience >= $%d", paramIndex)
		args = append(args, from)
		paramIndex++
	}
	if to != "" {
		query += fmt.Sprintf(" AND experience <= $%d", paramIndex)
		args = append(args, to)
		paramIndex++
	}
	// fmt.Println(query)
	// fmt.Println(args...)

	rows, err := rm.con.Query(query, args...)
	if err != nil {
		slog.Error("rm.con.Query","query", err)
		return nil, err
	}

	for rows.Next() {
		resume := structs_P.ResumeAll{}
		err = rows.Scan(&resume.ID, &resume.Position, &resume.Experience, &resume.Description,
			&resume.UserID.Name, &resume.UserID.Email, &resume.UserID.PhoneNumber, &resume.UserID.Birthday, &resume.UserID.Gender)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, &resume)
	}
	return resumes, nil
}

func (rm *ResumeManager) Update(resume *structs_P.ResumeUpdate) error {
	cloneResume, err := rm.GetByID(resume.ID)
	if err != nil {
		return err
	}

	fmt.Println(cloneResume)

	if resume.Description == "" {
		resume.Description = cloneResume.Description

	}
	if resume.Experience == -1 {
		resume.Experience = cloneResume.Experience

	}
	if resume.Position == "" {
		resume.Position = cloneResume.Position

	}
	fmt.Println(resume)

	query = `UPDATE resume SET position = $1, experience = $2, description = $3, updated_at = now() 
	WHERE id = $4 AND deleted_at = 0`
	_, err = rm.con.Exec(query, resume.Position, resume.Experience, resume.Description, resume.ID)
	if err != nil {
		slog.Error("Error updating resume", "err", err)
		return err
	}
	return nil
}

func (rm *ResumeManager) Delete(resumeID string) error {
	query = `UPDATE resume SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err = rm.con.Exec(query, resumeID)
	if err != nil {
		return err
	}
	return nil
}
