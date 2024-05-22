package handler_P

import (
	"net/http"
	sp "root/structs_P"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserCreateHandler(ctx *gin.Context) {
	var user *sp.UserCreate
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		return
	}

	err = h.UserRepo.Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorCreate": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) UserGetByIDHandler(ctx *gin.Context) {
	userID := ctx.Param("id")

	user, err := h.UserRepo.GetByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) UserGetAllFilterHandler(ctx *gin.Context) {
	gender := ctx.Query("gender")
	from := ctx.Query("from")
	to := ctx.Query("to")
	users, err := h.UserRepo.GetAllFilterAll(gender, from, to)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorGetall": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *Handler) GetAllUser_Interview_Handler(ctx *gin.Context) {
	userId := ctx.Param("id")
	users, err := h.UserRepo.GetAllUser_Interview(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *Handler) GetAllUser_Resume_Handler(ctx *gin.Context){
	userId := ctx.Param("id")
    users, err := h.UserRepo.GetAllUser_Resume(userId)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"users": users})
}


func (h *Handler) UserUpdateHandler(ctx *gin.Context) {
	User := sp.UserUpdate{}

	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorBody": err.Error()})
		return
	}

	err = h.UserRepo.Update(&User)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorUpdate": err.Error()})
		return
	}
}

func (h *Handler) UserDeleteHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	err := h.UserRepo.Delete(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorDelete": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
