package model

import "fmt"

// InvalidPropertyError は、Propertyが不適切な場合のエラー。
// 複数のpackageで使用するため、interfaceを通してではなく、構造体のインスタンスを使用する
type InvalidPropertyError struct {
Property string
Message  string
}

// Error は、エラー文を返す。
func (e *InvalidPropertyError) Error() string {
return fmt.Sprintf("%s is invalid. %s", e.Property, e.Message)
}