package database

import (
	"context"
	"database/sql"
)

type Tx struct {
	tx_       *sql.Tx
	hasError_ bool // 有一些错误 - -
}

func (t *Tx) Close() error {
	if t.hasError_ {
		return t.tx_.Rollback()
	}

	if err := t.tx_.Commit(); err != nil {
		return t.tx_.Rollback()
	}
	return nil
}

func (t *Tx) HasError() {
	t.hasError_ = true
}

func (t *Tx) Insert(query string, args ...any) (int64, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (t *Tx) InsertContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (t *Tx) ExecContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) Delete(query string, args ...any) (int64, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) DeleteContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) Update(query string, args ...any) (int64, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) UpdateContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) QueryForMap(query string, args ...any) (map[string]any, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmtMap(stmt, args...)
}

func (t *Tx) QueryMapContext(ctx context.Context, query string, args ...any) (map[string]any, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmtMapCtx(ctx, stmt, args...)
}

func (t *Tx) QueryForMapSlice(query string, args ...any) ([]map[string]any, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmtMapSlice(stmt, args...)
}

func (t *Tx) QueryMapSliceContext(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmtMapSliceCtx(ctx, stmt, args...)
}
