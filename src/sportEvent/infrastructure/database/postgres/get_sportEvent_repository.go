package postgres

import (
	"GameSport/src/sportEvent/domain/model"
	"context"
)

func (r *SportsEventRepository) GetByID(ctx context.Context, id int) (*model.SportsEvent, error) {
	query := `SELECT id, match_id, event_type, description, timestamp,scoreboard FROM sports_events WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	sportsEvent := &model.SportsEvent{}
	err := row.Scan(&sportsEvent.ID,
		&sportsEvent.MatchID,
		&sportsEvent.EventType,
		&sportsEvent.Description,
		&sportsEvent.Timestamp,
		&sportsEvent.Scoreboard,
	)
	return sportsEvent, err
}
