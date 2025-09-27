package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/go_todos/db/sqlc"
	"github.com/go_todos/util"
	"github.com/lib/pq"
)


type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=1"`
	Password string `json:"password" binding:"required,min=6"`
	Email string `json:"email" binding:"required,email"`
}

type createUserResponse struct{
	Username string `json:"username"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) createUserResponse {
	return createUserResponse{
		Username: user.Username,
		Email: user.Email,
		CreatedAt: user.CreatedAt.Time,
	}
}
func (server *Server) createUser(ctx *gin.Context){
	var req createUserRequest
	if err:= ctx.ShouldBindJSON(&req); err!= nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err  := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	user,err := server.store.CreateUser(ctx, arg)
	if err != nil{
		if pqErr, ok := err.(*pq.Error);ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
	}
	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}
