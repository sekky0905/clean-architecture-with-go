package model

import "fmt"

// RequiredError は、必要なものが存在しない場合のエラー。
type RequiredError struct {
	Property string
}

// Error は、エラー文を返す。
func (e *RequiredError) Error() string {
	return fmt.Sprintf("%s is required", e.Property)
}

// InvalidPropertyError は、Propertyが不適切な場合のエラー。
type InvalidPropertyError struct {
	Property string
	Message  string
}

// Error は、エラー文を返す。
func (e *InvalidPropertyError) Error() string {
	return fmt.Sprintf("%s is invalid. %s", e.Property, e.Message)
}

// InvalidParameterError は、Parameterが不適切な場合のエラー。
type InvalidParameterError struct {
	Parameter string
	Message   string
}

// Error は、エラー文を返す。
func (e *InvalidParameterError) Error() string {
	return fmt.Sprintf("%s is invalid. %s", e.Parameter, e.Message)
}

// DBError は、DBのエラーを表す。
type DBError struct {
	ModelName string
	DBMethod  string
	Detail    string
}

// Error は、エラー文を返す。
func (e *DBError) Error() string {
	return fmt.Sprintf("failed DB operation. Method : %s, Model : %s, Detail : %s", e.DBMethod, e.ModelName, e.Detail)
}

// AlreadyExistError は、既に存在することを表すエラー。
type AlreadyExistError struct {
	ID        int
	Name      string
	ModelName string
}

// Error は、エラーメッセージを返却する。
func (e *AlreadyExistError) Error() string {
	return fmt.Sprintf("already exists. model: %s, id: %d, name: %s", e.ModelName, e.ID, e.Name)
}

// NoSuchDataError は、データが存在しないことを表すエラー。
type NoSuchDataError struct {
	ID        int
	Name      string
	ModelName string
}

// Error は、エラーメッセージを返す。
func (e *NoSuchDataError) Error() string {
	return fmt.Sprintf("no such model.model: %s, id: %d, name: %s", e.ModelName, e.ID, e.Name)
}
