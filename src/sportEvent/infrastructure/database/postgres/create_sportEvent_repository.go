package postgres

import (
	"GameSport/src/sportEvent/domain/model"
	"context"
)

func (r *SportsEventRepository) Create(ctx context.Context, sportsEvent *model.SportsEvent) error {
	query := `INSERT INTO sports_events (match_id, event_type, description, timestamp, Scoreboard) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRowContext(ctx, query,
		sportsEvent.MatchID,
		sportsEvent.EventType,
		sportsEvent.Description,
		sportsEvent.Timestamp,
		sportsEvent.Scoreboard,
	).Scan(&sportsEvent.ID)
}
