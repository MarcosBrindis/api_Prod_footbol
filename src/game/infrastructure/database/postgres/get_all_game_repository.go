package postgres

import (
	"GameSport/src/game/domain/model"
	"context"
)

func (r *GameRepository) GetAll(ctx context.Context) ([]*model.Game, error) {
	query := `SELECT id, home_team,  away_team, date_time, status FROM games`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []*model.Game
	for rows.Next() {
		game := &model.Game{}
		if err := rows.Scan(&game.ID, &game.HomeTeam, &game.AwayTeam, &game.DateTime, &game.Status); err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}
