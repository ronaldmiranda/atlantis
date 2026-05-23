// Copyright 2026 The Atlantis Authors
// SPDX-License-Identifier: Apache-2.0

package i18n

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	EnglishLanguage = "en"
	SpanishLanguage = "es"
	DefaultLanguage = EnglishLanguage
)

var supportedLanguages = []string{
	EnglishLanguage,
	SpanishLanguage,
}

// SupportedLanguages returns a copy of all supported language codes.
func SupportedLanguages() []string {
	return slices.Clone(supportedLanguages)
}

// NormalizeLanguageCode canonicalizes the user language input.
func NormalizeLanguageCode(code string) string {
	normalized := strings.TrimSpace(strings.ToLower(code))
	if normalized == "" {
		return DefaultLanguage
	}
	if strings.Contains(normalized, "-") {
		parts := strings.SplitN(normalized, "-", 2)
		normalized = parts[0]
	}
	return normalized
}

// IsSupportedLanguage returns whether the language code is supported.
func IsSupportedLanguage(code string) bool {
	return slices.Contains(supportedLanguages, NormalizeLanguageCode(code))
}

// SupportedLanguagesDescription returns a stable human-readable language list.
func SupportedLanguagesDescription() string {
	return strings.Join(supportedLanguages, ", ")
}

// ValidateLanguage returns an error if the language is not supported.
func ValidateLanguage(code string) error {
	normalized := NormalizeLanguageCode(code)
	if IsSupportedLanguage(normalized) {
		return nil
	}
	return fmt.Errorf("unsupported language %q: supported languages are %s", normalized, SupportedLanguagesDescription())
}

// Translator contains localized strings for comment rendering.
type Translator struct {
	languageCode string
}

// NewTranslator creates a translator using a supported language.
func NewTranslator(code string) (*Translator, error) {
	normalized := NormalizeLanguageCode(code)
	if err := ValidateLanguage(normalized); err != nil {
		return nil, err
	}
	return &Translator{languageCode: normalized}, nil
}

// MustNewTranslator creates a translator or falls back to English.
func MustNewTranslator(code string) *Translator {
	translator, err := NewTranslator(code)
	if err != nil {
		translator, _ = NewTranslator(DefaultLanguage)
	}
	return translator
}

// LanguageCode returns the normalized language code.
func (t *Translator) LanguageCode() string {
	return t.languageCode
}

// CommandTitle returns the display title for a command name.
func (t *Translator) CommandTitle(commandName string) string {
	normalized := strings.TrimSpace(strings.ToLower(commandName))
	switch t.languageCode {
	case SpanishLanguage:
		switch normalized {
		case "apply":
			return "Aplicar"
		case "plan":
			return "Planificar"
		case "unlock":
			return "Desbloquear"
		case "policy_check":
			return "Verificar políticas"
		case "approve_policies":
			return "Aprobar políticas"
		case "version":
			return "Versión"
		case "import":
			return "Importar"
		case "state":
			return "Estado"
		case "cancel":
			return "Cancelar"
		}
	}
	return cases.Title(language.English).String(strings.ReplaceAll(normalized, "_", " "))
}

// PullRequestLabel returns a localized pull request label.
func (t *Translator) PullRequestLabel() string {
	if t.languageCode == SpanishLanguage {
		return "Solicitud de extracción"
	}
	return "Pull Request"
}

// MergeRequestLabel returns a localized merge request label.
func (t *Translator) MergeRequestLabel() string {
	if t.languageCode == SpanishLanguage {
		return "Solicitud de fusión"
	}
	return "Merge Request"
}
