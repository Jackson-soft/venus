package database

import (
	"context"
	"database/sql"
)

type Tx struct {
	tx_       *sql.Tx
	hasError_ bool // 有一些错误 - -
}

func (d *Database) BeginTxCtx(ctx context.Context) (*Tx, error) {
	tx, err := d.conn().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Tx{
		tx_:       tx,
		hasError_: false,
	}, nil
}

func (t *Tx) Close() error {
	if t.hasError_ {
		return t.tx_.Rollback()
	}

	err := t.tx_.Commit()
	if err != nil {
		return t.tx_.Rollback()
	}

	return nil
}

func (t *Tx) HasError() {
	t.hasError_ = true
}

func (t *Tx) InsertContext(ctx context.Context, query string, args ...any) (int64, error) {
	res, err := t.tx_.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (t *Tx) ExecContext(ctx context.Context, query string, args ...any) (int64, error) {
	res, err := t.tx_.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (t *Tx) QueryMapContext(ctx context.Context, query string, args ...any) (map[string]any, error) {
	rows, err := t.tx_.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rowMap(rows)
}

func (t *Tx) QueryMapSliceContext(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	rows, err := t.tx_.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rowMapSlice(rows)
}
