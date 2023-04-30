package storage

import (
	"context"
	"weshare/modules/authentication/model"

	"github.com/google/uuid"
)

func (repository *Repository) GetSession(ctx context.Context, id uuid.UUID) (*model.Session, error) {

	query := `
	SELECT id, account_id, refresh_token, is_blocked, expired_at 
	FROM sessions
	WHERE id = $1 LIMIT 1
	`
	row := repository.db.QueryRowContext(ctx, query, id)

	var session model.Session
	err := row.Scan(
		&session.Id,
		&session.AccountId,
		&session.RefreshToken,
		&session.IsBlocked,
		&session.ExpiredAt,
	)
	return &session, err
}
