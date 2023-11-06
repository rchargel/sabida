package main

import (
	"os"
	"log"
	
	"github.com/gin-gonic/gin"
	"github.com/rchargel/sabida/dao"
	"github.com/rchargel/sabida/handlers"
)

var router *gin.Engine

func GetEnv(variable string) (string) {
	v := os.Getenv(variable)
	if v == nil {
		log.Panicf("Missing environment variable %s", variable)
	}
	return v
}

func main() {
	db := connectToDatabase()
	runMigrations(db)

	conn := dao.New(db)

	indexHandler := handlers.NewIndexHandler(conn)
	loginHandler := handlers.NewLoginHandler(conn)

	router := gin.Default()

	router.Static("/static", "./static/")
	router.LoadHTMLGlob("templates/*")

	router.POST("/login", loginHandler.ProcessLoginForm)
	router.GET("/", indexHandler.ShowIndexPage)
	//router.GET("/article/view/:article_id", handlers.GetArticle)

	router.Run()
}
