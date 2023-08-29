package parser

import (
	"errors"
	"strings"
)

// Проверка названия сегмента

func ParseSegmentName(name string) (string, error) {
	name = strings.ToUpper(name) // Преобразование в верхний регистр

	if !strings.HasPrefix(name, "AVITO_") {
		return "", errors.New("Segment name must start with 'AVITO_'")
	}

	return name, nil
}
