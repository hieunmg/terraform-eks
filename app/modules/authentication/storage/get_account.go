package storage

import (
	"context"
	"weshare/modules/authentication/model"
)

func (repository *Repository) GetAccountById(ctx context.Context, id uint32) (*model.Account, error) {

	query := `
	SELECT id, username, full_name 
	FROM accounts
	WHERE id = $1 LIMIT 1
	`
	row := repository.db.QueryRowContext(ctx, query, id)

	var account model.Account

	err := row.Scan(
		&account.Id,
		&account.Username,
		&account.FullName,
	)

	return &account, err
}
