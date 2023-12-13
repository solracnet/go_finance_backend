package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/solracnet/go_finance_backend/db/sqlc"
)

type Server struct {
	store  *db.SqlStore
	router *gin.Engine
}

func NewServer(store *db.SqlStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/:id", server.getUserById)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}