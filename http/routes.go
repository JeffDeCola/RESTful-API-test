// RESTful-API-test routes.go

package http

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Id      int64 `db:"article_id"`
	Created int64
	Title   string
	Content string
}

// App is the APP
type App struct {
	db *gorp.DbMap
}

// NewApp is the NEW APP
func NewApp(db *gorp.DbMap) *App {
	return &App{db: db}
}

// createArticle the articles
func (a *App) createArticle(title, body string) Article {
	article := Article{
		Created: time.Now().UnixNano(),
		Title:   title,
		Content: body,
	}

	err := a.db.Insert(&article)
	checkErr(err, "Insert failed - sorry")
	return article
}

// getArticle the articles
func (a *App) getArticle(article_id int) Article {
	article := Article{}
	err := a.db.SelectOne(&article, "select * from articles where article_id=?", article_id)
	checkErr(err, "SelectOne failed")
	return article
}

// ArticlesList lists the articles
func (a *App) ArticlesList(c *gin.Context) {
	var articles []Article
	_, err := a.db.Select(&articles, "select * from articles order by article_id")
	checkErr(err, "Select failed")
	content := gin.H{}
	for k, v := range articles {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)
}

// ArticlesDetail shows the article details
func (a *App) ArticlesDetail(c *gin.Context) {
	article_id := c.Params.ByName("id")
	a_id, _ := strconv.Atoi(article_id)
	article := a.getArticle(a_id)
	content := gin.H{"title": article.Title, "content": article.Content}
	c.JSON(200, content)
}

// ArticlePost posts and article
func (a *App) ArticlePost(c *gin.Context) {
	var json Article

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	article := a.createArticle(json.Title, json.Content)
	if article.Title == json.Title {
		content := gin.H{
			"result":  "Success",
			"title":   article.Title,
			"content": article.Content,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occurred"})
	}
}

// InitDb CREATE A DATABASE
func InitDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(Article{}, "articles").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
