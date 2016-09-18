// RESTful-API-test main.go

package main

import (
	"github.com/JeffDeCola/RESTful-API-test/http"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

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
