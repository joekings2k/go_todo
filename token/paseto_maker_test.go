package token

import (
	"testing"
	"time"

	"github.com/go_todos/util"
	"github.com/stretchr/testify/require"
)


func TestPasetoMaker(t *testing.T) {

	maker, err :=  NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
	email := util.RandomEmail()
	userID := util.RandomInt(1,20)
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(int32(userID),email,duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	
	payload, err := maker.VerifyToken(token)
	require.NotZero(t, payload.ID)
	require.NotZero(t, payload.UserID)
	require.Equal(t, int32(userID),payload.UserID)
	require.Equal(t, email, payload.Email)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t,expiredAt, payload.ExpiredAt,time.Second)
}