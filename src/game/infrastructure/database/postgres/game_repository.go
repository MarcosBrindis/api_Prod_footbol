package postgres

import (
	"GameSport/src/game/domain/ports"
	"database/sql"
)

type GameRepository struct {
	db *sql.DB
}

// inicializa el repositorio.
func NewGameRepository(db *sql.DB) ports.GameRepository {
	return &GameRepository{db: db}
}
