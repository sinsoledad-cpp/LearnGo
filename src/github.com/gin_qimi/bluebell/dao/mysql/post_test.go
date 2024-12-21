package mysql

import (
	"testing"

	"github.com/gin_qimi/bluebell/models"
	"github.com/gin_qimi/bluebell/setting"
)

func init() {
	dbCfg := setting.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "sql3.1415",
		DbName:       "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
