package handler_P

import (
	"log"
	"net/http"
	sp "root/structs_P"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RecriuterCreateHandler(ctx *gin.Context) {
	var recruiter *sp.RecruiterCreate
	err := ctx.ShouldBindJSON(&recruiter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		log.Fatal(err)
		return
	}

	err = h.RecruiterRepo.Create(recruiter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorCreate": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) RecriuterGetByIDHandler(ctx *gin.Context) {
	recruiterID := ctx.Param("id")

	recruiter, err := h.RecruiterRepo.GetByID(recruiterID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"recruiter": recruiter})
}

func (h *Handler) RecriuterGetAllHandler(ctx *gin.Context) {
	gender := ctx.Query("gender")
	companyid := ctx.Query("companyId")
	from := ctx.Query("from")
	to := ctx.Query("to")

	recruiter, err := h.RecruiterRepo.GetAllRecruiter(gender, companyid, from, to)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorGetall": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"recruiter": recruiter})
}

func (h *Handler) RecriuterUpdateHandler(ctx *gin.Context) {
	var recruiter *sp.RecruiterUpdate
	err := ctx.ShouldBindJSON(&recruiter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		log.Fatal(err)
		return
	}
	err = h.RecruiterRepo.Update(recruiter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorUpdate": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *Handler) RecriuterDeleteHandler(ctx *gin.Context) {
	recruiterID := ctx.Param("id")
	err := h.RecruiterRepo.Delete(recruiterID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorDelete": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
