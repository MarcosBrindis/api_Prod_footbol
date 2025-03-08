package postgres

import (
	"GameSport/src/game/domain/model"
	"context"
)

func (r *GameRepository) GetByID(ctx context.Context, id int) (*model.Game, error) {
	query := `SELECT id, home_team, away_team, date_time, status FROM games WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	game := &model.Game{}
	err := row.Scan(&game.ID,
		&game.HomeTeam,
		&game.AwayTeam,
		&game.DateTime,
		&game.Status,
	)
	return game, err
}
