package postgres_P

import (
	"database/sql"
	"fmt"
	"log/slog"
	"root/structs_P"
)

type RecruiterManager struct {
	con *sql.DB
}

func NewRecruiterManager(conn *sql.DB) *RecruiterManager {
	return &RecruiterManager{con: conn}
}


func (rm *RecruiterManager) Create(recruiter *structs_P.RecruiterCreate) error {
	query = `INSERT INTO recruiter (name, email, phone_number, birthday, gender, company_id) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = rm.con.Exec(query, recruiter.Name, recruiter.Email, recruiter.PhoneNumber, recruiter.Birthday, recruiter.Gender, recruiter.CompanyID)
	if err != nil {
		slog.Error("Error creating", "err", err)
		return err
	}
	return nil
}


func (rm *RecruiterManager) GetByID(recruiterID string) (*structs_P.RecruiterAll, error) {
	recruiter := structs_P.RecruiterAll{}
	query = `SELECT r.id, r.name, r.email, r.phone_number, r.birthday, r.gender, r.created_at, c.name, c.location, c.workers 
	FROM recruiter r 
	JOIN company c ON r.company_id = c.id
	WHERE r.id = $1 `
	err = rm.con.QueryRow(query, recruiterID).Scan(
		&recruiter.ID, &recruiter.Name, &recruiter.Email, &recruiter.PhoneNumber,
		&recruiter.Birthday, &recruiter.Gender, &recruiter.CreatedAt,
		&recruiter.CompanyID.Name, &recruiter.CompanyID.Location, &recruiter.CompanyID.Workers)
	if err != nil {
		slog.Error("GetById metodi querysida xato:", "err",  err)
		return nil, err
	}
	return &recruiter, nil
}

func (rm *RecruiterManager) GetAllRecruiter(gender, companyid, from, to string) ([]*structs_P.RecruiterAll, error) {
	var resumes []*structs_P.RecruiterAll
	query = `SELECT r.name, r.email, r.phone_number, r.birthday,r.gender, r.created_at, c.name, c.location, c.workers 
	FROM recruiter r 
	JOIN  company c on r.company_id = c.id 
	WHERE r.deleted_at=0 AND c.deleted_at=0`
	
	args := []interface{}{}
	
	paramIndex := 1

	
	if gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", paramIndex)
		args = append(args, gender)
		paramIndex++
	}
	if companyid != "" {
		query += fmt.Sprintf(" AND company_id = $%d", paramIndex)
		args = append(args, companyid)
		paramIndex++
	}
	if from != "" {
		query += fmt.Sprintf(" AND EXTRACT(year FROM age(birthday)) >= $%d", paramIndex)
		args = append(args, from)
		paramIndex++
	}
	if to != "" {
		query += fmt.Sprintf(" AND EXTRACT(year FROM age(birthday)) <= $%d", paramIndex)
		args = append(args, to)
		paramIndex++
	}
	
	
	rows, err := rm.con.Query(query, args...)
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		resume := structs_P.RecruiterAll{}
		err = rows.Scan(&resume.Name, &resume.Email, &resume.PhoneNumber, &resume.Birthday, &resume.Gender, 
			&resume.CreatedAt,&resume.CompanyID.Name, &resume.CompanyID.Location, &resume.CompanyID.Workers)
			if err != nil {
				slog.Error("GetAllRecruiter metodi querysida xato:", "err", err)
				return nil, err
			}
			resumes = append(resumes, &resume)
		}
		return resumes, nil
}

func (rm *RecruiterManager) Update(recruiter *structs_P.RecruiterUpdate) error {
	cloneRecruiter, err := rm.GetByID(recruiter.ID)
	if err != nil {
		return err
	}


	if recruiter.Name == "" {
		recruiter.Name = cloneRecruiter.Name

	}
	if recruiter.Email == "" {
		recruiter.Email = cloneRecruiter.Email

	}
	if recruiter.PhoneNumber == "" {
		recruiter.PhoneNumber = cloneRecruiter.PhoneNumber

	}
	if recruiter.Birthday == "" {
		recruiter.Birthday = cloneRecruiter.Birthday

	}
	if recruiter.Gender == "" {
		recruiter.Gender = cloneRecruiter.Gender

	}
	if recruiter.CompanyID == "" {
		recruiter.CompanyID = cloneRecruiter.CompanyID.ID

	}

	query = `UPDATE recruiter SET name = $1, email = $2, phone_number = $3, birthday=$4, gender=$5, updated_at = now() 
	WHERE id = $6 AND deleted_at = 0`
	_, err = rm.con.Exec(query, recruiter.Name, recruiter.Email, recruiter.PhoneNumber, recruiter.Birthday, recruiter.Gender, recruiter.ID)
	if err != nil {
		slog.Error("Update metodi querysida sintaksis xato", "err:",  err)
		return err
	}
	return nil
}

func (rm *RecruiterManager) Delete(recruiterID string) error {
	query = `UPDATE recruiter SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err = rm.con.Exec(query, recruiterID)
	if err != nil {
		return err
	}
	return nil
}
