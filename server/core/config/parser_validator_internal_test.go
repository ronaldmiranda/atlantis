// Copyright 2025 The Atlantis Authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	. "github.com/runatlantis/atlantis/testing"
	yaml "gopkg.in/yaml.v3"
)

type panicUnmarshaler struct{}

func (p *panicUnmarshaler) UnmarshalYAML(*yaml.Node) error {
	panic("decode panic")
}

func TestDecodeYAMLKnownFields_RecoversPanic(t *testing.T) {
	var out struct {
		Value panicUnmarshaler `yaml:"value"`
	}

	err := decodeYAMLKnownFields([]byte("value: test"), &out)
	ErrContains(t, "panic while parsing yaml: decode panic", err)
}
