package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func InsertArticleTest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("err : %s", err)
	}
	defer db.Close()

	mdb, err := gorm.Open("mysql", db)
	if err != nil {
		t.Errorf("err : " + err.Error())
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO kumparan.articles").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := InsertArticles(mdb, &Articles{
		Author: "Test",
		Body:   "Test",
	}); err != nil {
		t.Error("failed to insert data" + err.Error())
	}
}
