package MyORM

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type user struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {

	engine, _ := NewEngine("mysql", "root:020408@tcp(127.0.0.1:3306)/gee")
	defer engine.Close()
	s := engine.NewSession().Model(&user{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
