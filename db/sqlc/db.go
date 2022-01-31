// Code generated by sqlc. DO NOT EDIT.

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
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.findAccountStmt, err = db.PrepareContext(ctx, findAccount); err != nil {
		return nil, fmt.Errorf("error preparing query FindAccount: %w", err)
	}
	if q.findAccountsStmt, err = db.PrepareContext(ctx, findAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query FindAccounts: %w", err)
	}
	if q.updateAccountBalanceStmt, err = db.PrepareContext(ctx, updateAccountBalance); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAccountBalance: %w", err)
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
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.findAccountStmt != nil {
		if cerr := q.findAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAccountStmt: %w", cerr)
		}
	}
	if q.findAccountsStmt != nil {
		if cerr := q.findAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAccountsStmt: %w", cerr)
		}
	}
	if q.updateAccountBalanceStmt != nil {
		if cerr := q.updateAccountBalanceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAccountBalanceStmt: %w", cerr)
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
	db                       DBTX
	tx                       *sql.Tx
	createAccountStmt        *sql.Stmt
	deleteAccountStmt        *sql.Stmt
	findAccountStmt          *sql.Stmt
	findAccountsStmt         *sql.Stmt
	updateAccountBalanceStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		createAccountStmt:        q.createAccountStmt,
		deleteAccountStmt:        q.deleteAccountStmt,
		findAccountStmt:          q.findAccountStmt,
		findAccountsStmt:         q.findAccountsStmt,
		updateAccountBalanceStmt: q.updateAccountBalanceStmt,
	}
}
