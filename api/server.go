package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/go_todos/db/sqlc"
	"github.com/go_todos/util"
)


type Server struct {
	config util.Config
	store *db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	server := &Server{store: store}
	


	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter(){
	router := gin.Default()

	router.GET("/", server.checkHealth )

	server.router = router

}


func (server *Server) Start(address string) error  {
	return  server.router.Run(address)
}

func errorMessgae(err error ) gin.H {
	return gin.H{"error":err.Error()}
}