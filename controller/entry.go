package controller

import (
	"diary_api/database"
	"diary_api/helper"
	"diary_api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddEntry(context *gin.Context) {
	var input model.Entry
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	input.UserId = user.ID

	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntry(context *gin.Context) {
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

func DeleteEntry(context *gin.Context) {
	entryId := context.Param("id")
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var entry model.Entry
	err = database.Database.Where("id = ? AND user_id = ?", entryId, user.ID).First(&entry).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = database.Database.Delete(&entry).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateEntry(context *gin.Context) {
	entryId := context.Param("id")
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var entry model.Entry
	err = database.Database.Where("id = ? AND user_id = ?", entryId, user.ID).First(&entry).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var input model.Entry
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry.Content = input.Content

	err = database.Database.Save(&entry).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": entry})
}
