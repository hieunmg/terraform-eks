package storage

import (
	"context"
	"weshare/modules/authentication/model"
)

func (repository *Repository) Register(
	ctx context.Context,
	req *model.AccountRegisterRequest) error {

	query := `
		INSERT INTO accounts (
		  username,
		  full_name,
		  password,
		  salt
		) VALUES (
		  $1, $2, $3, $4
		)
		`
	row := repository.db.QueryRowContext(
		ctx, query, req.Username,
		req.FullName, req.Password,
		req.Salt,
	)

	return row.Err()
}
