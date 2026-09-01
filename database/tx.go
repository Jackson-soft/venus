package database

import (
	"context"
	"database/sql"
	"errors"
)

type Tx struct {
	tx_       *sql.Tx
	hasError_ bool // 有错误需要回滚
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

// Close 提交或回滚事务：若期间任一操作出错则回滚，否则提交。
// 与 sql.Tx 一样，Tx 应由单一 goroutine 使用。
func (t *Tx) Close() error {
	if t.hasError_ {
		return t.tx_.Rollback()
	}

	err := t.tx_.Commit()
	if err != nil {
		// 提交失败时回滚；若回滚也失败，则合并两者，保留提交错误。
		rbErr := t.tx_.Rollback()
		if rbErr != nil {
			return errors.Join(err, rbErr)
		}

		return err
	}

	return nil
}

// HasError 手动标记事务出错，使 Close 执行回滚。
// 通常无需手动调用：Exec/Query 等方法出错时会自动标记。
func (t *Tx) HasError() {
	t.hasError_ = true
}

func (t *Tx) InsertContext(ctx context.Context, query string, args ...any) (int64, error) {
	res, err := t.tx_.ExecContext(ctx, query, args...)
	if err != nil {
		t.markError()

		return 0, err
	}

	return res.LastInsertId()
}

func (t *Tx) ExecContext(ctx context.Context, query string, args ...any) (int64, error) {
	res, err := t.tx_.ExecContext(ctx, query, args...)
	if err != nil {
		t.markError()

		return 0, err
	}

	return res.RowsAffected()
}

func (t *Tx) QueryMapContext(ctx context.Context, query string, args ...any) (map[string]any, error) {
	rows, err := t.tx_.QueryContext(ctx, query, args...)
	if err != nil {
		t.markError()

		return nil, err
	}

	result, err := rowMap(rows)
	if err != nil {
		t.markError()

		return nil, err
	}

	return result, nil
}

func (t *Tx) QueryMapSliceContext(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	rows, err := t.tx_.QueryContext(ctx, query, args...)
	if err != nil {
		t.markError()

		return nil, err
	}

	results, err := rowMapSlice(rows)
	if err != nil {
		t.markError()

		return nil, err
	}

	return results, nil
}

func (t *Tx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		t.markError()

		return nil, err
	}

	return stmt, nil
}

// markError 记录错误状态。
func (t *Tx) markError() {
	t.hasError_ = true
}
