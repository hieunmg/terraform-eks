package storage

import "context"

func (repository *Repository) Logout(ctx context.Context, accountId uint32) error {

	query := `DELETE FROM sessions WHERE account_id = $1`
	_, err := repository.db.ExecContext(ctx, query, accountId)

	return err
}
