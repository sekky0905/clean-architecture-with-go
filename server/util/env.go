package util

import (
	"github.com/joho/godotenv"
)

// LoadEnv は、.Envを読み込む。
func LoadEnv(fileName string) error {
	return godotenv.Load(fileName)
}
