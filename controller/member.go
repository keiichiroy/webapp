package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsunemisan88/webapp/lib"
	"github.com/tsunemisan88/webapp/models"
)

func Home(c *gin.Context) {
	var members []models.Member
	var err error

	err = lib.Sess.Table("member").Find(&members).Error

	c.JSON(http.StatusOK, gin.H{"error": err, "members": members})
}
