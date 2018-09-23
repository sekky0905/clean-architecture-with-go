package rdb

import (
	"context"
	"database/sql"
	// SQL Driver。
	_ "github.com/go-sql-driver/mysql"
)

// SQLManager は、SQLを管理する。
type SQLManager struct {
	Conn *sql.DB
}

// NewSQLManager は、SQLManagerを生成し、返す。
func NewSQLManager() SQLManagerInterface {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err.Error())
	}

	return &SQLManager{
		Conn: conn,
	}
}

// Exec は、SQL実行する。
func (s *SQLManager) Exec(query string, args ...interface{}) (Result, error) {
	return s.Conn.Exec(query, args...)
}

// ExecContext は、SQL実行する。
func (s *SQLManager) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error) {
	return s.Conn.ExecContext(ctx, query, args...)
}

// Query は、rowを返すようなQueryを実行する。
func (s *SQLManager) Query(query string, args ...interface{}) (Rows, error) {
	rows, err := s.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	row := &SQLRowManager{
		Rows: rows,
	}
	return row, nil
}

// QueryContext は、rowを返すようなQueryを実行する。
func (s *SQLManager) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rows, err := s.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	row := &SQLRowManager{
		Rows: rows,
	}
	return row, nil
}

// Prepare は、後でQueryやExecを行うために、準備された状態にする。
func (s *SQLManager) Prepare(query string) (*sql.Stmt, error) {
	return s.Conn.Prepare(query)
}

// PrepareContext は、後でQueryやExecを行うために、準備された状態にする。
func (s *SQLManager) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return s.Conn.PrepareContext(ctx, query)
}

// SQLRowManager は、Rowを管理する。
type SQLRowManager struct {
	Rows *sql.Rows
}

// Scan は、destに現在読み込んでいるrowのcolumnsをコピーする。
func (r SQLRowManager) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next は、Scanによって読み込まれる次のrowの結果を準備する。
func (r SQLRowManager) Next() bool {
	return r.Rows.Next()
}

// Close は、Rowsを終了する。
func (r SQLRowManager) Close() error {
	return r.Rows.Close()
}
