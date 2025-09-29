package api

import (
	"database/sql"
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
	ID int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) createUserResponse {
	return createUserResponse{
		ID: user.ID,
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

type loginUserRequest struct{
	Email string `json:"email" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}
type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	User createUserResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context){
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req);err !=nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
	}
	user,err := server.store.GetUserByEmail(ctx,req.Email)
	if err!= nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	err =util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	} 

	accessToken, err := server.tokenMaker.CreateToken(user.ID, user.Email,server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response:= loginUserResponse{
		AccessToken: accessToken,
		User: newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, response)
}