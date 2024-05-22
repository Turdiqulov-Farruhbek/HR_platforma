package postgres_P

import (
	"database/sql"
	"fmt"
	"log/slog"
	"root/structs_P"
)

type CompanyManager struct {
	con *sql.DB
}

func NewCompanyManager(conn *sql.DB) *CompanyManager {
	return &CompanyManager{con: conn}
}

func (cm *CompanyManager) Create(company *structs_P.CompanyCreate) error {
	query = `INSERT INTO company (name, location, workers) VALUES ($1, $2, $3)`
	_, err = cm.con.Exec(query, company.Name, company.Location, company.Workers)
	if err != nil {
		return err
	}
	return nil
}

func (cm *CompanyManager) GetByID(companyID string) (*structs_P.Company, error) {
	company := structs_P.Company{}
	query = `SELECT c.name, c.location, c.workers FROM company c 
	WHERE c.id = $1 AND c.deleted_at=0`
	err = cm.con.QueryRow(query, companyID).Scan(&company.ID, &company.Name, &company.Location, &company.Workers)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (cm *CompanyManager) GetAll(location string) (*[]structs_P.Company, error) {
	companies := []structs_P.Company{}
	query = `SELECT c.name, c.location, c.workers FROM company c
	WHERE c.deleted_at=0`

	args := []interface{}{}
	paramIndex := 1

	if location != "" {
		query += fmt.Sprintf(" AND location = $%d", paramIndex)
		args = append(args, location)
		paramIndex++
	}

	rows, err := cm.con.Query(query, args...)
	if err != nil {
		slog.Error("Query", "err:", err)
		return nil, err
	}
	for rows.Next() {
		company := structs_P.Company{}
		err = rows.Scan(&company.Name, &company.Location, &company.Workers)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	return &companies, nil
}

func (cm *CompanyManager) Update(company *structs_P.CompanyUpdate) error {
	cloneCompany, err := cm.GetByID(company.ID)
	if err != nil {
		return err
	}

	if company.Name == "" {
		company.Name = cloneCompany.Name
	}
	if company.Location == "" {
		company.Location = cloneCompany.Location
	}
	if company.Workers == 0 {
		company.Workers = cloneCompany.Workers
	}

	query = `UPDATE company SET name = $1, location = $2, workers = $3, updated_at = NOW() WHERE id = $4`
	_, err = cm.con.Exec(query, company.Name, company.Location, company.Workers, company.ID)
	if err != nil {
		slog.Error("Error updating company", "err", err)
		return err
	}
	return nil
}

func (cm *CompanyManager) Delete(companyID string) error {
	query = `UPDATE company SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err = cm.con.Exec(query, companyID)
	if err != nil {
		return err
	}
	return nil
}
