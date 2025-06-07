package server

import (
	"fmt"

	"github.com/AjayKodavati/CMS/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	Router            *gin.Engine
	RepositoryService repository.DBRepositories
}

func NewServer(pool *pgxpool.Pool) *Server {
	repoService := repository.SetUpDBRepositories(pool)
	return &Server{
		Router:            gin.Default(),
		RepositoryService: *repoService,
	}
}

func (s *Server) Start(port string)  {
	s.Router.Run(port)
	fmt.Printf("Server is running on port %s\n", port)
}
