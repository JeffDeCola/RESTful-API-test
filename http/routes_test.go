// RESTful-API-test routes_test.go

package http

import (
	"reflect"
	"testing"

	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
)

func TestNewApp(t *testing.T) {
	type args struct {
		db *gorp.DbMap
	}
	tests := []struct {
		name string
		args args
		want *App
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewApp(tt.args.db); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewApp() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestApp_createArticle(t *testing.T) {
	type args struct {
		title string
		body  string
	}
	tests := []struct {
		name string
		a    *App
		args args
		want Article
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.a.createArticle(tt.args.title, tt.args.body); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. App.createArticle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestApp_getArticle(t *testing.T) {
	type args struct {
		article_id int
	}
	tests := []struct {
		name string
		a    *App
		args args
		want Article
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.a.getArticle(tt.args.article_id); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. App.getArticle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestApp_ArticlesList(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.a.ArticlesList(tt.args.c)
	}
}

func TestApp_ArticlesDetail(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.a.ArticlesDetail(tt.args.c)
	}
}

func TestApp_ArticlePost(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.a.ArticlePost(tt.args.c)
	}
}

func TestInitDb(t *testing.T) {
	tests := []struct {
		name string
		want *gorp.DbMap
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := InitDb(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. InitDb() = %v, want %v", tt.name, got, tt.want)
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
