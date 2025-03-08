package postgres

import (
	"GameSport/src/game/domain/model"
	"context"
)

func (r *GameRepository) Create(ctx context.Context, game *model.Game) error {
	game.Status = "waiting"
	query := `INSERT INTO games (home_team, away_team, date_time, status) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRowContext(ctx, query,
		game.HomeTeam,
		game.AwayTeam,
		game.DateTime,
		game.Status,
	).Scan(&game.ID)
}
