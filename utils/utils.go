package utils

import (
	"regexp"
	"strings"
)

type KnownUtils struct {
}

type UtilsInterface interface {
	GenerateSlug(token string) string
}

func NewUtils() UtilsInterface {
	return &KnownUtils{}
}

func (u *KnownUtils) GenerateSlug(token string) string {

	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	clean := re.ReplaceAllString(token, "")
	clean = strings.ToLower(clean)
	clean = strings.ReplaceAll(clean, " ", "-")
	return clean
}
