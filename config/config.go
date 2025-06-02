package config

import (
	"cmp"
	"os"

	"github.com/ascii-arcade/moonrollers/language"
	"github.com/charmbracelet/log"
)

const (
	MinimumHeight = 40
	MinimumWidth  = 120

	Host = "localhost"
	Port = "23234"
)

var (
	Language *language.Language = setDefaultLanguage()

	Version string = "dev"
)

func setDefaultLanguage() *language.Language {
	langCode := cmp.Or(os.Getenv("ASCII_ARCADE_LANG"), "EN")
	lang, exists := language.Languages[langCode]
	if !exists {
		log.Warn("Unknown language code %s, defaulting to English", langCode)
		return language.Languages["EN"]
	}
	return lang
}
