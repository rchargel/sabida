package handlers

import (
	"github.com/rchargel/sabida/models"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Sabida",
		"payload": articles,
	}, "index.html")

}
