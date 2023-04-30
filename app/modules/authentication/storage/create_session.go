package storage

import (
	"context"
	"weshare/modules/authentication/model"
)

func (repository *Repository) CreateSession(ctx context.Context, req *model.SessionCreate) error {
	query := `
	INSERT INTO sessions (
	id,
	account_id,
	refresh_token,
	is_blocked,
	expired_at
	) VALUES (
	$1, $2, $3, $4, $5
	) RETURNING id, account_id, refresh_token, is_blocked, expired_at, created_at
	`
	row := repository.db.QueryRowContext(ctx, query,
		req.Id,
		req.AccountId,
		req.RefreshToken,
		req.IsBlocked,
		req.ExpiredAt,
	)
	var i model.Session
	err := row.Scan(
		&i.Id,
		&i.AccountId,
		&i.RefreshToken,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return err
}
