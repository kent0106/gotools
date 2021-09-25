// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package useany_test

import (
	"testing"

	"github.com/kent0106/gotools/go/analysis/analysistest"
	"github.com/kent0106/gotools/internal/lsp/analysis/useany"
	"github.com/kent0106/gotools/internal/typeparams"
)

func Test(t *testing.T) {
	if !typeparams.Enabled {
		t.Skip("type params are not enabled")
	}
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, useany.Analyzer, "a")
}
