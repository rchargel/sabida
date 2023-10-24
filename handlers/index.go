package handlers

import (
	"context"
	"log"

	"github.com/rchargel/sabida/dao"

	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
	Conn *dao.Queries
}

func NewIndexHandler(conn *dao.Queries) *IndexHandler {
	return &IndexHandler{conn}
}

func (i *IndexHandler) ShowIndexPage(c *gin.Context) {
	ctx := context.Background()
	categories, err := i.Conn.ListCategories(ctx)
	if err != nil {
		log.Panic(err)
	}

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Sabida",
		"payload": categories,
	}, "index.html")

}
