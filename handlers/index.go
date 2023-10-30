package handlers

import (
	// "context"
	// "log"
	"database/sql"

	"github.com/rchargel/sabida/dao"

	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
	Conn *dao.Queries
}

func NewIndexHandler(conn *dao.Queries) *IndexHandler {
	return &IndexHandler{conn}
}

func ToNullString(text string) sql.NullString {
	return sql.NullString{
		text,
		true,
	}
}

func (i *IndexHandler) ShowIndexPage(c *gin.Context) {
	// ctx := context.Background()
	// categories, err := i.Conn.ListCategories(ctx)
	// if err != nil {
	// 	log.Panic(err)
	// }
	var categories []dao.Category

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Sabida",
		"payload": categories,
	}, "index.html")

}
