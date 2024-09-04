package pgsql

import (
	"context"
	"fmt"
	"stepic-go-basic/micro/pkg/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func (d *DB) Redirect(ctx context.Context, s string) (string, error) {

	rows, err := d.pool.Query(ctx, `
		SELECT orig
		FRoM urls WHERE short LIKE '%' || $1 || '%'
		`, s,
	)
	if err != nil {
		println(err)
		return "Ошибка", err

	}
	defer rows.Close()

	var b string
	for rows.Next() {

		err := rows.Scan(
			&b,
		)
		if err != nil {
			return "", err
		}

	}
	err = rows.Err()
	if err != nil {
		return "", err
	}
	return b, nil

}

func (d *DB) AddURL(ctx context.Context, url db.URL) error {

	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(ctx)

	// пакетный запрос
	batch := &pgx.Batch{}
	// добавление заданий в пакет

	batch.Queue(`INSERT INTO urls(short, orig) VALUES ($1, $2)`, url.Short, url.Orig)

	res := tx.SendBatch(ctx, batch)

	err = res.Close()
	if err != nil {
		return err
	}
	// отправка пакета в БД (может выполняться для транзакции или соединения)

	// обязательная операция закрытия соединения

	// подтверждение транзакции
	return tx.Commit(ctx)
}

func NewDB(ctx context.Context, connString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &DB{pool: pool}, nil
}

// Close закрывает пул соединений
func (db *DB) Close() {
	db.pool.Close()
}
