package controller

import (
	"encoding/json"
	"net/http"

	"github.com/davidfim/atmosphere-fc/player"
)

type PlayerController struct {
	service *player.PlayerService
}

func NewPlayerController(service *player.PlayerService) *PlayerController {
	return &PlayerController{service: service}
}

func (c *PlayerController) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var p player.Player
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreatePlayer(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *PlayerController) RemovePlayer(w http.ResponseWriter, r *http.Request) {
	var p player.Player
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.RemovePlayer(p.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *PlayerController) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := c.service.GetPlayers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(players); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
