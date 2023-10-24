package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rchargel/sabida/dao"
	"github.com/rchargel/sabida/handlers"
)

var router *gin.Engine

func main() {
	db := connectToDatabase()
	runMigrations(db)

	conn := dao.New(db)

	indexHandler := handlers.NewIndexHandler(conn)

	router := gin.Default()

	router.Static("/static", "./static/")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", indexHandler.ShowIndexPage)
	//router.GET("/article/view/:article_id", handlers.GetArticle)

	router.Run()
}
