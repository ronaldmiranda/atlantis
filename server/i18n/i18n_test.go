// Copyright 2026 The Atlantis Authors
// SPDX-License-Identifier: Apache-2.0

package i18n_test

import (
	"testing"

	"github.com/runatlantis/atlantis/server/i18n"
	. "github.com/runatlantis/atlantis/testing"
)

func TestNormalizeLanguageCode(t *testing.T) {
	Equals(t, "en", i18n.NormalizeLanguageCode(""))
	Equals(t, "es", i18n.NormalizeLanguageCode("es-MX"))
	Equals(t, "en", i18n.NormalizeLanguageCode(" EN "))
}

func TestValidateLanguage(t *testing.T) {
	Ok(t, i18n.ValidateLanguage("en"))
	Ok(t, i18n.ValidateLanguage("es-MX"))
	ErrEquals(t, `unsupported language "de": supported languages are en, es`, i18n.ValidateLanguage("de"))
}

func TestTranslator_CommandTitle(t *testing.T) {
	translator, err := i18n.NewTranslator("es")
	Ok(t, err)
	Equals(t, "Aplicar", translator.CommandTitle("apply"))
	Equals(t, "Verificar políticas", translator.CommandTitle("policy_check"))
}
