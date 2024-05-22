package postgres_P

import (
	"database/sql"
	"fmt"
	"root/structs_P"
)

type VacancyManager struct {
	con *sql.DB
}

func NewVacansyManager(conn *sql.DB) *VacancyManager {
	return &VacancyManager{con: conn}
}

func (vm *VacancyManager) Create(vacancy *structs_P.VacancyCreated) error {
	query = `INSERT INTO vacancy (name, position, min_exp, company_id, description) VALUES ($1, $2, $3, $4, $5)`
	_, err = vm.con.Exec(query, vacancy.Name, vacancy.Position, vacancy.MinExp, vacancy.CompanyID, vacancy.Description)
	if err != nil {
		panic(err)
	}
	return nil
}

func (vm *VacancyManager) GetByID(vacancyID string) (*structs_P.VacancyAll, error) {
	vacancy := structs_P.VacancyAll{}
	query = `SELECT v.id, v.name, v.position, v.min_exp, v.description,  v.created_at, c.name, c.location, c.workers
	FROM vacancy v 
	JOIN company c ON v.company_id = c.id 
	WHERE v.id = $1`
	err = vm.con.QueryRow(query, vacancyID).Scan(&vacancy.ID, &vacancy.Name, &vacancy.Position, &vacancy.MinExp, &vacancy.Description, &vacancy.CreatedAt,
		&vacancy.CompanyID.Name, &vacancy.CompanyID.Location, &vacancy.CompanyID.Workers)
	if err != nil {
		return nil, err
	}
	return &vacancy, nil
}

func (vm *VacancyManager) GetAllFilterAll(position, minExp, companyId string) ([]*structs_P.VacancyAll, error) {
	var vacancies []*structs_P.VacancyAll
	query = `SELECT v.id, v.name, v.position, v.min_exp, v.description,  v.created_at, c.name, c.location, c.workers
	FROM vacancy v 
	JOIN company c ON v.company_id = c.id 
	WHERE v.deleted_at=0 AND c.deleted_at=0`

	args := []interface{}{}

	paramIndex := 1

	if position != "" {
		query += fmt.Sprintf(" AND position = $%d", paramIndex)
		args = append(args, position)
		paramIndex++
	}
	if minExp != "" {
		query += fmt.Sprintf(" AND min_exp = $%d", paramIndex)
		args = append(args, minExp)
		paramIndex++
	}
	if companyId != "" {
		query += fmt.Sprintf(" AND company_id = $%d", paramIndex)
		args = append(args, companyId)
		paramIndex++
	}

	rows, err := vm.con.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		vacancy := structs_P.VacancyAll{}
		err = rows.Scan(&vacancy.ID, &vacancy.Name, &vacancy.Position, &vacancy.MinExp, &vacancy.Description, &vacancy.CreatedAt,
			&vacancy.CompanyID.Name, &vacancy.CompanyID.Location, &vacancy.CompanyID.Workers)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, &vacancy)
	}
	fmt.Println("vacance", vacancies)
	return vacancies, nil
}

func (vm *VacancyManager) Update(vacancy *structs_P.VacancyUpdate) error {
	cloneVacancy := structs_P.VacancyUpdate{}
	query := `SELECT v.name, v.position, v.min_exp, v.description, company_id FROM vacancy v WHERE v.id=$1`
	err := vm.con.QueryRow(query, vacancy.ID).Scan(&cloneVacancy.Name, &cloneVacancy.Position, &cloneVacancy.MinExp,
		&cloneVacancy.Description, &cloneVacancy.CompanyID)
	if err!= nil {
	    return err
	}

	if vacancy.Name == "" {
		vacancy.Name = cloneVacancy.Name
	}
	if vacancy.Position == "" {
		vacancy.Position = cloneVacancy.Position
	}
	if vacancy.MinExp == -1 {
		vacancy.MinExp = cloneVacancy.MinExp
	}
	if vacancy.CompanyID == "" {
		vacancy.CompanyID = cloneVacancy.CompanyID
	}
	if vacancy.Description == "" {
		vacancy.Description = cloneVacancy.Description
	}

	query = `UPDATE vacancy SET name = $1, position = $2, min_exp = $3, company_id = $4, description = $5, updated_at = NOW() WHERE id = $6`
	_, err = vm.con.Exec(query, vacancy.Name, vacancy.Position, vacancy.MinExp, vacancy.CompanyID, vacancy.Description, vacancy.ID)
	if err != nil {
		return err
	}
	return nil
}

func (vm *VacancyManager) Delete(vacancyID string) error {
	query = `UPDATE vacancy SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err = vm.con.Exec(query, vacancyID)
	if err != nil {
		return err
	}

	return nil
}
