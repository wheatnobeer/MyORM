package MyORM

import (
	"MyORM/session"
	"testing"
)

var (
	user1 = &user{"Tom", 18}
	user2 = &user{"Sam", 25}
	user3 = &user{"Jack", 25}
)

func testRecordInit(t *testing.T) (*session.Session, *Engine) {
	t.Helper()
	engine, _ := NewEngine("mysql", "root:020408@tcp(127.0.0.1:3306)/gee")
	s := engine.NewSession().Model(&user{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s, engine
}

func TestSession_Insert(t *testing.T) {
	s, engine := testRecordInit(t)
	defer engine.Close()

	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s, engine := testRecordInit(t)
	defer engine.Close()

	var users []user
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
}
