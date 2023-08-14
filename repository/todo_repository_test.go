package repository_test

import (
	"testing"

	"github.com/kidus-tiliksew/gohtmx/models"
	"github.com/kidus-tiliksew/gohtmx/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Todo{})

	return db, nil
}

func TestCreate(t *testing.T) {
	db, err := initDb()
	if err != nil {
		t.Fatal(err)
	}

	repo := repository.TodoRepository{DB: db}

	todo := models.Todo{Title: "Test todo", Done: false}
	if err := repo.Create(&todo); err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, todo.ID)
	t.Log(todo)
}

func TestGet(t *testing.T) {
	db, err := initDb()
	if err != nil {
		t.Fatal(err)
	}

	repo := repository.TodoRepository{DB: db}

	todo := models.Todo{Title: "Test todo", Done: false}
	if err := repo.Create(&todo); err != nil {
		t.Error(err)
	}

	res, err := repo.Get(todo.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, todo.ID, res.ID)
}


func TestGetAll(t *testing.T) {
	db, err := initDb()
	if err != nil {
		t.Fatal(err)
	}

	repo := repository.TodoRepository{DB: db}

	todo := models.Todo{Title: "Test todo", Done: false}
	if err := repo.Create(&todo); err != nil {
		t.Error(err)
	}

	res, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, res)
}

func TestUpdate(t *testing.T) {
	db, err := initDb()
	if err != nil {
		t.Fatal(err)
	}

	repo := repository.TodoRepository{DB: db}

	todo := models.Todo{Title: "Test todo", Done: false}
	if err := repo.Create(&todo); err != nil {
		t.Error(err)
		return
	}

	newTitle := "new title"
	todo.Title = newTitle
	if err := repo.Update(&todo); err != nil {
		t.Error(err)
		return
	}

	res, err := repo.Get(todo.ID)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, newTitle, res.Title)
}

func TestDelete(t *testing.T) {
	db, err := initDb()
	if err != nil {
		t.Fatal(err)
	}

	repo := repository.TodoRepository{DB: db}

	todo := models.Todo{Title: "Test todo", Done: false}
	if err := repo.Create(&todo); err != nil {
		t.Error(err)
	}

	if err := repo.Delete(todo.ID); err != nil {
		t.Error(err)
	}
}
