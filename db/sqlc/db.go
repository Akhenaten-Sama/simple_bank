// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAccountStmt, err = db.PrepareContext(ctx, createAccount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccount: %w", err)
	}
	if q.createTransferStmt, err = db.PrepareContext(ctx, createTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTransfer: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.deleteTransferStmt, err = db.PrepareContext(ctx, deleteTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTransfer: %w", err)
	}
	if q.getAccountStmt, err = db.PrepareContext(ctx, getAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccount: %w", err)
	}
	if q.getTransferStmt, err = db.PrepareContext(ctx, getTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransfer: %w", err)
	}
	if q.listAccountsStmt, err = db.PrepareContext(ctx, listAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query ListAccounts: %w", err)
	}
	if q.listTransfersStmt, err = db.PrepareContext(ctx, listTransfers); err != nil {
		return nil, fmt.Errorf("error preparing query ListTransfers: %w", err)
	}
	if q.updateAccountStmt, err = db.PrepareContext(ctx, updateAccount); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAccount: %w", err)
	}
	if q.updateTransferStmt, err = db.PrepareContext(ctx, updateTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTransfer: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAccountStmt != nil {
		if cerr := q.createAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountStmt: %w", cerr)
		}
	}
	if q.createTransferStmt != nil {
		if cerr := q.createTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTransferStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.deleteTransferStmt != nil {
		if cerr := q.deleteTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTransferStmt: %w", cerr)
		}
	}
	if q.getAccountStmt != nil {
		if cerr := q.getAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountStmt: %w", cerr)
		}
	}
	if q.getTransferStmt != nil {
		if cerr := q.getTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransferStmt: %w", cerr)
		}
	}
	if q.listAccountsStmt != nil {
		if cerr := q.listAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAccountsStmt: %w", cerr)
		}
	}
	if q.listTransfersStmt != nil {
		if cerr := q.listTransfersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTransfersStmt: %w", cerr)
		}
	}
	if q.updateAccountStmt != nil {
		if cerr := q.updateAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAccountStmt: %w", cerr)
		}
	}
	if q.updateTransferStmt != nil {
		if cerr := q.updateTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTransferStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                 DBTX
	tx                 *sql.Tx
	createAccountStmt  *sql.Stmt
	createTransferStmt *sql.Stmt
	deleteAccountStmt  *sql.Stmt
	deleteTransferStmt *sql.Stmt
	getAccountStmt     *sql.Stmt
	getTransferStmt    *sql.Stmt
	listAccountsStmt   *sql.Stmt
	listTransfersStmt  *sql.Stmt
	updateAccountStmt  *sql.Stmt
	updateTransferStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                 tx,
		tx:                 tx,
		createAccountStmt:  q.createAccountStmt,
		createTransferStmt: q.createTransferStmt,
		deleteAccountStmt:  q.deleteAccountStmt,
		deleteTransferStmt: q.deleteTransferStmt,
		getAccountStmt:     q.getAccountStmt,
		getTransferStmt:    q.getTransferStmt,
		listAccountsStmt:   q.listAccountsStmt,
		listTransfersStmt:  q.listTransfersStmt,
		updateAccountStmt:  q.updateAccountStmt,
		updateTransferStmt: q.updateTransferStmt,
	}
}
