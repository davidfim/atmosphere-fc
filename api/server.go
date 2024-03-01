package api

import (
	"github.com/davidfim/atmosphere-fc/api/controller"
	"github.com/davidfim/atmosphere-fc/player"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	*mux.Router
	playerController *controller.PlayerController
}

func NewServer(db *gorm.DB) *Server {
	playerService := player.NewPlayerService(db)
	playerController := controller.NewPlayerController(playerService)

	s := &Server{
		Router:           mux.NewRouter(),
		playerController: playerController,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/players", s.playerController.GetPlayers).Methods("GET")
	s.HandleFunc("/players", s.playerController.CreatePlayer).Methods("POST")
	s.HandleFunc("/players/{id}", s.playerController.RemovePlayer).Methods("DELETE")
}
