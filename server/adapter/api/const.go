package api

// パスの定義。
const (
	ProgrammingLangAPIPath = "/langs"
)

// クエリストリングの属性。
const (
	Limit = "limit"
)

// Limitの定義。
const (
	MaxLimit     = 100
	MinLimit = 5
	DefaultLimit = 20
)

// パラメータの属性
const (
	ID = "id"
)

// HTTPのメソッド。
const (
	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"
)
