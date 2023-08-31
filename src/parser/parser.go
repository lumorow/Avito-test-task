package parser

import (
	"errors"
	_ "github.com/bmuller/arrow"
	"strings"
	"time"
)

// Проверка названия сегмента

func ParseSegmentName(name string) (string, error) {
	name = strings.ToUpper(name) // Преобразование в верхний регистр

	if !strings.HasPrefix(name, "AVITO_") {
		return "", errors.New("segment name must start with 'AVITO_'")
	}

	return name, nil
}

func ParseInputTime(inputTime string) (string, error) {
	layout := "2006-01-02"
	inputTime += "-01"
	date, err := time.Parse(layout, inputTime)
	//parsed, err := arrow.CParse("%Y-%m", inputTime)
	if err != nil {
		return "", err
	}
	return date.Format(layout), nil
}
