package controllers

import (
	"fmt"
	"net/http"

	"github.com/aleroxac/alura-golang-gin/database"
	"github.com/aleroxac/alura-golang-gin/models"
	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func List(c *gin.Context) {
	var skills []models.Skill
	database.DB.Find(&skills)

	c.JSON(200, skills)
}

func GetByName(c *gin.Context) {
	var skill models.Skill
	skill_name := c.Params.ByName("name")
	database.DB.Where("sk_name = ?", skill_name).First(&skill)

	if skill.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "skill not found",
		})
		return
	}

	c.JSON(200, skill)
}

func Create(c *gin.Context) {
	var new_skill models.Skill

	if err := c.ShouldBindJSON(&new_skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateSkillData(&new_skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&new_skill)
	c.JSON(http.StatusCreated, new_skill)
}

func Update(c *gin.Context) {
	var updated_skill models.Skill
	skill_name := c.Params.ByName("name")
	database.DB.Where("sk_name = ?", skill_name).Find(&updated_skill)

	if updated_skill.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "skill not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&updated_skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateSkillData(&updated_skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&updated_skill).UpdateColumns(updated_skill)
	c.JSON(200, updated_skill)
}

func Delete(c *gin.Context) {
	var deleted_skill models.Skill
	skill_name := c.Params.ByName("name")

	database.DB.Where("sk_name = ?", skill_name).Find(&deleted_skill)
	if deleted_skill.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "skill not found",
		})
		return
	}

	database.DB.Where("sk_name = ?", skill_name).Delete(&deleted_skill)
	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("skill %s has deleted successfully", skill_name),
	})
}
