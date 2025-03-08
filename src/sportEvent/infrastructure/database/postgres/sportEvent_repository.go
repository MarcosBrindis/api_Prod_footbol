package postgres

import (
	"GameSport/src/sportEvent/domain/ports"
	"database/sql"
)

type SportsEventRepository struct {
	db *sql.DB
}

// inicializa el repositorio.
func NewSportEventRepository(db *sql.DB) ports.SportsEventRepository {
	return &SportsEventRepository{db: db}
}
