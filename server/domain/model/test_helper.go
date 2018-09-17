package model

import (
	"fmt"
	"time"
)

// GetTestTime は、テスト用の時刻を生成し、返す。
func GetTestTime(month time.Month, day int) time.Time {
	return time.Date(2018, month, day, 12, 0, 0, 0, time.UTC)
}

// CreateProgrammingLangs は、引数で与えられた数だけProgrammingLangを生成し、返す。
func CreateProgrammingLangs(num int) []*ProgrammingLang {
	langSlice := make([]*ProgrammingLang, num, num)
	for i := range langSlice {
		langSlice[i] = &ProgrammingLang{
			ID:        i + 1,
			Name:      fmt.Sprintf("%s%d", TestName, i),
			Feature:   fmt.Sprintf("%s%d", TestFeature, i),
			CreatedAt: GetTestTime(time.October, i+1),
			UpdatedAt: GetTestTime(time.October, i+1),
		}
	}
	return langSlice
}
