package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Player struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router

	PlayerList []Player //in real case not defined here
}

func NewServer() *Server {
	s := &Server{
		Router:     mux.NewRouter(),
		PlayerList: []Player{},
	}
	s.routes()
	return s
}

func (s *Server) createPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Player
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		p.ID = uuid.New()
		s.PlayerList = append(s.PlayerList, p)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removePlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, player := range s.PlayerList {
			if player.ID == id {
				s.PlayerList = append(s.PlayerList[:i], s.PlayerList[i+1:]...)
				break
			}
		}
	}
}

func (s *Server) listPlayers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.PlayerList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) routes() {
	s.HandleFunc("/players", s.listPlayers()).Methods("GET")
	s.HandleFunc("/players", s.createPlayer()).Methods("POST")
	s.HandleFunc("/players/{id}", s.removePlayer()).Methods("DELETE")

}
