package session

import "testing"

type User struct {
	Name string `orm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HashTable() {
		t.Fatal("Failed to create table User")
	}
}

func TestModel(t *testing.T) {
	s := NewSession().Model(&User{})
	table := s.RefTable()
	s.Model(&Session{})
	if table.Name != "User" || s.refTable.Name != "Session" {
		t.Fatal("Failed to change model")
	}
}