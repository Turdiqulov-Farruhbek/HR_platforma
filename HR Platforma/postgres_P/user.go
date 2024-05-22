package postgres_P

import (
	"database/sql"
	"fmt"
	"root/structs_P"
	"log/slog"
)

var (
	query string
	err   error
)

type UserManager struct {
	con *sql.DB
}

func NewUserManager(conn *sql.DB) *UserManager {
	return &UserManager{con: conn}
}

func (um *UserManager) Create(user *structs_P.UserCreate) error {
	query = `INSERT INTO users (name, email, phone_number, birthday, gender) VALUES ($1, $2, $3, $4, $5)`
	_, err = um.con.Exec(query, user.Name, user.Email, user.PhoneNumber, user.Birthday, user.Gender)
	if err != nil {
		slog.Error("Create metodi querysida sintaksis xato", "err:",  err)
		return err
	}
	return nil
}

func (um *UserManager) GetByID(userID string) (*structs_P.User, error) {
	user := structs_P.User{}
	query = `SELECT id, name, email, phone_number, birthday, gender, created_at, updated_at, deleted_at FROM users WHERE id = $1 `
	err = um.con.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Birthday, &user.Gender, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		slog.Error("GetById metodi querysida sintaksis xato", "err:",  err)
		return nil, err
	}
	return &user, nil
}

func (um *UserManager) GetAllFilterAll(gender, from, to string) ([]*structs_P.User, error) {
	var users []*structs_P.User
	query := `
		  SELECT id, name, email, phone_number, birthday, gender, created_at, updated_at, deleted_at
		  myResume users
		  WHERE deleted_at = 0
	  `
	args := []interface{}{}

	paramIndex := 1

	if gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", paramIndex)
		args = append(args, gender)
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

	rows, err := um.con.Query(query, args...)
	if err != nil {
		slog.Error("GetAllFilterAll metodi querysida sintaksis xato", "err:",  err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := structs_P.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Birthday, &user.Gender, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			slog.Error("GetAllFilterAll metodi Sacanida xato", "err:",  err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (um UserManager) GetAllUser_Interview(userID string) ([]structs_P.InterviewAll, error) {
	var interviews []structs_P.InterviewAll

	query := `SELECT 
		i.vacancy_id, i.recruiter_id, i.interview_date,
	    u.id, u.name, u.email, u.phone_number, u.birthday, u.gender,
	    v.name, v.position, v.min_exp, c.id, c.name, c.location, c.workers, v.description,
	    r.name, r.email, r.phone_number, r.birthday, r.gender,
		c.id, c.name, c.location, c.workers
		
	FROM interview i
	JOIN users u ON i.user_id = u.id
	JOIN vacancy v ON i.vacancy_id = v.id
	JOIN recruiter r ON i.recruiter_id = r.id
	JOIN company c ON v.company_id = c.id
	WHERE i.deleted_at = 0 AND u.deleted_at = 0 AND v.deleted_at = 0 AND r.deleted_at = 0 AND c.deleted_at=0 AND i.user_id = $1`

	rows, err := um.con.Query(query, userID)
	if err != nil {
		slog.Error("Queryda sintaksis xato", "err:",  err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var interv structs_P.InterviewAll

		if err := rows.Scan(
			&interv.VacancyID.ID, &interv.RecruiterID.ID, &interv.InterviewDate,
			&interv.UserID.ID, &interv.UserID.Name, &interv.UserID.Email, &interv.UserID.PhoneNumber, &interv.UserID.Birthday, &interv.UserID.Gender,
			&interv.VacancyID.Name, &interv.VacancyID.Position, &interv.VacancyID.MinExp,
			&interv.VacancyID.CompanyID.ID, &interv.VacancyID.CompanyID.Name, &interv.VacancyID.CompanyID.Location, &interv.VacancyID.CompanyID.Workers, &interv.VacancyID.Description,
			&interv.RecruiterID.Name, &interv.RecruiterID.Email, &interv.RecruiterID.PhoneNumber, &interv.RecruiterID.Birthday, &interv.RecruiterID.Gender,
			&interv.RecruiterComp.ID, &interv.RecruiterComp.Name, &interv.RecruiterComp.Location, &interv.RecruiterComp.Workers); err != nil {
			return nil, err
		}
		interviews = append(interviews, interv)
	}
	return interviews, nil
}

func (um UserManager) GetAllUser_Resume(userID string) (*[]structs_P.ResumeAll, error) {
	var resums []structs_P.ResumeAll

	query := `SELECT 
		r.id, r.position, r.experience, r.description,
		u.id, u.name, u.email, u.phone_number, u.birthday, u.gender	
	FROM resume as r
	JOIN users u ON r.user_id = u.id
	WHERE r.deleted_at = 0 AND u.deleted_at = 0 AND r.user_id = $1`

	rows, err := um.con.Query(query, userID)
	if err != nil {
		slog.Error("Queryda sintaksis xato", "err:",  err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var resum structs_P.ResumeAll
		if err := rows.Scan(
			&resum.ID, &resum.Position, &resum.Experience, &resum.Description,
			&resum.UserID.ID, &resum.UserID.Name, &resum.UserID.Email, &resum.UserID.PhoneNumber, &resum.UserID.Birthday, &resum.UserID.Gender); err != nil {
			return nil, err
		}
		resums = append(resums, resum)
	}

	return &resums, nil

}

func (um *UserManager) Update(user *structs_P.UserUpdate) error {
	cloneUser, err := um.GetByID(user.ID)
	if err != nil {
		return err
	}

	if user.Name == "" {
		user.Name = cloneUser.Name
	}
	if user.Birthday == "" {
		user.Birthday = cloneUser.Birthday
	}
	if user.Gender == "" {
		user.Gender = cloneUser.Gender
	}
	if user.Email == "" {
		user.Email = cloneUser.Email
	}
	if user.PhoneNumber == "" {
		user.PhoneNumber = cloneUser.PhoneNumber
	}

	query = `UPDATE users SET name = $1, email = $2, phone_number = $3, birthday = $4, gender = $5, updated_at = now() WHERE id = $6 AND deleted_at = 0`
	_, err = um.con.Exec(query, user.Name, user.Email, user.PhoneNumber, user.Birthday, user.Gender, user.ID)
	if err != nil {
		slog.Error("Update metodi querysida sintaksis xato", "err:",  err)
		return err

	}
	return nil
}

func (um *UserManager) Delete(userID string) error {
	query = `UPDATE users SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err = um.con.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}
