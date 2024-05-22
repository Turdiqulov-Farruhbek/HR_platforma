package main

import (
	"database/sql"
	"fmt"
	api "root/api_P"
	cp "root/config_P"
	ps "root/postgres_P"

	_ "github.com/lib/pq"
)

func main() {
	cfg := cp.Load()

	dbCon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)
	fmt.Println(dbCon)
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	userMN := ps.NewUserManager(db)
	resumeMN := ps.NewResumeManager(db)
	recruiterMN := ps.NewRecruiterManager(db)
	companyMN := ps.NewCompanyManager(db)
	vacancyMN := ps.NewVacansyManager(db)
	interviewMN := ps.NewInterviewManager(db)
	
	r := api.NewGin(userMN, resumeMN, recruiterMN, companyMN, vacancyMN, interviewMN)
	defer db.Close()
	r.Run(":8080")

}
