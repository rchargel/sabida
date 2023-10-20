package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rchargel/sabida/handlers"
)

var router *gin.Engine

func main() {
	router := gin.Default()

	router.Static("/static", "./static/")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)

	router.Run()
}
