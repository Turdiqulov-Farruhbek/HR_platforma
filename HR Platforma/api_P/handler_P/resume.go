package handler_P

import (
	"log"
	"net/http"
	sp "root/structs_P"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ResumeCreateHandler(ctx *gin.Context) {
	var user *sp.ResumeCreate
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		log.Fatal(err)
		return
	}

	err = h.ResumeRepo.Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorCreate": err.Error()})
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) ResumeGetByIDHandler(ctx *gin.Context) {
	resumeID := ctx.Param("id")

    resume, err := h.ResumeRepo.GetByID(resumeID)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        log.Fatal(err)
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"resume": resume})
}

func (h *Handler) ResumeGetAllHandler(ctx *gin.Context) {

	position := ctx.Query("position")
	experience := ctx.Query("experience")
	from := ctx.Query("from")
	to := ctx.Query("to")

    resume, err := h.ResumeRepo.GetAllFiltrResume(position, experience, from, to)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        log.Fatal(err)
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"resume": resume})
}

func (h *Handler) ResumeUpdateHandler(ctx *gin.Context) {
	var resume *sp.ResumeUpdate
    err := ctx.ShouldBindJSON(&resume)
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
        log.Fatal(err)
        return
    }

    err = h.ResumeRepo.Update(resume)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errorUpdate": err.Error()})
        log.Fatal(err)
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}


func (h *Handler) ResumeDeleteHandler(ctx *gin.Context) {
	resumeID := ctx.Param("id")
    err := h.ResumeRepo.Delete(resumeID)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errorDelete": err.Error()})
        log.Fatal(err)
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

