// Copyright 2025 The Atlantis Authors
// SPDX-License-Identifier: Apache-2.0

// Package cfgfuzz contains fuzz tests for the Atlantis config parser.
// It lives in a dedicated subdirectory so that compile_native_go_fuzzer
// only needs to compile this package and its direct imports, avoiding
// the test-only imports in the parent package (e.g. parser_validator_test.go).
package cfgfuzz

import (
	"testing"

	"github.com/runatlantis/atlantis/server/core/config"
	"github.com/runatlantis/atlantis/server/core/config/valid"
)

func FuzzParseRepoCfgData(f *testing.F) {
	f.Add([]byte(`version: 3
projects:
- dir: .
`))
	f.Add([]byte(`version: 3
automerge: true
projects:
- dir: .
  workspace: default
  autoplan:
    when_modified: ["*.tf"]
    enabled: true
`))
	// Regression seed: yaml.v3 merge-key panic (hash of unhashable type []interface{}).
	// The YAML merge key (<<) with an illegal value triggered a panic inside gopkg.in/yaml.v3
	// instead of returning a decode error.  Fixed by wrapping Decode with recover().
	f.Add([]byte("0nnnn:\n\n<< :\n\n0.0:\n"))

	pv := config.ParserValidator{}
	globalCfg := valid.NewGlobalCfgFromArgs(valid.GlobalCfgArgs{AllowAllRepoSettings: true})

	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = pv.ParseRepoCfgData(data, globalCfg, "github.com/test/repo", "main")
	})
}
