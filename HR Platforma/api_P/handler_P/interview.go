package handler_P

import (
	"log"
	"log/slog"
	"net/http"
	sp "root/structs_P"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (h *Handler) InterviewCreateHandler(ctx *gin.Context) {
	var inter sp.Interv

	err := ctx.ShouldBindJSON(&inter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		log.Fatal(err)
		return
	}

	var interview sp.InterviewCreate
	interview.UserID.ID = inter.UserID
	interview.VacancyID.ID = inter.VacancyID
	interview.RecruiterID.ID = inter.RecruiterID

	err = h.InterviewRepo.Create(&interview)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorCreate": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) InterviewGetByIDHandler(ctx *gin.Context) {
	interviewID := ctx.Param("id")

	interview, err := h.InterviewRepo.GetByID(interviewID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"interview": interview})
}

func (h *Handler) InterviewUpdateHandler(ctx *gin.Context) {
	var interv sp.Interv
	err := ctx.ShouldBindJSON(&interv)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		slog.Error("Yangilanishda body Xatosi", err)
		return
	}
	var interview sp.InterviewUpdate
	interview.ID = interv.ID
	interview.UserID.ID = interv.UserID
	interview.VacancyID.ID = interv.VacancyID
	interview.RecruiterID.ID = interv.RecruiterID

	pp.Println(interv)
	pp.Println(interview)


	err = h.InterviewRepo.Update(&interview)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorUpdate": err.Error()})
		slog.Error("Yangilanishda update Xatosi", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *Handler) InterviewGetAllHandler(ctx *gin.Context) {
	compID := ctx.Query("id")

	interviews, err := h.InterviewRepo.GetAll(compID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"interviews": interviews})
}

func (h *Handler) InterviewDeleteHandler(ctx *gin.Context) {
	interviewID := ctx.Param("id")
	err := h.InterviewRepo.Delete(interviewID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorDelete": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
