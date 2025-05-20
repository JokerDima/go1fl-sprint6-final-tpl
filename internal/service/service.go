package service

import (
	"errors"
	"log"
	"strings"

	"go1fl-sprint6-final-tpl/pkg/morse"
)

func Replace(s string) (string, error) {
	if len(s) == 0 {
		log.Fatal("no data available")
		return "", errors.New("no data available")
	}

	if strings.ContainsAny(s, "ауоиэыяюеё") {
		return morse.ToMorse(s), nil
	} else {
		return morse.ToText(s), nil
	}
}
