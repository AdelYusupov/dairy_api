package controller

import (
	"diary_api/database"
	"diary_api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(context *gin.Context) {
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.Database.Delete(&user).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}
func UpdateUser(context *gin.Context) {
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input struct {
		OldPassword string `json:"old_password" binding:"required"`
		Username    string `json:"username"`
		NewPassword string `json:"new_password"`
	}
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.OldPassword)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Username != "" {
		user.Username = input.Username
	}
	if input.NewPassword != "" {
		user.Password = input.NewPassword
	}
	err = database.Database.Save(&user).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User updated"})
}
