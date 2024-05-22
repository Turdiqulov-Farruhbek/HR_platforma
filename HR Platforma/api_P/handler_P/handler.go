package handler_P

import (
	ps "root/postgres_P"
)

type Handler struct {
	UserRepo   *ps.UserManager
	ResumeRepo *ps.ResumeManager
	RecruiterRepo *ps.RecruiterManager
	CompanyRepo *ps.CompanyManager
	VacancyRepo *ps.VacancyManager
	InterviewRepo *ps.InterviewManager
}

func NewHandler(um *ps.UserManager, rm *ps.ResumeManager, rcm *ps.RecruiterManager, 
	cm  *ps.CompanyManager, vm *ps.VacancyManager, im *ps.InterviewManager) *Handler {
	return &Handler{UserRepo: um, ResumeRepo: rm, RecruiterRepo: rcm, CompanyRepo: cm, VacancyRepo: vm, InterviewRepo: im}
}
