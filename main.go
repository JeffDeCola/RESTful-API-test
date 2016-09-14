package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffdecola/RESTful-API-test/http"
	_ "github.com/mattn/go-sqlite3"
)

// The database
// var dbmap = initDb()

func main() {

	var dbmap = http.InitDb()

	defer dbmap.Db.Close()
	a := http.NewApp(dbmap)

	router := gin.Default()
	router.GET("/articles", a.ArticlesList)
	router.POST("/articles", a.ArticlePost)
	router.GET("/articles/:id", a.ArticlesDetail)
	router.Run(":8000")
}
