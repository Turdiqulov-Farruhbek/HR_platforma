package handler_P

import (
	"fmt"
	"log"
	"net/http"
	sp "root/structs_P"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CompanyCreateHandler(ctx *gin.Context) {
	var company *sp.CompanyCreate
	err := ctx.ShouldBindJSON(&company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		log.Fatal(err)
		return
	}
	err = h.CompanyRepo.Create(company)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorCreate": err.Error()})
		log.Fatal(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Company created successfully"})
}

func (h *Handler) CompanyGetByIDHandler(ctx *gin.Context) {
	companyID := ctx.Param("id")

	company, err := h.CompanyRepo.GetByID(companyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"company": company})
}

func (h *Handler) CompanyGetAllHandler(ctx *gin.Context) {
	location := ctx.Query("location")
	companies, err := h.CompanyRepo.GetAll(location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"companies": companies})
}

func (h *Handler) CompanyUpdateHandler(ctx *gin.Context) {
	var company *sp.CompanyUpdate
	err := ctx.ShouldBindJSON(&company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		log.Fatal(err)
		return
	}
	
	fmt.Println(company)
	err = h.CompanyRepo.Update(company)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorUpdate": err.Error()})
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}

func (h *Handler) CompanyDeleteHandler(ctx *gin.Context) {
	companyID := ctx.Param("id")
	err := h.CompanyRepo.Delete(companyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorDelete": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
