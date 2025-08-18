package db

import (
	"context"
	"testing"

	"github.com/go_todos/util"
	"github.com/stretchr/testify/require"
)


func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
		Email: util.RandomEmail(),
	}
	user ,err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func clearUsers(t *testing.T){
_,err := testQueries.db.ExecContext(context.Background(), "DELETE FROM users")
require.NoError(t, err)
}

func TestCreateUser (t *testing.T){
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(),user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.ID, user2.ID)
}

func TestDeleteUser(t *testing.T){
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T){
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(),user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.ID, user2.ID)
}

func TestListUsers (t *testing.T){
	clearUsers(t)
	var createdUsers []User
	for i :=0;i<10;i++{
		user := createRandomUser(t)
		createdUsers = append(createdUsers, user)
	}
	arg := ListUsersParams{
		Limit: 5,
		Offset: 5,
	}
	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)
	for i,user := range users{
		require.NotEmpty(t, user)
		require.NotZero(t, user.ID)
		require.Equal(t, createdUsers[5+i].Username, user.Username)
		require.Equal(t, createdUsers[5+i].Password, user.Password)
		require.Equal(t, createdUsers[5+i].Email, user.Email)
	}

	// Boundary case 
	arg = ListUsersParams{
		Limit:5,
		Offset:20,
	}
	users,err = testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 0)

	arg= ListUsersParams{
		Limit: 0,
		Offset: 0,
	}
	users, err = testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 0)
}

func TestUpdateUser(t *testing.T){
	user1:= createRandomUser(t)
	arg := UpdateUserParams{
		ID: user1.ID,
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
		Email: util.RandomEmail(),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, user1.ID, user2.ID)

	user3 , err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.Equal(t, user2.Username, user3.Username)
	require.Equal(t, user2.Password, user3.Password)
	require.Equal(t, user2.Email, user3.Email)

	arg.ID = -1
	_, err = testQueries.UpdateUser(context.Background(), arg)
	require.Error(t, err)
}