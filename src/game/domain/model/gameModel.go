package model

import "time"

type Game struct {
	ID       int       `json:"id"`
	HomeTeam string    `json:"home_team"` // Nombre del equipo local
	AwayTeam string    `json:"away_team"` // Nombre del equipo visitante
	DateTime time.Time `json:"date_time"` // Fecha y hora del partido
	Status   string    `json:"status"`    // Estado del partido (waiting ,live, Finalized)
}
