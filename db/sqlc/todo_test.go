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

func createRandomTodosPerUserID (userID int32, t *testing.T) Todo{
arg := CreateTodoParams{
    Title: util.RandomTodoTitle(),
    Description: sql.NullString{String:"This is a random todo description that explains what the todo is all about.", Valid: true},
		UserID:userID,
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

func createRandomTodo(t *testing.T) Todo {
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

func clearTodos(t *testing.T) {
	_,err := testQueries.db.ExecContext(context.Background(), "DELETE FROM todos")
	require.NoError(t, err)
}

func TestCreateTodo(t *testing.T) {
	createRandomTodo(t)
}

func TestDeleteTodo(t *testing.T){
	todo1 := createRandomTodo(t)
	err := testQueries.DeleteTodo(context.Background(), todo1.ID)
	require.NoError(t, err)
}

func TestGetTodo(t *testing.T) {
	todo1 := createRandomTodo(t)
	todo2,err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.NoError(t, err)
	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, todo1.Title, todo2.Title)
	require.Equal(t, todo1.Description, todo2.Description)
	require.Equal(t, todo1.UserID, todo2.UserID)
	require.Equal(t, todo1.Status, todo2.Status)
}

func TestUpdateTodo(t *testing.T) {
	todo1 := createRandomTodo(t)
	arg := UpdateTodoParams{
		ID: todo1.ID,
		Title: util.RandomTodoTitle(),
		Description: sql.NullString{String:"This is a random todo description that updates the todo and explains what the todo is all about.", Valid: true},
		Status: sql.NullString{String:util.RandomTodoStatus(), Valid: true},
	}
	todo2, err := testQueries.UpdateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, arg.Title, todo2.Title)
	require.Equal(t, arg.Description, todo2.Description)
	require.Equal(t, arg.Status, todo2.Status)
}

func TestListTodos(t *testing.T) {
	clearTodos(t)
	for i := 0; i < 10; i++ {
		createRandomTodo(t)
	}

	arg := ListTodosParams{
		Limit: 5,
		Offset: 5,
	}

	todos, err := testQueries.ListTodos(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, todos, 5)
	for i,todo := range todos{
		require.NotEmpty(t, todo)
		require.Equal(t, todo.ID, todos[i].ID)
		require.Equal(t, todo.Title, todos[i].Title)
		require.Equal(t, todo.Description, todos[i].Description)
		require.Equal(t, todo.UserID, todos[i].UserID)
		require.Equal(t, todo.Status, todos[i].Status)
	}
}

func TestListTodosByID(t *testing.T) {
	clearTodos(t)
	user := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createRandomTodosPerUserID(user.ID, t)
	}
	arg := ListTodosByUserIDParams {
		UserID: user.ID,
		Limit: 5,
		Offset: 5,
	}
	todos, err := testQueries.ListTodosByUserID(context.Background(), arg)
	require.NoError(t,err)
	require.NotEmpty(t,todos)
	require.Len(t, todos, 5)

	for i,todo  := range todos{
		require.NotEmpty(t,todo)
		require.NotEmpty(t, todo)
		require.Equal(t, todo.ID, todos[i].ID)
		require.Equal(t, todo.Title, todos[i].Title)
		require.Equal(t, todo.Description, todos[i].Description)
		require.Equal(t, todo.UserID, user.ID)
		require.Equal(t, todo.Status, todos[i].Status)
	}	
}

