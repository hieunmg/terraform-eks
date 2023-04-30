package storage

import (
	"context"
	"weshare/modules/authentication/model"
)

func (repository *Repository) Login(
	ctx context.Context,
	username string) (*model.Account, error) {

	query := `
	SELECT id, username, password, salt, full_name, status, birthday, created_at, updated_at FROM accounts
	WHERE username = $1 LIMIT 1`

	row := repository.db.QueryRowContext(ctx, query, username)
	var account model.Account
	err := row.Scan(
		&account.Id,
		&account.Username,
		&account.Password,
		&account.Salt,
		&account.FullName,
		&account.Status,
		&account.Birthday,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	return &account, err
}
