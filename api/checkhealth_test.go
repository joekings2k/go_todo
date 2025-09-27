package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/go_todos/db/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)


func TestCheckHealth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)
	server := newTestServer(t,store)
	recorder := httptest.NewRecorder()
	url := "/"
	request,err := http.NewRequest(http.MethodGet, url,nil)
	require.NoError(t,err)
	server.router.ServeHTTP(recorder,request)
	require.Equal(t,http.StatusOK, recorder.Code)
}