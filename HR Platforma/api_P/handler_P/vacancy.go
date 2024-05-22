package handler_P

import (
	"net/http"
	sp "root/structs_P"

	"github.com/gin-gonic/gin"
)

func (h *Handler) VacancyCreateHandler(ctx *gin.Context) {
	var vacancy *sp.VacancyCreated
	err := ctx.ShouldBindJSON(&vacancy)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		return
	}

	err = h.VacancyRepo.Create(vacancy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorCreate": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vacancy created successfully"})
}


func (h *Handler) VacancyGetByIDHandler(ctx *gin.Context) {
	vacancyID := ctx.Param("id")

    vacancy, err := h.VacancyRepo.GetByID(vacancyID)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"vacancy": vacancy})
}


func (h *Handler) VacancyGetAllHandler(ctx *gin.Context) {
	position := ctx.Query("position")
    minExp := ctx.Query("minExp")
    companyId := ctx.Query("companyId")

    vacancies, err := h.VacancyRepo.GetAllFilterAll(position, minExp, companyId)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errorGetall": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"vacancies": vacancies})
}


func (h *Handler) VacancyUpdateHandler(ctx *gin.Context) {
	var vacancy *sp.VacancyUpdate
    err := ctx.ShouldBindJSON(&vacancy)
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
        return
    }

    err = h.VacancyRepo.Update(vacancy)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errorUpdate": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Vacancy updated successfully"})
}


func (h *Handler) VacancyDeleteHandler(ctx *gin.Context) {
	vacancyID := ctx.Param("id")
    err := h.VacancyRepo.Delete(vacancyID)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errorDelete": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Vacancy deleted successfully"})
}