package api_P

import (
	hp "root/api_P/handler_P"
	ps "root/postgres_P"

	"github.com/gin-gonic/gin"
)

func NewGin(um *ps.UserManager, rm *ps.ResumeManager, rcm *ps.RecruiterManager, cm *ps.CompanyManager, vm *ps.VacancyManager, im *ps.InterviewManager) *gin.Engine {
	r := gin.Default()

	handler := hp.NewHandler(um, rm, rcm, cm, vm, im)

	r.POST("/user", handler.UserCreateHandler)
	r.GET("/user/:id", handler.UserGetByIDHandler)
	r.GET("/user", handler.UserGetAllFilterHandler)
	r.PUT("/user", handler.UserUpdateHandler)
	r.DELETE("/user/:id", handler.UserDeleteHandler)
	r.GET("/user/:id/myinterview", handler.GetAllUser_Interview_Handler)
	r.GET("/user/:id/myresume", handler.GetAllUser_Resume_Handler)

	r.POST("/resume", handler.ResumeCreateHandler)
	r.GET("/resume/:id", handler.ResumeGetByIDHandler)
	r.GET("/resume", handler.ResumeGetAllHandler)
	r.PUT("/resume", handler.ResumeUpdateHandler)
	r.DELETE("/resume/:id", handler.ResumeDeleteHandler)

	r.POST("/recruiter", handler.RecriuterCreateHandler)
	r.GET("/recruiter/:id", handler.RecriuterGetByIDHandler)
	r.GET("/recruiter", handler.RecriuterGetAllHandler)
	r.PUT("/recruiter", handler.RecriuterUpdateHandler)
	r.DELETE("/recruiter/:id", handler.RecriuterDeleteHandler)

	r.POST("/vacancy", handler.VacancyCreateHandler)
	r.GET("/vacancy/:id", handler.VacancyGetByIDHandler)
	r.GET("/vacancy", handler.VacancyGetAllHandler)
	r.PUT("/vacancy", handler.VacancyUpdateHandler)
	r.DELETE("/vacancy/:id", handler.VacancyDeleteHandler)

	r.POST("/company", handler.CompanyCreateHandler)
	r.GET("/company/:id", handler.CompanyGetByIDHandler)
	r.GET("/company", handler.CompanyGetAllHandler)
	r.PUT("/company", handler.CompanyUpdateHandler)
	r.DELETE("/company/:id", handler.CompanyDeleteHandler)

	r.POST("/interview", handler.InterviewCreateHandler)
	r.GET("/interview/:id", handler.InterviewGetByIDHandler)
	r.GET("/interview/company", handler.InterviewGetAllHandler)
	r.PUT("/interview", handler.InterviewUpdateHandler)
	r.DELETE("/interview/:id", handler.InterviewDeleteHandler)

	

	return r
}
