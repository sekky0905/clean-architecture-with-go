package rdb

import (
	"context"
	"database/sql"
)

// SQLManagerInterface は、SQLManagerのinterface。
type SQLManagerInterface interface {
	Executor
	Preparer
	Queryer
}

// DBに関するInterfaceの定義。
type (
	// Executor は、Rowが返ってこないQueryを実行するためのメソッドを集めたinterface。
	Executor interface {
		// Exec は、Rowが返ってこないQueryを実行する
		Exec(query string, args ...interface{}) (Result, error)
		// ExecContext は、Rowが返ってこないQueryを実行する
		ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)
	}

	// Preparer は、後でQueryやExecを行うために、準備された状態にするメソッドを集めたinterface。
	Preparer interface {
		// Prepare は、後でQueryやExecを行うために、準備された状態にする。
		// callerは、準備された状態が必要ない場合には、Closeを呼び出す必要がある。
		Prepare(query string) (*sql.Stmt, error)
		// PrepareContext は、後でQueryやExecを行うために、準備された状態にする。
		// callerは、準備された状態が必要ない場合には、Closeを呼び出す必要がある。
		// 引数で渡されたcontextは、状態の準備に使用するのであって、状態の実行に使用するのではない。
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	}

	// Queryer は、rowを返すようなQueryを実行するメソッドを集めたinterface。
	Queryer interface {
		// Query は、rowを返すようなQueryを実行する。SELECTが典型的な例。
		Query(query string, args ...interface{}) (Rows, error)
		// QueryContext は、rowを返すようなQueryを実行する。SELECTが典型的な例。
		QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
	}

	// Row は、1行を選択するQueryRowの結果を表す。
	Row interface {
		// Scan は、destに現在読み込んでいるrowのcolumnsをコピーする。
		Scan(...interface{}) error
	}

	// Rows は、Queryの結果を表す。
	Rows interface {
		Row
		// Next は、Scanによって読み込まれる次のrowの結果を準備する。
		// Nextが成功した場合には、trueを返して、失敗した場合にはfalseを返す。
		Next() bool
		// Close は、Rowsを終了する。
		// Nextが呼ばれて、falseが返ってきて結果がそれ以上ない場合は、Rowsは自動でCloseする。
		// Errの結果を確認するにはそれで十分である。
		// Closeを使うと冪等になるし、Errの結果に影響を受けなくなる。
		Close() error
	}

	// Result は、DBからのレスポンスを扱うメソッドを集めたinterface。
	Result interface {
		// LastInsertId は、DBのレスポンスによって生成された数字を返す。
		LastInsertId() (int64, error)
		// RowsAffected は、update、insert、deleteで影響を受けたrowの数を返す。
		RowsAffected() (int64, error)
	}
)
