package postgres

import (
	"GameSport/src/game/domain/model"
	"context"
	"strconv"
	"strings"
)

func (r *GameRepository) Update(ctx context.Context, id int, game *model.Game) error {
	var fields []string
	var args []interface{}
	argID := 1

	if game.Status != "" {
		fields = append(fields, "status = $"+strconv.Itoa(argID))
		args = append(args, game.Status)
		argID++
	}

	args = append(args, id)
	query := "UPDATE games SET " + strings.Join(fields, ", ") + " WHERE id = $" + strconv.Itoa(argID)

	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}
