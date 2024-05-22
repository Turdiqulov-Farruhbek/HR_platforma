package postgres_P

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"root/structs_P"
	"strings"
)

type InterviewManager struct {
	con *sql.DB
}

func NewInterviewManager(conn *sql.DB) *InterviewManager {
	return &InterviewManager{con: conn}
}

func (im *InterviewManager) Create(interview *structs_P.InterviewCreate) error {
	var cheak bool
	var userAge int
	err := im.con.QueryRow("SELECT EXTRACT(YEAR FROM AGE(birthday)) FROM users WHERE id = $1", &interview.UserID.ID).Scan(&userAge)
	if err != nil {
		panic(err)
	}

	if userAge >= 18 {
		cheak = true
	}

	if !cheak {
		err = fmt.Errorf("yoshingiz %s", "mos kelmadi")
		return err
	}

	position, err := im.con.Query(`SELECT r.position FROM resume r 
	JOIN vacancy v on r.position = v.position
	WHERE  r.user_id=$1 AND v.id=$2 AND v.deleted_at=0 AND r.deleted_at=0`, interview.UserID.ID, interview.VacancyID.ID)
	if err != nil {
		slog.Error("Greate metodi querysida xato:", "err", err)
		return err
	}

	var resumePosition string
	for position.Next() {
		position.Scan(&resumePosition)
		cheak = true
	}
	if !cheak {
		err = fmt.Errorf("position: %s", "mos kelmadi")
		return err
	}
	defer position.Close()
	fmt.Print(resumePosition)

	_, err = im.con.Exec(`INSERT INTO interview(user_id, vacancy_id, recruiter_id, interview_date) VALUES($1, $2, $3, now())`,
		interview.UserID.ID,
		interview.VacancyID.ID,
		interview.RecruiterID.ID,
	)
	if err != nil {
		slog.Error("Greate metodi querysida xato:", "err", err)
		return err
	}
	return err
}

func (im *InterviewManager) GetByID(interviewID string) (*structs_P.Interview, error) {
	var interview structs_P.Interview
	row := im.con.QueryRow(`SELECT id, user_id, vacancy_id, recruiter_id, interview_date, created_at, updated_at, deleted_at 
	FROM interview 
	WHERE id=$1 AND deleted_at=0`, interviewID)
	err := row.Scan(&interview.ID, &interview.UserID, &interview.VacancyID, &interview.RecruiterID,
		&interview.InterviewDate, &interview.CreatedAt, &interview.UpdatedAt, &interview.DeletedAt)
	if err != nil {
		slog.Error("Error getting", err)
		return nil, err
	}
	return &interview, nil
}

func (im *InterviewManager) GetAll(compID string) (*[]structs_P.InterviewAll, error) {
	var interviews []structs_P.InterviewAll
	query := `
	SELECT 
		v.id AS vacancy_id, v.name AS vacancy_name, v.position, v.min_exp, v.company_id, c.name, c.location, c.workers, v.description, 
		i.id AS interview_id, 
		u.id AS user_id, u.name AS user_name, u.email AS user_email, u.phone_number AS user_phone_number, u.birthday AS user_birthday, u.gender AS user_gender,
		r.id AS recruiter_id, r.name AS recruiter_name, r.email AS recruiter_email, r.phone_number AS recruiter_phone_number, r.birthday AS recruiter_birthday, r.gender AS recruiter_gender,
		c.name, c.location, c.workers,
		i.interview_date
	FROM vacancy v
	JOIN interview i ON i.vacancy_id = v.id
	JOIN users u ON i.user_id = u.id
	JOIN recruiter r ON i.recruiter_id = r.id
	JOIN company c ON v.company_id = c.id
	WHERE i.deleted_at = 0 AND v.deleted_at = 0`
	var args []interface{}
	if compID != "" {
		query += " AND v.company_id = $1"
		args = append(args, compID)
	}

	rows, err := im.con.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		interview := structs_P.InterviewAll{}
		err = rows.Scan(
			&interview.VacancyID.ID, &interview.VacancyID.Name, &interview.VacancyID.Position, &interview.VacancyID.MinExp, &interview.VacancyID.CompanyID.ID,
			&interview.VacancyID.CompanyID.Name, &interview.VacancyID.CompanyID.Location, &interview.VacancyID.CompanyID.Workers, &interview.VacancyID.Description,
			&interview.ID,
			&interview.UserID.ID, &interview.UserID.Name, &interview.UserID.Email, &interview.UserID.PhoneNumber, &interview.UserID.Birthday, &interview.UserID.Gender,
			&interview.RecruiterID.ID, &interview.RecruiterID.Name, &interview.RecruiterID.Email, &interview.RecruiterID.PhoneNumber, &interview.RecruiterID.Birthday, &interview.RecruiterID.Gender,
			&interview.RecruiterComp.Name, &interview.RecruiterComp.Location, &interview.RecruiterComp.Workers,
			&interview.InterviewDate,
		)
		if err != nil {
			slog.Error("Error scanning", "err", err)
			return nil, err
		}
		interviews = append(interviews, interview)
	}

	return &interviews, nil
}

func (im *InterviewManager) Update(interview *structs_P.InterviewUpdate) error {
	var cheak bool
	row, err := im.con.Query(`
	SELECT id FROM users 
	WHERE EXTRACT(YEAR FROM AGE(birthday)) > 18
	AND id=$1 AND deleted_at=0`, interview.UserID.ID)
	if err != nil {

		slog.Error("18 dab katta bulish shart", "error", err)
	}

	if !row.Next() {
		err = fmt.Errorf("yosh: %s", "dan katta bolishi kerak")
		slog.Error("chek", err)
		log.Println("err next")

		return err
	}
	log.Println(3)

	position, err := im.con.Query(`
	SELECT r.position FROM resume r 
	JOIN vacancy v on r.position = v.position
	WHERE  r.user_id=$1 AND v.id=$2 AND v.deleted_at=0 AND r.deleted_at=0`,
		interview.UserID.ID, interview.VacancyID.ID)
	if err != nil {
		slog.Error("chek", err)
	}

	var resumePosition string
	for position.Next() {
		position.Scan(&resumePosition)
		cheak = true

	}
	fmt.Println(position)
	if !cheak {
		err = fmt.Errorf("position: %s", "mos kelmadi")
		return err
	}
	defer position.Close()
	fmt.Print(resumePosition)

	query := `UPDATE interview SET`

	var args []interface{}
	paramIndex := 1
	if interview.UserID.ID != "" {
		query += fmt.Sprintf(" user_id = $%d,", paramIndex)
		args = append(args, interview.UserID.ID)
		paramIndex++

	}
	if interview.VacancyID.ID != "" {
		query += fmt.Sprintf(" vacancy_id = $%d,", paramIndex)
		args = append(args, interview.VacancyID.ID)
		paramIndex++

	}
	if interview.RecruiterID.ID != "" {
		query += fmt.Sprintf(" recruiter_id = $%d,", paramIndex)
		args = append(args, interview.RecruiterID.ID)
		paramIndex++

	}
	if interview.InterviewDate != "" {
		query += fmt.Sprintf(" interview_date = $%d,", paramIndex)
		args = append(args, interview.InterviewDate)
		paramIndex++

	}
	query = strings.TrimSuffix(query, ",")
	log.Println(7)
	query += fmt.Sprintf(" WHERE id = $%d", paramIndex)
	args = append(args,interview.ID)
	query += " and deleted_at = 0"

	
	_, err = im.con.Exec(query, args...)
	if err != nil {
		slog.Error("update qilishda posgres papka", err)
	}
	return err
}

func (im *InterviewManager) Delete(interviewID string) error {
	query = `UPDATE interview SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err = im.con.Exec(query, interviewID)
	if err != nil {
		return err
	}
	return nil
}
