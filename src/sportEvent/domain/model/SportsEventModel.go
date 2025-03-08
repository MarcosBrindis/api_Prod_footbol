package model

import "time"

type SportsEvent struct {
	ID          int       `json:"id"`
	MatchID     int       `json:"match_id"`    // ID del partido relacionado
	EventType   string    `json:"event_type"`  // Tipo de evento (gol, tarjeta_amarilla, etc.)
	Description string    `json:"description"` // Detalles adicionales (jugador, minuto, etc.)
	Timestamp   time.Time `json:"timestamp"`   // Momento en que ocurrió el evento
	Scoreboard  string    `json:"scoreboard"`  // Marcador del partido
}
