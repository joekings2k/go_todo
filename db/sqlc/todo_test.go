package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go_todos/util"
	"github.com/stretchr/testify/require"
)

type NullString struct {
	Valid bool
	String string
}

func createRandomTodo (t *testing.T) Todo {
	user := createRandomUser(t)
	arg := CreateTodoParams{
    Title: util.RandomTodoTitle(),
    Description: sql.NullString{String:"This is a random todo description that explains what the todo is all about.", Valid: true},
		UserID:user.ID,
		Status: sql.NullString{String:util.RandomTodoStatus(), Valid: true},
	}
	todo, err := testQueries.CreateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)
	require.Equal(t, arg.Title, todo.Title)
	require.Equal(t, arg.Description, todo.Description)
	require.Equal(t, arg.UserID, todo.UserID)
	require.Equal(t, arg.Status, todo.Status)
	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreatedAt)
	return todo
}


func TestCreateTodo(t *testing.T) {
	createRandomTodo(t)
}