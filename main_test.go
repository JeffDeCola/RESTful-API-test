package main

import (
	"reflect"
	"testing"

	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}

func Test_createArticle(t *testing.T) {
	type args struct {
		title string
		body  string
	}
	tests := []struct {
		name string
		args args
		want Article
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := createArticle(tt.args.title, tt.args.body); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. createArticle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getArticle(t *testing.T) {
	type args struct {
		article_id int
	}
	tests := []struct {
		name string
		args args
		want Article
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getArticle(tt.args.article_id); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getArticle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestArticlesList(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		ArticlesList(tt.args.c)
	}
}

func TestArticlesDetail(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		ArticlesDetail(tt.args.c)
	}
}

func TestArticlePost(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		ArticlePost(tt.args.c)
	}
}

func Test_initDb(t *testing.T) {
	tests := []struct {
		name string
		want *gorp.DbMap
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := initDb(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. initDb() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_checkErr(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		checkErr(tt.args.err, tt.args.msg)
	}
}
