package connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectionDB(ctx context.Context) (*pgx.Conn, error) {
	conn_string := os.Getenv("CONN_STRING")

	conn, err := pgx.Connect(ctx, conn_string)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
